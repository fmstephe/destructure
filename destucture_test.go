package destructure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	in := "foo"
	out := Destructure[string](in)
	require.Equal(t, in, out)
}

func TestNotString(t *testing.T) {
	in := true
	out := Destructure[string](in)
	require.Equal(t, "", out)
}

func TestBool(t *testing.T) {
	in := true
	out := Destructure[bool](in)
	require.Equal(t, in, out)
}

func TestNotBool(t *testing.T) {
	in := "foo"
	out := Destructure[bool](in)
	require.Equal(t, false, out)
}

func TestInt(t *testing.T) {
	in := 123
	out := Destructure[int](in)
	require.Equal(t, in, out)
}

func TestNotInt(t *testing.T) {
	in := "foo"
	out := Destructure[int](in)
	require.Equal(t, 0, out)
}

func TestInt8(t *testing.T) {
	in := int8(5)
	out := Destructure[int8](in)
	require.Equal(t, in, out)
}

func TestNotInt8(t *testing.T) {
	in := true
	out := Destructure[int8](in)
	require.Equal(t, int8(0), out)
}

func TestInt16(t *testing.T) {
	in := int16(5)
	out := Destructure[int16](in)
	require.Equal(t, in, out)
}

func TestNotInt16(t *testing.T) {
	in := true
	out := Destructure[int16](in)
	require.Equal(t, int16(0), out)
}

func TestInt32(t *testing.T) {
	in := int32(5)
	out := Destructure[int32](in)
	require.Equal(t, in, out)
}

func TestNotInt32(t *testing.T) {
	in := true
	out := Destructure[int32](in)
	require.Equal(t, int32(0), out)
}

func TestInt64(t *testing.T) {
	in := int64(5)
	out := Destructure[int64](in)
	require.Equal(t, in, out)
}

func TestNotInt64(t *testing.T) {
	in := true
	out := Destructure[int64](in)
	require.Equal(t, int64(0), out)
}

func TestUint(t *testing.T) {
	in := uint(5)
	out := Destructure[uint](in)
	require.Equal(t, in, out)
}

func TestNotUint(t *testing.T) {
	in := true
	out := Destructure[uint](in)
	require.Equal(t, uint(0), out)
}

func TestUint8(t *testing.T) {
	in := uint8(5)
	out := Destructure[uint8](in)
	require.Equal(t, in, out)
}

func TestNotUint8(t *testing.T) {
	in := true
	out := Destructure[uint8](in)
	require.Equal(t, uint8(0), out)
}

func TestUint16(t *testing.T) {
	in := uint16(5)
	out := Destructure[uint16](in)
	require.Equal(t, in, out)
}

func TestNotUint16(t *testing.T) {
	in := true
	out := Destructure[uint16](in)
	require.Equal(t, uint16(0), out)
}

func TestUint32(t *testing.T) {
	in := uint32(5)
	out := Destructure[uint32](in)
	require.Equal(t, in, out)
}

func TestNotUint32(t *testing.T) {
	in := true
	out := Destructure[uint32](in)
	require.Equal(t, uint32(0), out)
}

func TestUint64(t *testing.T) {
	in := uint64(5)
	out := Destructure[uint64](in)
	require.Equal(t, in, out)
}

func TestNotUint64(t *testing.T) {
	in := true
	out := Destructure[uint64](in)
	require.Equal(t, uint64(0), out)
}

func TestFloat32(t *testing.T) {
	in := float32(5)
	out := Destructure[float32](in)
	require.Equal(t, in, out)
}

func TestNotFloat32(t *testing.T) {
	in := true
	out := Destructure[float32](in)
	require.Equal(t, float32(0), out)
}

func TestFloat64(t *testing.T) {
	in := float64(5)
	out := Destructure[float64](in)
	require.Equal(t, in, out)
}

func TestNotFloat64(t *testing.T) {
	in := true
	out := Destructure[float64](in)
	require.Equal(t, float64(0), out)
}

func TestComplex64(t *testing.T) {
	in := complex64(5)
	out := Destructure[complex64](in)
	require.Equal(t, in, out)
}

func TestNotComplex64(t *testing.T) {
	in := true
	out := Destructure[complex64](in)
	require.Equal(t, complex64(0), out)
}

func TestComplex128(t *testing.T) {
	in := complex128(5)
	out := Destructure[complex128](in)
	require.Equal(t, in, out)
}

func TestNotComplex128(t *testing.T) {
	in := true
	out := Destructure[complex128](in)
	require.Equal(t, complex128(0), out)
}

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

type OneIntStruct struct {
	Field int
}

type TwoIntStruct struct {
	Field1 int
	Field2 int
}

func TestStruct_HiddenBehindAny(t *testing.T) {
	in := any(OneIntStruct{
		Field: 1,
	})
	out := Destructure[OneIntStruct](in)
	require.Equal(t, in.(OneIntStruct), out)
}

func TestStruct_HiddenInMap(t *testing.T) {
	in := map[any]any{
		"foo": OneIntStruct{1},
		"bar": OneIntStruct{2},
	}
	out := Destructure[map[string]OneIntStruct](in)
	require.Equal(t, map[string]OneIntStruct{
		"foo": OneIntStruct{1},
		"bar": OneIntStruct{2},
	}, out)
}
