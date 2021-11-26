package memory

import (
	"github.com/yumaojun03/dmidecode"
)

type MemoryList []*Memory

var (
	_memList MemoryList
)

type Memory struct {
	Name                       string
	Manufacturer               string
	Capacity                   uint64
	PhysicalMemoryArrayHandle  uint16
	ErrorInformationHandle     uint16
	TotalWidth                 uint16
	DataWidth                  uint16
	FormFactor                 string
	DeviceSet                  byte
	DeviceLocator              string
	BankLocator                string
	Type                       string
	TypeDetail                 string
	Speed                      uint16
	SerialNumber               string
	AssetTag                   string
	PartNumber                 string
	Attributes                 byte
	ExtendedSize               uint32
	ConfiguredMemoryClockSpeed uint16
	MinimumVoltage             uint16
	MaximumVoltage             uint16
	ConfiguredVoltage          uint16
}

func GetMemoryList() (MemoryList, error) {
	if len(_memList) == 0 {
		dmi, err := dmidecode.New()
		if err != nil {
			return nil, err
		}

		devices, err := dmi.MemoryDevice()
		for _, device := range devices {
			_memList = append(_memList, &Memory{
				Manufacturer:               device.Manufacturer,
				Capacity:                   uint64(device.Size) * 1024, // MB TO KB
				PhysicalMemoryArrayHandle:  device.PhysicalMemoryArrayHandle,
				ErrorInformationHandle:     device.ErrorInformationHandle,
				TotalWidth:                 device.TotalWidth,
				DataWidth:                  device.DataWidth,
				FormFactor:                 device.FormFactor.String(),
				DeviceSet:                  device.DeviceSet,
				DeviceLocator:              device.DeviceLocator,
				BankLocator:                device.BankLocator,
				Type:                       device.Type.String(),
				TypeDetail:                 device.TypeDetail.String(),
				Speed:                      device.Speed,
				SerialNumber:               device.SerialNumber,
				AssetTag:                   device.AssetTag,
				PartNumber:                 device.PartNumber,
				Attributes:                 device.Attributes,
				ExtendedSize:               device.ExtendedSize,
				ConfiguredMemoryClockSpeed: device.ConfiguredVoltage,
				MinimumVoltage:             device.ConfiguredVoltage,
				MaximumVoltage:             device.ConfiguredVoltage,
				ConfiguredVoltage:          device.ConfiguredVoltage,
			})
		}
	}
	return _memList, nil
}
