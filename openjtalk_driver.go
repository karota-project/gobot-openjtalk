package openjtalk

import (
  "github.com/hybridgroup/gobot"
)

type OpenjtalkDriver struct {
  gobot.Driver
}

type OpenjtalkInterface interface {
}

func NewOpenjtalkDriver(a *OpenjtalkAdaptor, name string) *OpenjtalkDriver {
  return &OpenjtalkDriver{
    Driver: *gobot.NewDriver(
      name,
      "openjtalk.OpenjtalkDriver",
      a,
    ),
  }
}

func (o *OpenjtalkDriver) adaptor() *OpenjtalkAdaptor {
  return o.Driver.Adaptor().(*OpenjtalkAdaptor)
}

func (o *OpenjtalkDriver) Start() bool { return true }
func (o *OpenjtalkDriver) Halt() bool { return true }
