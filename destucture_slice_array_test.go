package destructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArray(t *testing.T) {
	in := [4]string{"1", "2", "3", "4"}
	out := Destructure[[4]string](in)
	require.Equal(t, in, out)
}

func TestArray_InterfaceToString(t *testing.T) {
	in := [4]any{"1", "2", "3", "4"}
	out := Destructure[[4]string](in)
	require.Equal(t, [4]string{"1", "2", "3", "4"}, out)
}

func TestArray_InterfaceToString_MixedTypes(t *testing.T) {
	in := [8]any{1, "1", 1.2, "2", true, "3", -1, "4"}
	out := Destructure[[4]string](in)
	require.Equal(t, [4]string{"1", "2", "3", "4"}, out)
}

func TestArray_InterfaceToString_MixedTypes_NotEnoughValues(t *testing.T) {
	in := [10]any{1, "1", 1.2, "2"}
	out := Destructure[[4]string](in)
	require.Equal(t, [4]string{"1", "2", "", ""}, out)
}

func TestArray_InterfaceToString_MixedTypes_TooManyValues(t *testing.T) {
	in := [8]any{1, "1", 1.2, "2", true, "3", -1, "4"}
	out := Destructure[[2]string](in)
	require.Equal(t, [2]string{"1", "2"}, out)
}

func TestNotArray(t *testing.T) {
	in := [4]int{1, 2, 3, 4}
	out := Destructure[[4]string](in)
	require.Equal(t, [4]string{}, out)
}

func TestSlice(t *testing.T) {
	in := []string{"1", "2", "3", "4"}
	out := Destructure[[]string](in)
	require.Equal(t, in, out)
}

func TestNotSlice(t *testing.T) {
	in := []int{1, 2, 3, 4}
	out := Destructure[[]string](in)
	require.Equal(t, []string{}, out)
}

func TestSlice_InterfaceToString(t *testing.T) {
	in := []any{"1", "2", "3", "4"}
	out := Destructure[[]string](in)
	require.Equal(t, []string{"1", "2", "3", "4"}, out)
}

func TestSlice_InterfaceToString_MixedTypes(t *testing.T) {
	in := []any{1, "1", 1.2, "2", true, "3", -1, "4", false, 1, 2, 3}
	out := Destructure[[]string](in)
	require.Equal(t, []string{"1", "2", "3", "4"}, out)
}
