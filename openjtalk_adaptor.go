package openjtalk

import (
  "github.com/hybridgroup/gobot"
)

type OpenjtalkAdaptor struct {
  gobot.Adaptor
}

func NewOpenjtalkAdaptor(name string) *OpenjtalkAdaptor {
  return &OpenjtalkAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "openjtalk.OpenjtalkAdaptor",
    ),
  }
}

func (o *OpenjtalkAdaptor) Connect() bool {
  return true
}

func (o *OpenjtalkAdaptor) Finalize() bool {
  return true
}
