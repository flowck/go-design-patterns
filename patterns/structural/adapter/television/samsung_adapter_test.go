package television

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSamsungAdapter_VolumeUp(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	assert.Equal(t, tvAdapter.VolumeUp(), 56)
	assert.Equal(t, tv.CurrentVolume, 56)
}

func TestSamsungAdapter_VolumeDown(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	assert.Equal(t, tvAdapter.VolumeDown(), 54)
	assert.Equal(t, tv.CurrentVolume, 54)
}

func TestSamsungAdapter_ChannelUp(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	assert.Equal(t, tvAdapter.ChannelUp(), 3)
	assert.Equal(t, tv.CurrentChan, 3)
}

func TestSamsungAdapter_ChannelDown(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	assert.Equal(t, tvAdapter.ChannelDown(), 1)
	assert.Equal(t, tv.CurrentChan, 1)
}

func TestSamsungAdapter_TurnOn(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	tvAdapter.TurnOn()
	assert.Equal(t, tv.TvOn, true)
}

func TestSamsungAdapter_TurnOff(t *testing.T) {
	tv := &Samsung{
		CurrentChan:   2,
		CurrentVolume: 55,
		TvOn:          true,
	}

	tvAdapter := NewSamsungAdapter(tv)
	tvAdapter.TurnOff()
	assert.Equal(t, tv.TvOn, false)
}
