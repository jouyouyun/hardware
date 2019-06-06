package network

// #cgo CFLAGS: -Wall -g
// #include "ip.h"
// #include <stdlib.h>
import "C"

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"unsafe"

	"github.com/jouyouyun/hardware/utils"
)

const (
	netSysfsDir   = "/sys/class/net"
	netVirtualDir = "/sys/devices/virtual/net"

	SlotTypePCI = "pci"
	SlotTypeUSB = "usb"
)

// Network store network info
type Network struct {
	Name    string
	Vendor  string
	Product string
	Address string
	IP      string

	Slot string
}

// NetworkList network list
type NetworkList []*Network

func newNetwork(dir, iface string) (*Network, error) {
	uevent := filepath.Join(dir, iface, "device", "uevent")
	uinfo, err := utils.NewUEvent(uevent)
	if err != nil {
		return nil, err
	}

	var net = Network{Name: uinfo.Name}
	if uinfo.Type == utils.UEventTypePCI {
		pci := uinfo.Data.(*utils.PCIUEvent)
		net.Slot = SlotTypePCI
		net.Vendor = pci.Vendor.Name
		net.Product = pci.Device.Name
	} else {
		usb := uinfo.Data.(*utils.USBUEvent)
		net.Slot = SlotTypeUSB
		net.Vendor = usb.Vendor
		net.Product = usb.Product
	}

	net.Address, _ = utils.ReadFileContent(filepath.Join(dir, iface, "address"))
	net.IP = getIfaceIP(iface)
	return &net, nil
}

func getIfaceList(dir string, filter func(string) bool) []string {
	finfos, _ := ioutil.ReadDir(dir)
	var ifaceList []string
	for _, finfo := range finfos {
		if filter(finfo.Name()) {
			continue
		}
		ifaceList = append(ifaceList, finfo.Name())
	}
	return ifaceList
}

func getIfaceIP(iface string) string {
	ciface := C.CString(iface)
	defer C.free(unsafe.Pointer(ciface))

	cret := C.get_iface_ip(ciface)
	defer C.free(unsafe.Pointer(cret))

	ret := C.GoString(cret)
	return ret
}

func filterIface(iface string) bool {
	return isVirtualIface(iface, netVirtualDir)
}

func isVirtualIface(iface, dir string) bool {
	_, err := os.Stat(filepath.Join(dir, iface))
	return err == nil || os.IsExist(err)
}
