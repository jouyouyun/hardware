package disk

import (
	"testing"
)

func TestGetRootMountInfo(t *testing.T) {
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
		   {"name":"zzzz", "serial":"10c34x45c45155d024f55sd4a4ba0001", "type":"rom", "size":2226057216, "vendor":"NECVMWar", "model":"VMware_Virtual_IDE_CDROM_Drive", "mountpoint":"/", "uuid":"2020-01-14-08-15-26-00"},
		   {"name":"adc", "serial":"10c34x45c45155d024f55sd4a4ba0001", "type":"rom", "size":2226057216, "vendor":"NECVMWar", "model":"VMware_Virtual_IDE_CDROM_Drive", "mountpoint":"/media/kyrie/uos 20", "uuid":"2020-01-14-08-15-26-00"}
		]
	 }`
	disk, err := newDiskListFromOutput([]byte(str))
	if err != nil {
		t.Error("json format error")
	}
	var infos = DiskList{
		&Disk{
			Name:        "sdb",
			Model:       "VMware_Virtual_S",
			Serial:      "68f9ds8fd7f8d52c8a2dj78fg79ss9c",
			Vendor:      "VMware, ",
			Size:        15548554655,
			RootMounted: true,
		},
		&Disk{
			Name:        "zzzz",
			Model:       "VMware_Virtual_IDE_CDROM_Drive",
			Serial:      "10c34x45c45155d024f55sd4a4ba0001",
			Vendor:      "NECVMWar",
			Size:        2226057216,
			RootMounted: true,
		},
		&Disk{
			Name:        "adc",
			Model:       "VMware_Virtual_IDE_CDROM_Drive",
			Serial:      "10c34x45c45155d024f55sd4a4ba0001",
			Vendor:      "NECVMWar",
			Size:        2226057216,
			RootMounted: false,
		},
	}
	for i, v := range disk {
		if v.Name != infos[i].Name {
			t.Errorf("Disk name test failed: excepted %s, but got %s", infos[i].Name, v.Name)
		}
		if v.Model != infos[i].Model {
			t.Errorf("Disk model test failed: excepted %s, but got %s", infos[i].Model, v.Model)
		}
		if v.Vendor != infos[i].Vendor {
			t.Errorf("Disk vendor test failed: excepted %s, but got %s", infos[i].Vendor, v.Vendor)
		}
		if v.Serial != infos[i].Serial {
			t.Errorf("Disk serial test failed: excepted %s, but got %s", infos[i].Serial, v.Serial)
		}
		if v.Size != infos[i].Size {
			t.Errorf("Disk size test failed: excepted %v, but got %v", infos[i].Size, v.Size)
		}
		if v.RootMounted != infos[i].RootMounted {
			t.Errorf("Disk root mounted test failed: excepted %v, but got %v", infos[i].RootMounted, v.RootMounted)
		}
	}
}
