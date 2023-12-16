package main

import (
	"fmt"
	homeautomation "microservice_solid/pkg/home_automation"
)

func main() {
	// Create a new DeviceManager
	deviceManager := &homeautomation.DeviceManager{}

	// Example: Register a device
	device1 := homeautomation.Device{ID: "1", Name: "Smart Light", Type: "Lighting"}
	deviceManager.RegisterDevice(device1)

	// Example: Register another device
	device2 := homeautomation.Device{ID: "2", Name: "Smart Thermostat", Type: "Climate Control"}
	deviceManager.RegisterDevice(device2)

	// Create a new DataManagerImpl
	dataManager := &homeautomation.DataManagerImpl{}

	// Example: Process data
	data := map[string]interface{}{"temperature": 22.5, "humidity": 50}
	err := dataManager.ProcessData(data)
	if err != nil {
		fmt.Println("Error processing data:", err)
	}
}
