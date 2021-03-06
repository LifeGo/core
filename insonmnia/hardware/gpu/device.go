package gpu

import (
	"encoding/json"

	pb "github.com/sonm-io/core/proto"
)

// Device describes a GPU device.
type Device interface {
	// Name returns GPU model name.
	Name() string
	// VendorName returns GPU vendor name.
	VendorName() string
	// MaxMemorySize returns the total maximum memory size the device can hold
	// in bytes.
	MaxMemorySize() uint64
}

type device struct {
	d pb.GPUDevice
}

func NewDevice(name, vendorName string, maxMemorySize uint64) Device {
	return &device{
		d: pb.GPUDevice{
			Name:          name,
			VendorName:    vendorName,
			MaxMemorySize: maxMemorySize,
		},
	}
}

func (d *device) Name() string {
	return d.d.GetName()
}

func (d *device) VendorName() string {
	return d.d.GetVendorName()
}

func (d *device) MaxMemorySize() uint64 {
	return d.d.GetMaxMemorySize()
}

func (d *device) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name":          d.Name(),
		"vendorName":    d.VendorName(),
		"maxMemorySize": d.MaxMemorySize(),
	})
}

// GetGPUDevices returns a list of available GPU devices on the machine.
func GetGPUDevices() ([]Device, error) {
	devices, err := GetGPUDevicesUsingOpenCL()
	if err != nil {
		return nil, err
	}

	return devices, nil
}
