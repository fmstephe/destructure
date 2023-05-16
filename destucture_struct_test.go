package destructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type FirstIntStruct struct {
	Field1 int
}

type SecondIntStruct struct {
	Field2 int
}

type TwoIntStruct struct {
	Field1 int
	Field2 int
}

func TestStruct_HiddenBehindAny(t *testing.T) {
	in := any(FirstIntStruct{
		Field1: 1,
	})
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, in.(FirstIntStruct), out)
}

func TestStruct_HiddenInMap(t *testing.T) {
	in := map[any]any{
		"foo": FirstIntStruct{1},
		"bar": FirstIntStruct{2},
	}
	out := Destructure[map[string]FirstIntStruct](in)
	require.Equal(t, map[string]FirstIntStruct{
		"foo": FirstIntStruct{1},
		"bar": FirstIntStruct{2},
	}, out)
}

func TestStruct_FieldSubset_Field1(t *testing.T) {
	in := TwoIntStruct{
		Field1: 1,
		Field2: 2,
	}
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, FirstIntStruct{
		Field1: 1,
	}, out)
}

func TestStruct_FieldSubset_Field2(t *testing.T) {
	in := TwoIntStruct{
		Field1: 1,
		Field2: 2,
	}
	out := Destructure[SecondIntStruct](in)
	require.Equal(t, SecondIntStruct{
		Field2: 2,
	}, out)
}
