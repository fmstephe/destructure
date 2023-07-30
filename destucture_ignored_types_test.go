package destructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)


type ChanStruct struct {
	ChanField chan int
}

// Here is an example of a channel being ignored
func TestChan_ChanIsIgnored(t *testing.T) {
	require.Nil(t, Destructure[chan int](make(chan any)))
	require.Nil(t, Destructure[chan any](make(chan bool)))
	require.Nil(t, Destructure[chan int](1))
}

func TestChan_FieldIgnored(t *testing.T) {
	in := ChanStruct{
		ChanField: make(chan int),
	}
	out := Destructure[ChanStruct](in)
	require.Equal(t, in, out)
}

