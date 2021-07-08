package cpu

import (
	"testing"
)

func Test_NewCPU(t *testing.T) {

	cpu, err := newCPU(cpuFilename)
	if err != nil {
		t.Error("get cpu info error")
	}

	cpu, err = newCPUForSW(cpuFilename)
	if err != nil {
		t.Error("get sw cpu info error")
	}

	cpu, err = newCPUForLoonson(cpuFilename)
	if err != nil {
		t.Error("get loonson cpu info error")
	}

	cpu, err = newCPUForARM(cpuFilename)
	if err != nil {
		t.Error("get arm cpu info error")
	}

	println(cpu)
}
