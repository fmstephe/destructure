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
	out := Destructure[string](true)
	require.Equal(t, "", out)
}

func TestBool(t *testing.T) {
	out := Destructure[bool](true)
	require.Equal(t, true, out)
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

func TestInt8_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, int8(5), Destructure[int8](5))
	require.Equal(t, int8(5), Destructure[int8](int16(5)))
	require.Equal(t, int8(5), Destructure[int8](int32(5)))
	require.Equal(t, int8(5), Destructure[int8](int64(5)))
}

func TestNotInt8(t *testing.T) {
	in := uint8(5)
	out := Destructure[int8](in)
	require.Equal(t, int8(0), out)
}

func TestInt16(t *testing.T) {
	in := int16(5)
	out := Destructure[int16](in)
	require.Equal(t, in, out)
}

func TestInt16_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, int16(5), Destructure[int16](5))
	require.Equal(t, int16(5), Destructure[int16](int8(5)))
	require.Equal(t, int16(5), Destructure[int16](int32(5)))
	require.Equal(t, int16(5), Destructure[int16](int64(5)))
}

func TestNotInt16(t *testing.T) {
	out := Destructure[int16](true)
	require.Equal(t, int16(0), out)
}

func TestInt32(t *testing.T) {
	in := int32(5)
	out := Destructure[int32](in)
	require.Equal(t, in, out)
}

func TestInt32_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, int32(5), Destructure[int32](5))
	require.Equal(t, int32(5), Destructure[int32](int8(5)))
	require.Equal(t, int32(5), Destructure[int32](int16(5)))
	require.Equal(t, int32(5), Destructure[int32](int64(5)))
}

func TestNotInt32(t *testing.T) {
	in := "actually a string"
	out := Destructure[int32](in)
	require.Equal(t, int32(0), out)
}

func TestInt64(t *testing.T) {
	in := int64(5)
	out := Destructure[int64](in)
	require.Equal(t, in, out)
}

func TestInt64_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, int64(5), Destructure[int64](5))
	require.Equal(t, int64(5), Destructure[int64](int8(5)))
	require.Equal(t, int64(5), Destructure[int64](int16(5)))
	require.Equal(t, int64(5), Destructure[int64](int32(5)))
}

func TestNotInt64(t *testing.T) {
	in := "actually a string"
	out := Destructure[int64](in)
	require.Equal(t, int64(0), out)
}

func TestUint(t *testing.T) {
	in := uint(5)
	out := Destructure[uint](in)
	require.Equal(t, in, out)
}

func TestUint_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, uint(5), Destructure[uint](uint8(5)))
	require.Equal(t, uint(5), Destructure[uint](uint16(5)))
	require.Equal(t, uint(5), Destructure[uint](uint32(5)))
	require.Equal(t, uint(5), Destructure[uint](uint32(5)))
}

func TestNotUint(t *testing.T) {
	in := "actually a string"
	out := Destructure[uint](in)
	require.Equal(t, uint(0), out)
}

func TestUint8(t *testing.T) {
	in := uint8(5)
	out := Destructure[uint8](in)
	require.Equal(t, in, out)
}

func TestUint8_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, uint8(5), Destructure[uint8](uint8(5)))
	require.Equal(t, uint8(5), Destructure[uint8](uint16(5)))
	require.Equal(t, uint8(5), Destructure[uint8](uint32(5)))
	require.Equal(t, uint8(5), Destructure[uint8](uint64(5)))
}

func TestNotUint8(t *testing.T) {
	in := "actually a string"
	out := Destructure[uint8](in)
	require.Equal(t, uint8(0), out)
}

func TestUint16(t *testing.T) {
	in := uint16(5)
	out := Destructure[uint16](in)
	require.Equal(t, in, out)
}

func TestUint16_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, uint16(5), Destructure[uint16](uint8(5)))
	require.Equal(t, uint16(5), Destructure[uint16](uint16(5)))
	require.Equal(t, uint16(5), Destructure[uint16](uint32(5)))
	require.Equal(t, uint16(5), Destructure[uint16](uint64(5)))
}

func TestNotUint16(t *testing.T) {
	in := "actually a string"
	out := Destructure[uint16](in)
	require.Equal(t, uint16(0), out)
}

func TestUint32(t *testing.T) {
	in := uint32(5)
	out := Destructure[uint32](in)
	require.Equal(t, in, out)
}

func TestUint32_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, uint32(5), Destructure[uint32](uint8(5)))
	require.Equal(t, uint32(5), Destructure[uint32](uint16(5)))
	require.Equal(t, uint32(5), Destructure[uint32](uint32(5)))
	require.Equal(t, uint32(5), Destructure[uint32](uint64(5)))
}

func TestNotUint32(t *testing.T) {
	in := "actually a string"
	out := Destructure[uint32](in)
	require.Equal(t, uint32(0), out)
}

func TestUint64(t *testing.T) {
	in := uint64(5)
	out := Destructure[uint64](in)
	require.Equal(t, in, out)
}

func TestUint64_FromOtherIntSizes(t *testing.T) {
	require.Equal(t, uint64(5), Destructure[uint64](uint8(5)))
	require.Equal(t, uint64(5), Destructure[uint64](uint16(5)))
	require.Equal(t, uint64(5), Destructure[uint64](uint32(5)))
	require.Equal(t, uint64(5), Destructure[uint64](uint64(5)))
}

func TestNotUint64(t *testing.T) {
	in := "actually a string"
	out := Destructure[uint64](in)
	require.Equal(t, uint64(0), out)
}

func TestFloat32(t *testing.T) {
	in := float32(5)
	out := Destructure[float32](in)
	require.Equal(t, in, out)
}

func TestFloat32_FromOtherFloatSizes(t *testing.T) {
	require.Equal(t, float32(5.1), Destructure[float32](5.1))
}

func TestNotFloat32(t *testing.T) {
	in := "actually a string"
	out := Destructure[float32](in)
	require.Equal(t, float32(0), out)
}

func TestFloat64(t *testing.T) {
	in := float64(5)
	out := Destructure[float64](in)
	require.Equal(t, in, out)
}

func TestFloat64_FromOtherFloatSizes(t *testing.T) {
	require.Equal(t, float64(float32(5.1)), Destructure[float64](float32(5.1)))
}

func TestNotFloat64(t *testing.T) {
	in := "actually a string"
	out := Destructure[float64](in)
	require.Equal(t, float64(0), out)
}

func TestComplex64(t *testing.T) {
	in := complex64(5)
	out := Destructure[complex64](in)
	require.Equal(t, in, out)
}

func TestComplex64_FromOtherComplexSizes(t *testing.T) {
	// TODO how do we initialise complex values
	require.Equal(t, complex64(5.1), Destructure[complex64](complex128(5.1)))
}

func TestNotComplex64(t *testing.T) {
	in := "actually a string"
	out := Destructure[complex64](in)
	require.Equal(t, complex64(0), out)
}

func TestComplex128(t *testing.T) {
	in := complex128(5)
	out := Destructure[complex128](in)
	require.Equal(t, in, out)
}

func TestComplex128_FromOtherComplexSizes(t *testing.T) {
	// TODO how do we initialise complex values
	require.Equal(t, complex128(complex64(5.1)), Destructure[complex128](complex64(5.1)))
}

func TestNotComplex128(t *testing.T) {
	in := "actually a string"
	out := Destructure[complex128](in)
	require.Equal(t, complex128(0), out)
}

