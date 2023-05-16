package destructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_StringAny_To_StringInt(t *testing.T) {
	in := map[string]any{
		"foo": 1,
		"bar": 2,
	}
	out := Destructure[map[string]int](in)
	require.Equal(t, map[string]int{
		"foo": 1,
		"bar": 2,
	}, out)
}

func TestMap_AnyAny_To_StringInt(t *testing.T) {
	in := map[any]any{
		"foo": 1,
		"bar": 2,
	}
	out := Destructure[map[string]int](in)
	require.Equal(t, map[string]int{
		"foo": 1,
		"bar": 2,
	}, out)
}

func TestMap_AnyAnyAny_To_StringInt(t *testing.T) {
	in := any(map[any]any{
		"foo": 1,
		"bar": 2,
	})
	out := Destructure[map[string]int](in)
	require.Equal(t, map[string]int{
		"foo": 1,
		"bar": 2,
	}, out)
}

func TestNotMap(t *testing.T) {
	in := []int{1, 2, 3, 4}
	out := Destructure[map[string]int](in)
	require.Equal(t, map[string]int(nil), out)
}
