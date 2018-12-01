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
			logrus.Fatal(err)
		}
		//power := values["battery_charge"]

		/*
			switch power.(type) {
			case int32:
				fmt.Printf("int32")
			case uint32:
				fmt.Printf("uint32")
			case int64:
				fmt.Printf("int64")
			case uint64:
				fmt.Printf("uint64")
			}

			_ = power

			for k, v := range values {
				fmt.Printf("%s -> %v\n", k, v)
			}

			_ = values*/
		fmt.Printf("%.1f W\n", float64(values["0:1.4.0"].(uint64))/10.0)
		fmt.Printf("%.1f W\n", float64(values["0:1.4.0"].(uint64))/10.0)
	}
}
