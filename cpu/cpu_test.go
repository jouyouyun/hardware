package cpu

import (
	"testing"
)

func Test_NewCPU(t *testing.T) {

	cpu, err := newCPU(cpuFilename)
	println("cpu info:", cpu, ", err:", err)

	cpu, err = newCPUForSW(cpuFilename)
	println("cpu info:", cpu, ", err:", err)

	cpu, err = newCPUForLoonson(cpuFilename)
	println("cpu info:", cpu, ", err:", err)

	cpu, err = newCPUForARM(cpuFilename)
	println("cpu info:", cpu, ", err:", err)

}
