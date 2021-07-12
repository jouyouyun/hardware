package network

import (
	"github.com/jouyouyun/hardware/utils"
	"testing"
)

func TestNewNetwork(t *testing.T) {
	info, err := newNetwork("./testdata", "jouyouyun_enp3s0")
	if err != nil {
		t.Error("failed to new network from testdata:", err)
		return
	}

	var v = Network{
		CardInfo: utils.CardInfo{
			Vendor:  "10EC",
			Product: "8168",
			Slot:    "pci",
		},
		Address: "00:e0:70:c4:66:c1",
		IP:      "0.0.0.0",
	}
	if info.Address != v.Address {
		t.Errorf("Address excepted %q, but %q", v.Address, info.Address)
	}
	if info.IP != v.IP {
		t.Errorf("IP excepted %q, but %q", v.IP, info.IP)
	}
	if info.Vendor != v.Vendor {
		t.Errorf("Vendor excepted %q, but %q", v.Vendor, info.Vendor)
	}
	if info.Product != v.Product {
		t.Errorf("Product excepted %q, but %q", v.Product, info.Product)
	}
	if info.Slot != v.Slot {
		t.Errorf("Slot excepted %q, but %q", v.Slot, info.Slot)
	}
}
