package main

//go:generate go run -tags=dev ./asset_generator.go

import (
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"

	"gitlab.com/bboehmke/sunny"
	"gitlab.com/bboehmke/sunny/cmd/monitor/assets"
)

var websocketUpgrader = websocket.Upgrader{}

// SolarState contains current inverter and energy meter values
type SolarState struct {
	// >0 - Solar production
	SolarPower int32 `json:"solar_power"`
	// >0 - buy energy | <0 sell energy
	SupplierPower int32 `json:"supplier_power"`
	// <0 charge | >0 discharge
	BatteryPower int32 `json:"battery_power"`

	HomePower         uint32 `json:"home_power"`
	HomeSupplierPower uint32 `json:"home_supplier_power"`
	HomeOwnPower      uint32 `json:"home_own_power"`

	BatteryCharge uint32 `json:"battery_charge"`
}

// MonitorService for updating the actual solar state
type MonitorService struct {
	sync.RWMutex

	EnergyMeter     *sunny.Device
	SolarInverter   *sunny.Device
	BatteryInverter *sunny.Device

	state SolarState
}

func (s *MonitorService) getState() SolarState {
	s.RLock()
	defer s.RUnlock()
	return s.state
}

func (s *MonitorService) updateLoop() {
	for {
		// no sleep required -> EnergyMeter values will received once a second
		s.update()
	}
}
func (s *MonitorService) update() {
	energyMeterValues, err := s.EnergyMeter.GetValues()
	if err != nil {
		logrus.Warning(err)
		return
	}
	pvValues, err := s.SolarInverter.GetValues()
	if err != nil {
		logrus.Warning(err)
		return
	}
	batteryValues, err := s.BatteryInverter.GetValues()
	if err != nil {
		logrus.Warning(err)
		return
	}

	// get energy meter values
	supplierPlus, ok := energyMeterValues["0:1.4.0"]
	if !ok || supplierPlus == nil {
		supplierPlus = uint64(0)
	}
	supplierMinus, ok := energyMeterValues["0:2.4.0"]
	if !ok || supplierMinus == nil {
		supplierMinus = uint64(0)
	}

	// get pv value
	pvPower := pvValues["power_ac_total"]
	if !ok || pvPower == nil {
		pvPower = int32(0)
	}

	// get battery values
	batteryPower := batteryValues["power_ac_total"]
	if !ok || batteryPower == nil {
		batteryPower = int32(0)
	}
	batteryCharge := batteryValues["battery_charge"]
	if !ok || batteryCharge == nil {
		batteryCharge = uint32(0)
	}

	state := SolarState{
		SolarPower:    pvPower.(int32),
		SupplierPower: int32(supplierPlus.(uint32)-supplierMinus.(uint32)) / 10,
		BatteryPower:  batteryPower.(int32),
		BatteryCharge: batteryCharge.(uint32),
	}

	usage := state.SolarPower + state.SupplierPower + state.BatteryPower
	if usage < 0 {
		return // ignore this update
	}

	state.HomePower = uint32(usage)
	if state.SupplierPower > 0 {
		state.HomeSupplierPower = uint32(state.SupplierPower)
	}
	state.HomeOwnPower = state.HomePower - state.HomeSupplierPower

	s.Lock()
	defer s.Unlock()
	s.state = state
}

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}

func getDeviceFromEnv(key, password string) (*sunny.Device, error) {
	address, ok := os.LookupEnv(key)
	if !ok {
		return nil, nil
	}

	return sunny.NewDevice(address, password)
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableColors:   true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// read configuration from environment variables
	if strings.ToLower(os.Getenv("DEBUG")) != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	password := getEnv("PASSWORD", "0000")

	service := MonitorService{}
	var err error

	service.SolarInverter, err = getDeviceFromEnv("SOLAR_INVERTER", password)
	if err != nil {
		logrus.Fatal(err)
	}
	service.BatteryInverter, err = getDeviceFromEnv("BATTERY_INVERTER", password)
	if err != nil {
		logrus.Fatal(err)
	}
	service.EnergyMeter, err = getDeviceFromEnv("ENERGY_METER", password)
	if err != nil {
		logrus.Fatal(err)
	}

	devices, err := sunny.DiscoverDevices(password)
	if err != nil {
		logrus.Fatal(err)
	}

	for _, dev := range devices {
		deviceClass, err := dev.GetDeviceClass()
		if err != nil {
			logrus.Fatalf("Failed to get device type %v", err)
		}

		switch deviceClass {
		case 1: // energy meter
			if service.EnergyMeter == nil {
				service.EnergyMeter = dev
			}

		case 8001: // solar
			if service.SolarInverter == nil {
				service.SolarInverter = dev
			}

		case 8007: // battery
			if service.BatteryInverter == nil {
				service.BatteryInverter = dev
			}

		default:
			logrus.Warningf("Unknown device class %d", deviceClass)
		}
	}

	if service.EnergyMeter == nil {
		logrus.Fatal("Energy meter missing")
	}
	if service.SolarInverter == nil {
		logrus.Fatal("Solar inverter missing")
	}
	if service.BatteryInverter == nil {
		logrus.Fatal("Battery inverter missing")
	}

	go service.updateLoop()

	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	handler := fsHandler(assets.Assets)
	router.GET("/", handler)
	router.HEAD("/", handler)
	router.GET("/status_img.svg", handler)
	router.HEAD("/status_img.svg", handler)

	router.GET("/ws", ws(&service))

	err = router.Run(":8080")
	if err != nil {
		logrus.Fatal(err)
	}
}

func fsHandler(fs http.FileSystem) gin.HandlerFunc {
	fileServer := http.FileServer(fs)
	return func(ctx *gin.Context) {
		fileServer.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func ws(service *MonitorService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conn, err := websocketUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			logrus.Error("Failed to upgrade protocol:", err)
			return
		}

		counter := 1
		for {
			err := conn.WriteJSON(service.getState())
			if err != nil {
				conn.Close()
				break
			}
			time.Sleep(time.Second)
			counter++
		}
	}
}
