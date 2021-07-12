package hardware

import (
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/jouyouyun/hardware/cpu"
	"github.com/jouyouyun/hardware/disk"
	hdisk "github.com/jouyouyun/hardware/disk"
	hdmi "github.com/jouyouyun/hardware/dmi"
	"github.com/jouyouyun/hardware/network"
	"github.com/jouyouyun/hardware/utils"
)

const (
	etcMachineIDFile = "/etc/machine-id"
)

var (
	_mid            string
	IncludeDiskInfo bool
)

// GenMachineID generate this machine's id
func GenMachineID() (string, error) {
	if len(_mid) != 0 {
		return _mid, nil
	}

	var (
		mid   string
		err   error
		dmi   *hdmi.DMI
		disks disk.DiskList
		root  *disk.Disk
	)

	dmi, err = hdmi.GetDMI()
	if !IncludeDiskInfo && err == nil && dmi.IsValid() {
		mid, err = genMachineIDWithDMI(*dmi)
		if len(mid) != 0 {
			goto out
		}
	} else if dmi == nil {
		dmi = &hdmi.DMI{}
	}

	// if dmi product uuid null or IncludeDiskInfo is true, generate machine id with root disk serial
	disks, err = hdisk.GetDiskList()
	if err == nil {
		root = disks.GetRoot()
		if root != nil && len(root.Serial) != 0 {
			mid, err = genMachineIDWithDisk(*dmi, root)
			if len(mid) != 0 {
				goto out
			}
		}
	}

	mid, err = genMachineIDWithNet(*dmi)
	if len(mid) != 0 {
		goto out
	}

	mid, err = genMachineIDWithFile(etcMachineIDFile)

out:
	_mid = mid
	return mid, err
}

func genMachineIDWithDMI(dmi hdmi.DMI) (string, error) {
	// bios info maybe changed after upgraded
	dmi.BiosDate = ""
	dmi.BiosVendor = ""
	dmi.BiosVersion = ""
	return doGenMachineID(&dmi)
}

func genMachineIDWithDisk(dmi hdmi.DMI, disk *hdisk.Disk) (string, error) {
	// bios info maybe changed after upgraded
	dmi.BiosDate = ""
	dmi.BiosVendor = ""
	dmi.BiosVersion = ""
	var info = struct {
		hdmi.DMI
		DiskSerial string
	}{
		DMI:        dmi,
		DiskSerial: disk.Serial,
	}
	return doGenMachineID(&info)
}

func doGenMachineID(info interface{}) (string, error) {
	data, err := json.Marshal(info)
	if err != nil {
		return "", err
	}
	return utils.SHA256Sum(data), nil
}

func genMachineIDWithNet(dmi hdmi.DMI) (string, error) {
	cpuInfo, err := cpu.NewCPU()
	if err != nil {
		return "", err
	}
	netInfo, err := network.GetNetworkList()
	if err != nil {
		return "", err
	}

	var (
		plist []string
		vlist []string
		alist []string
	)
	for _, info := range netInfo {
		plist = append(plist, info.Product)
		vlist = append(vlist, info.Vendor)
		alist = append(alist, info.Address)
	}
	sort.Strings(plist)
	sort.Strings(vlist)
	sort.Strings(alist)

	// bios info maybe changed after upgraded
	dmi.BiosDate = ""
	dmi.BiosVendor = ""
	dmi.BiosVersion = ""
	var info = struct {
		hdmi.DMI
		CPU         string
		ProductList []string
		VendorList  []string
		AddressList []string
	}{
		DMI:         dmi,
		CPU:         cpuInfo.Name,
		ProductList: plist,
		VendorList:  vlist,
		AddressList: alist,
	}
	return doGenMachineID(&info)
}

func genMachineIDWithFile(filename string) (string, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	var info = struct {
		ID string
	}{
		ID: string(contents),
	}
	return doGenMachineID(&info)
}
