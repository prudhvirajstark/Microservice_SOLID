package homeautomation

import "fmt"

// Device struct represents a smart device
type Device struct {
	ID   string
	Name string
	Type string
}

// DeviceManager struct implements DataManager for device-related operations
type DeviceManager struct {
	RegisteredDevices []Device
}

// RegisterDevice registers a new smart device
func (dm *DeviceManager) RegisterDevice(device Device) {
	dm.RegisteredDevices = append(dm.RegisteredDevices, device)
	fmt.Printf("Device %s registered successfully.\n", device.Name)
}