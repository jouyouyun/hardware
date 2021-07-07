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
