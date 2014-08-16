package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-openjtalk"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	openjtalkAdaptor := openjtalk.NewOpenjtalkAdaptor("openjtalk-a01")
	openjtalkDriver := openjtalk.NewOpenjtalkDriver(openjtalkAdaptor, "openjtalk-d01")

	master.AddRobot(
		gobot.NewRobot(
			"openjtalk",
			[]gobot.Connection{openjtalkAdaptor},
			[]gobot.Device{openjtalkDriver},
			func() {
				fmt.Println("work")
			}))

	master.Start()
}
