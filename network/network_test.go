package network

import (
	"strings"
	"testing"
)

func Test_GetNetworkList(t *testing.T) {
	nets, err := GetNetworkList()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, no := range nets {
		if strings.EqualFold(no.Name, "docker0") ||
			strings.EqualFold(no.Name, "lo") ||
			strings.EqualFold(no.Name, "virbr0") ||
			strings.EqualFold(no.Name, "virbr0-nic") ||
			strings.EqualFold(no.Name, "vmnet1") ||
			strings.EqualFold(no.Name, "vmnet8") {
			t.Errorf("error: contine virtual network")
			t.FailNow()
		}
		println(no.Name)
	}
}

func Test_GetNetworkListByCustomDir(t *testing.T) {
	SetDir("./testdata", "/sys/devices/virtual/net")
	nets, err := GetNetworkList()
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, no := range nets {
		if !strings.EqualFold(no.Name, "Realtek Semiconductor Co., Ltd. RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller") {
			t.Error("test local network name failed")
		}
		if !strings.EqualFold(no.Vendor, "10EC") {
			t.Error("test local network vendor failed")
		}
		if !strings.EqualFold(no.Product, "8168") {
			t.Error("test local network product failed")
		}
		if !strings.EqualFold(no.Slot, "pci") {
			t.Error("test local network slot failed")
		}
		if !strings.EqualFold(no.Address, "00:e0:70:c4:66:c1") {
			t.Error("test local network address failed")
		}
		if !strings.EqualFold(no.IP, "10.20.22.135") {
			t.Error("test local network ip failed")
		}
	}
}
