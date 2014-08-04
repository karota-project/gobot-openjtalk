package openjtalk

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestOpenjtalkDriver() *OpenjtalkDriver {
	return NewOpenjtalkDriver(NewOpenjtalkAdaptor("myAdaptor"), "myDriver")
}

func TestOpenjtalkDriverStart(t *testing.T) {
	d := initTestOpenjtalkDriver()
	gobot.Expect(t, d.Start(), true)
}

func TestOpenjtalkDriverHalt(t *testing.T) {
	d := initTestOpenjtalkDriver()
	gobot.Expect(t, d.Halt(), true)
}
