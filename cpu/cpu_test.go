package cpu

import (
	"testing"
)

func Test_NewCPU(t *testing.T) {

	cpu, err := newCPU("./testdata/cpuinfo")
	if err != nil {
		t.Fatal("get cpu info error")
	}
	c1 := &CPU{
		Name:       "Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz",
		Processors: 8,
	}
	if cpu.Name != c1.Name {
		t.Errorf("cpu name test failed: excepted %s, but got %s", c1.Name, cpu.Name)
	}
	if cpu.Processors != c1.Processors {
		t.Errorf("cpu processor test failed: excepted %d, but got %d", c1.Processors, cpu.Processors)
	}

	swCpu, err := newCPUForSW("./testdata/swinfo")
	if err != nil {
		t.Fatal("get sw cpu info error")
	}
	c2 := &CPU{
		Name:       "sw",
		Processors: 4,
	}
	if swCpu.Name != c2.Name {
		t.Errorf("SW cpu name test failed: excepted %s, but got %s", c2.Name, swCpu.Name)
	}
	if swCpu.Processors != c2.Processors {
		t.Errorf("SW cpu processor test failed: excepted %d, but got %d", c2.Processors, swCpu.Processors)
	}

	armCpu, err := newCPUForARM("./testdata/arminfo")
	if err != nil {
		t.Fatal("get arm cpu info error")
	}
	c3 := &CPU{
		Name:       "ARMv7 Processor rev 0 (v7l)",
		Processors: 4,
	}
	if armCpu.Name != c3.Name {
		t.Errorf("ARM cpu name test failed: excepted %s, but got %s", c3.Name, armCpu.Name)
	}
	if armCpu.Processors != c3.Processors {
		t.Errorf("ARM cpu processor test failed: excepted %d, but got %d", c3.Processors, armCpu.Processors)
	}

	loonsonCpu, err := newCPUForLoonson("./testdata/loonsoninfo")
	if err != nil {
		t.Fatal("get loonson cpu info error")
	}
	c4 := &CPU{
		Name:       "Loongson-3B V0.7  FPU V0.1",
		Processors: 6,
	}
	if loonsonCpu.Name != c4.Name {
		t.Errorf("LOONSON cpu name test failed: excepted %s, but got %s", c3.Name, loonsonCpu.Name)
	}
	if loonsonCpu.Processors != c4.Processors {
		t.Errorf("LOONSON cpu processor test failed: excepted %d, but got %d", c3.Processors, loonsonCpu.Processors)
	}
}
