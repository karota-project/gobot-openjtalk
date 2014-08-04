package openjtalk

import (
  "github.com/hybridgroup/gobot"
  "testing"
)

func initTestOpenjtalkAdaptor() *OpenjtalkAdaptor {
  return NewOpenjtalkAdaptor("myAdaptor")
}

func TestOpenjtalkAdaptorConnect(t *testing.T) {
  a := initTestOpenjtalkAdaptor()
  gobot.Expect(t, a.Connect(), true)
}

func TestOpenjtalkAdaptorFinalize(t *testing.T) {
  a := initTestOpenjtalkAdaptor()
  gobot.Expect(t, a.Finalize(), true)
}
