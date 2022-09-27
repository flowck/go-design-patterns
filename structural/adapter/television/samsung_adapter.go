package television

type SamsungAdapter struct {
	sstv *Samsung
}

func NewSamsungAdapter(sstv *Samsung) *SamsungAdapter {
	return &SamsungAdapter{sstv}
}

func (s *SamsungAdapter) VolumeUp() int {
	vol := s.sstv.GetVolume() + 1
	s.sstv.SetVolume(vol)
	return vol
}

func (s *SamsungAdapter) VolumeDown() int {
	vol := s.sstv.GetVolume() - 1
	s.sstv.SetVolume(vol)
	return vol
}

func (s *SamsungAdapter) ChannelUp() int {
	ch := s.sstv.GetCurrentChan() + 1
	s.sstv.SetCurrentChan(ch)
	return ch
}

func (s *SamsungAdapter) ChannelDown() int {
	ch := s.sstv.GetCurrentChan() - 1
	s.sstv.SetCurrentChan(ch)
	return ch
}

func (s *SamsungAdapter) TurnOn() {
	s.sstv.SetOnState(true)
}

func (s *SamsungAdapter) TurnOff() {
	s.sstv.TvOn = false
}

func (s *SamsungAdapter) GoToChannel(ch int) {
	s.sstv.SetCurrentChan(ch)
}
