package television

import "fmt"

type Samsung struct {
	CurrentChan   int
	CurrentVolume int
	TvOn          bool
}

func (tv *Samsung) GetVolume() int {
	fmt.Println("Samsung volume is", tv.CurrentVolume)
	return tv.CurrentVolume
}

func (tv *Samsung) SetVolume(vol int) {
	fmt.Println("Setting Samsung volume to", vol)
	tv.CurrentVolume = vol
	fmt.Println("Samsung vol after change", tv.CurrentVolume)
}

func (tv *Samsung) GetCurrentChan() int {
	fmt.Println("Samsung current Channel is", tv.CurrentChan)
	return tv.CurrentChan
}

func (tv *Samsung) SetCurrentChan(ch int) {
	fmt.Println("Setting Samsung current Channel  to", ch)
	tv.CurrentChan = ch
}

func (tv *Samsung) SetOnState(tvOn bool) {
	if tvOn {
		fmt.Println("Samsung tv is on")
	} else {
		fmt.Println("Samsung tv is off")
	}

	tv.TvOn = tvOn
}
