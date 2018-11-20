package main

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"gitlab.com/bboehmke/sunny"
)

func main() {
	dev, err := sunny.NewDevice("192.168.2.123", "0000")
	if err != nil {
		logrus.Fatal(err)
	}

	for {
		values, err := dev.GetValues()
		if err != nil {
			logrus.Error(err)
			continue
		}

		fmt.Printf("%.1f W\n", float64(values["0:1.4.0"].(uint64))/10.0)
	}
}
