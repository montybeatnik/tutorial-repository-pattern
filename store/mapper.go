package store

import "github.com/montybeatnik/tutorials/repository-pattern/devices/models"

// mapDeviceAttrs assings the NewDeviceRequest fields to a device,
// returning it to the caller.
func mapDeviceAttrs(newDevice models.NewDeviceRequest) models.Device {
	return models.Device{
		CLLI:     newDevice.CLLI,
		Hostname: newDevice.Hostname,
		IP:       newDevice.IP,
	}
}
