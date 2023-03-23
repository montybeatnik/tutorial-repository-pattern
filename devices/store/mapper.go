package store

import "github.com/montybeatnik/tutorials/repository-pattern/devices/models"

func mapDeviceAttrs(newDevice models.NewDeviceRequest) models.Device {
	return models.Device{
		CLLI:     newDevice.CLLI,
		Hostname: newDevice.Hostname,
		IP:       newDevice.IP,
	}
}
