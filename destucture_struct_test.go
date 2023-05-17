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

type FirstAnyStruct struct {
	Field1 any
}

type SecondAnyStruct struct {
	Field2 any
}

type FirstIntPointerStruct struct {
	Field1 *int
}

type SecondIntPointerStruct struct {
	Field2 *int
}

type TwoIntPointerStruct struct {
	Field1 *int
	Field2 *int
}

type TwoAnyStruct struct {
	Field1 any
	Field2 any
}

type DeeplyIndirectedAny struct {
	Field1 *any
}

type DeeplyIndirectedInt struct {
	Field1 *int
}

type BigInStruct struct {
	StringField    *string
	IntField       *int64
	UintField      *uint64
	BoolField      *bool
	FloatField     *float64
	ComplexField   *complex128
	ArrayField     [2]*BigInStruct
	SliceField     []*BigInStruct
	MapField       map[any]any
	StructField    *BigInStruct
	InterfaceField any
}

type BigOutStruct struct {
	StringField    string
	IntField       int32
	UintField      uint32
	BoolField      bool
	FloatField     float32
	ComplexField   complex64
	ArrayField     [2]*BigOutStruct
	SliceField     []BigOutStruct
	MapField       map[int32]BigOutStruct
	StructField    *BigOutStruct
	InterfaceField *BigOutStruct
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

func TestStruct_AnyFieldSubset_Field1(t *testing.T) {
	in := TwoAnyStruct{
		Field1: 1,
		Field2: 2,
	}
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, FirstIntStruct{
		Field1: 1,
	}, out)
}

func TestStruct_AnyFieldSubset_Field2(t *testing.T) {
	in := TwoAnyStruct{
		Field1: 1,
		Field2: 2,
	}
	out := Destructure[SecondIntStruct](in)
	require.Equal(t, SecondIntStruct{
		Field2: 2,
	}, out)
}

func TestStruct_PointerField_ToPointerField_HiddenByInterface(t *testing.T) {
	i := 1
	in := FirstAnyStruct{
		Field1: &i,
	}
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, FirstIntStruct{
		Field1: 1,
	}, out)
}

func TestStruct_PointerField_ToValueField(t *testing.T) {
	i := 1
	in := FirstIntPointerStruct{
		Field1: &i,
	}
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, FirstIntStruct{
		Field1: 1,
	}, out)
}

func TestStruct_DeeplyIndirectedAny_To_FirstIntStruct(t *testing.T) {
	i := any(1)
	in := DeeplyIndirectedAny{
		Field1: &i,
	}
	out := Destructure[FirstIntStruct](in)
	require.Equal(t, FirstIntStruct{
		Field1: 1,
	}, out)
}

func TestStruct_DeeplyIndirectedAny_To_DeeplyIndirectedInt(t *testing.T) {
	iAny := any(1)
	in := DeeplyIndirectedAny{
		Field1: &iAny,
	}
	out := Destructure[DeeplyIndirectedInt](in)
	i := 1
	require.Equal(t, DeeplyIndirectedInt{
		Field1: &i,
	}, out)
}

func TestStruct_BigInStruct_To_BigOutStruct(t *testing.T) {
	stringField := "string"
	intField := int64(1)
	uintField := uint64(2)
	boolField := true
	floatField := 3.333
	complexField := complex128(4.444)
	in := buildBigInStruct(bigInStructFun(stringField, intField, uintField, boolField, floatField, complexField))
	expectedOut := buildBigOutStruct(bigOutStructFun(stringField, intField, uintField, boolField, floatField, complexField))
	require.Equal(t, expectedOut, Destructure[BigOutStruct](in))
}

func bigInStructFun(
	stringField string,
	intField int64,
	uintField uint64,
	boolField bool,
	floatField float64,
	complexField complex128,
) func() *BigInStruct {
	return func() *BigInStruct {
		return &BigInStruct{
			StringField:  &stringField,
			IntField:     &intField,
			UintField:    &uintField,
			BoolField:    &boolField,
			FloatField:   &floatField,
			ComplexField: &complexField,
		}
	}
}

func bigOutStructFun(
	stringField string,
	intField int64,
	uintField uint64,
	boolField bool,
	floatField float64,
	complexField complex128,
) func() *BigOutStruct {
	return func() *BigOutStruct {
		return &BigOutStruct{
			StringField:  stringField,
			IntField:     int32(intField),
			UintField:    uint32(uintField),
			BoolField:    boolField,
			FloatField:   float32(floatField),
			ComplexField: complex64(complexField),
			SliceField:   []BigOutStruct{},
			MapField:     map[int32]BigOutStruct{},
		}
	}
}

func buildBigInStruct(
	structFun func() *BigInStruct,
) *BigInStruct {
	strct := structFun()
	strct.ArrayField = [2]*BigInStruct{
		structFun(),
		structFun(),
	}
	strct.SliceField = []*BigInStruct{
		structFun(),
		structFun(),
	}
	strct.MapField = map[any]any{
		1: structFun(),
		2: structFun(),
	}
	strct.StructField = structFun()
	strct.InterfaceField = structFun()
	return strct
}

func buildBigOutStruct(
	structFun func() *BigOutStruct,
) BigOutStruct {
	strct := *structFun()
	strct.ArrayField = [2]*BigOutStruct{
		structFun(),
		structFun(),
	}
	strct.SliceField = []BigOutStruct{
		*structFun(),
		*structFun(),
	}
	strct.MapField = map[int32]BigOutStruct{
		1: *structFun(),
		2: *structFun(),
	}
	strct.StructField = structFun()
	strct.InterfaceField = structFun()
	return strct
}
