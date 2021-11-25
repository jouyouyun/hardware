package memory

import (
	"github.com/yumaojun03/dmidecode"
	"github.com/yumaojun03/dmidecode/parser/memory"
)

func GetMemoryList() ([]*memory.MemoryDevice, error) {
	dmi, err := dmidecode.New()
	if err != nil {
		return nil, err
	}

	devices, err := dmi.MemoryDevice()
	return devices, err
}
