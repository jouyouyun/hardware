package network

import (
	"github.com/jouyouyun/hardware/utils"
	"testing"
)

func Test_NewNetwork(t *testing.T) {
	net, err := newNetwork("./testdata", "enp3s0")
	if err != nil {
		t.Fatal(err.Error())
	}
	n := Network{
		utils.CardInfo{
			Name:    "Realtek Semiconductor Co., Ltd. RTL8111/8168/8411 PCI Express Gigabit Ethernet Controller",
			Vendor:  "10EC",
			Product: "8168",
			Slot:    "pci",
		},
		"00:e0:70:c4:66:c1",
		"10.20.22.135",
	}
	if net.Name != n.Name {
		t.Errorf("network name test failed: excepted %s, but got %s", n.Name, net.Name)
	}
	if net.Vendor != n.Vendor {
		t.Errorf("network vendor test failed: excepted %s, but got %s", n.Vendor, net.Vendor)
	}
	if net.Product != n.Product {
		t.Errorf("network product test failed: excepted %s, but got %s", n.Product, net.Product)
	}
	if net.Slot != n.Slot {
		t.Errorf("network slot test failed: excepted %s, but got %s", n.Slot, net.Slot)
	}
	if net.Address != n.Address {
		t.Errorf("network address test failed: excepted %s, but got %s", n.Address, net.Address)
	}
	if net.IP != n.IP {
		t.Errorf("network ip test failed: excepted %s, but got %s", n.IP, net.IP)
	}

}
