package television

import "fmt"

type Sony struct {
	Vol     int
	Channel int
	IsOn    bool
}

func (tv *Sony) TurnOn() {
	fmt.Println("Sony tv is on")
	tv.IsOn = true
}

func (tv *Sony) TurnOff() {
	fmt.Println("Sony tv is off")
	tv.IsOn = false
}

func (tv *Sony) VolumeUp() int {
	tv.Vol++
	fmt.Println("Increasing Sony tv volume to", tv.Vol)
	return tv.Vol
}

func (tv *Sony) VolumeDown() int {
	if tv.Vol > 0 {
		tv.Vol--
	}

	fmt.Println("Decreasing Sony tv volume to", tv.Vol)
	return tv.Vol
}

func (tv *Sony) ChannelUp() int {
	tv.Channel++
	fmt.Println("Increasing Sony tv Channel to", tv.Vol)
	return tv.Channel
}

func (tv *Sony) ChannelDown() int {
	if tv.Channel > 0 {
		tv.Channel--
	}

	fmt.Println("Decreasing Sony tv Channel to", tv.Channel)
	return tv.Channel
}

func (tv *Sony) GoToChannel(ch int) int {
	if tv.Channel >= 0 {
		fmt.Println("Sony tv Channel is set to", ch)
		tv.Channel = ch
		return tv.Channel
	}

	return 0
}
