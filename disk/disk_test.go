package disk

import (
	"strings"
	"testing"
)

func Test_GetRootMountInfo(t *testing.T) {
	var str = `{
		"blockdevices": [
		   {"name":"sdb", "serial":"68f9ds8fd7f8d52c8a2dj78fg79ss9c", "type":"disk", "size":15548554655, "vendor":"VMware, ", "model":"VMware_Virtual_S", "mountpoint":null, "uuid":null,
			  "children": [
				 {"name":"sda1", "serial":null, "type":"part", "size":1610612736, "vendor":null, "model":null, "mountpoint":"/", "uuid":"c41673e5-638f-4f3c-b52d-cd1667e024b3"},
				 {"name":"sda3", "serial":null, "type":"part", "size":2147483648, "vendor":null, "model":null, "mountpoint":"[SWAP]", "uuid":"b8604489-15fc-40e8-bd69-74bad4045624",
			 		"children": [
				 		{"name":"sda4", "serial":null, "type":"part1", "size":85894356591, "vendor":null, "model":null, "mountpoint":"/bin", "uuid":"dd52f15b-876a-4cb1-8eec-013c974c568a"}
			  		]
			 	 }
			  ]
		   },
		   {"name":"zzzz", "serial":"10c34x45c45155d024f55sd4a4ba0001", "type":"rom", "size":2226057216, "vendor":"NECVMWar", "model":"VMware_Virtual_IDE_CDROM_Drive", "mountpoint":"/data/", "uuid":"2020-01-14-08-15-26-00"}
		]
	 }`
	info := []byte(str)
	disk, err := newDiskListFromOutput(info)
	if err != nil {
		t.Error("json format error")
	}
	for index, v := range disk {
		if index == 0 {
			if !strings.EqualFold(v.Name, "sdb") {
				t.Error("test disk name failed")
			}
			if !strings.EqualFold(v.Model, "VMware_Virtual_S") {
				t.Error("test disk model failed")
			}
			if !strings.EqualFold(v.Serial, "68f9ds8fd7f8d52c8a2dj78fg79ss9c") {
				t.Error("test disk serial failed")
			}
			if !strings.EqualFold(v.Vendor, "VMware, ") {
				t.Error("test disk vendor failed")
			}
			if v.Size != 15548554655 {
				t.Error("test disk size failed")
			}
			if !v.RootMounted {
				t.Error("test disk root mounted failed")
			}
		} else {
			if !strings.EqualFold(v.Name, "zzzz") {
				t.Error("test disk name failed")
			}
			if !strings.EqualFold(v.Model, "VMware_Virtual_IDE_CDROM_Drive") {
				t.Error("test disk model failed")
			}
			if !strings.EqualFold(v.Serial, "10c34x45c45155d024f55sd4a4ba0001") {
				t.Error("test disk serial failed")
			}
			if !strings.EqualFold(v.Vendor, "NECVMWar") {
				t.Error("test disk vendor failed")
			}
			if v.Size != 2226057216 {
				t.Error("test disk size failed")
			}
			if v.RootMounted {
				t.Error("test disk root mounted failed")
			}
		}
	}
}
