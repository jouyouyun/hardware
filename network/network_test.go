package network

import "testing"

func Test_GetNetworkList(t *testing.T) {

	nets, err := GetNetworkList()

	println("networklist:", nets, ", error:", err)

}
