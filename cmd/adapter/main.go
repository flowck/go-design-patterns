package main

import "go-design-patterns/structural/adapter/television"

func main() {
	tv1 := &television.Samsung{CurrentChan: 13, CurrentVolume: 35, TvOn: true}
	tv2 := &television.Sony{Vol: 20, Channel: 9, IsOn: true}
	tvAdapted := television.NewSamsungAdapter(tv1)

	tv2.TurnOn()
	tv2.VolumeUp()
	tv2.VolumeDown()
	tv2.ChannelUp()
	tv2.ChannelDown()
	tv2.GoToChannel(68)
	tv2.TurnOff()

	tvAdapted.TurnOn()
	tvAdapted.VolumeUp()
	tvAdapted.VolumeDown()
	tvAdapted.ChannelUp()
	tvAdapted.ChannelDown()
	tvAdapted.GoToChannel(68)
	tvAdapted.TurnOff()
}
