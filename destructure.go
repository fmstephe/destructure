package destructure

import (
	"fmt"
	"reflect"
)

func Destructure[T any](value any) T {
	//vVal := reflect.ValueOf(v)
	newT := (*T)(nil)
	// Grab the reflective type of T
	tType := reflect.TypeOf(newT).Elem()
	d, ok := destructure(reflect.ValueOf(value), tType)
	if !ok {
		return nilOf[T]()
	}

	return d.Interface().(T)
}

var zeroValue = reflect.Value{}

func destructure(value reflect.Value, structure reflect.Type) (reflect.Value, bool) {
	if value.Kind() == reflect.Interface {
		// Remove the layer of indirection for this interface value
		if value.IsNil() {
			return reflect.Value{}, false
		}
		value = value.Elem()
	}

	if value.Type().AssignableTo(structure) {
		newVal := reflect.New(structure).Elem()
		newVal.Set(value)
		return newVal, true
	}

	ifcValue := value.Interface()
	// Switch on the next expected kind
	switch structure.Kind() {
	case reflect.String:
		v, ok := ifcValue.(string)
		return reflect.ValueOf(v), ok

	case reflect.Bool:
		v, ok := ifcValue.(bool)
		return reflect.ValueOf(v), ok

	case reflect.Int:
		v, ok := ifcValue.(int)
		return reflect.ValueOf(v), ok

	case reflect.Int8:
		v, ok := ifcValue.(int8)
		return reflect.ValueOf(v), ok

	case reflect.Int16:
		v, ok := ifcValue.(int16)
		return reflect.ValueOf(v), ok

	case reflect.Int32:
		v, ok := ifcValue.(int32)
		return reflect.ValueOf(v), ok

	case reflect.Int64:
		v, ok := ifcValue.(int64)
		return reflect.ValueOf(v), ok

	case reflect.Uint:
		v, ok := ifcValue.(uint)
		return reflect.ValueOf(v), ok

	case reflect.Uint8:
		v, ok := ifcValue.(uint8)
		return reflect.ValueOf(v), ok

	case reflect.Uint16:
		v, ok := ifcValue.(uint16)
		return reflect.ValueOf(v), ok

	case reflect.Uint32:
		v, ok := ifcValue.(uint32)
		return reflect.ValueOf(v), ok

	case reflect.Uint64:
		v, ok := ifcValue.(uint64)
		return reflect.ValueOf(v), ok

	case reflect.Float32:
		v, ok := ifcValue.(float32)
		return reflect.ValueOf(v), ok

	case reflect.Float64:
		v, ok := ifcValue.(float64)
		return reflect.ValueOf(v), ok

	case reflect.Complex64:
		v, ok := ifcValue.(complex64)
		return reflect.ValueOf(v), ok

	case reflect.Complex128:
		v, ok := ifcValue.(complex128)
		return reflect.ValueOf(v), ok

	case reflect.Array:
		return copySliceArray(value, structure)

	case reflect.Slice:
		return copySliceArray(value, structure)

	case reflect.Map:
		return copyMap(value, structure)

	default:
		fmt.Printf("Unsupported kind %s\n", structure.Kind())
		return reflect.Value{}, false
	}
}

func copySliceArray(value reflect.Value, structure reflect.Type) (reflect.Value, bool) {
	if structure.Kind() != reflect.Array && structure.Kind() != reflect.Slice {
		panic(fmt.Errorf("requires structure type of either slice or array, got %s", structure.Kind()))
	}

	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return reflect.Value{}, false
	}

	retypedValues := []reflect.Value{}
	structureElem := structure.Elem()

	for i := 0; i < value.Len(); i++ {
		valueItem := value.Index(i)
		//if valueItem.IsZero() {
		//	continue
		//}
		if retypedItem, ok := destructure(valueItem, structureElem); ok {
			retypedValues = append(retypedValues, retypedItem)
		}
	}

	newArray := reflect.New(structure).Elem()

	// If we are copying a slice - grow it to the size needed for copy
	if newArray.Kind() == reflect.Slice {
		newArray = reflect.MakeSlice(structure, len(retypedValues), len(retypedValues))
	}

	for i, retypedValue := range retypedValues {
		if i >= newArray.Len() {
			break
		}
		newItem := newArray.Index(i)
		newItem.Set(retypedValue)
	}

	return newArray, true
}

func copyMap(value reflect.Value, structure reflect.Type) (reflect.Value, bool) {
	if structure.Kind() != reflect.Map {
		panic(fmt.Errorf("requires structure type of either slice or array, got %s", structure.Kind()))
	}

	if value.Kind() != reflect.Map {
		return reflect.Value{}, false
	}

	keyType := structure.Key()
	elemType := structure.Elem()

	newMap := reflect.MakeMap(structure)
	iter := value.MapRange()

	for iter.Next() {
		mapKey, ok := destructure(iter.Key(), keyType)
		if !ok {
			continue
		}
		mapVal, ok := destructure(iter.Value(), elemType)
		if !ok {
			continue
		}

		newMap.SetMapIndex(mapKey, mapVal)
	}

	return newMap, true
}

func nilOf[T any]() T {
	newT := (*T)(nil)
	// Grab the reflective type of T
	tType := reflect.TypeOf(newT).Elem()
	return reflect.New(tType).Elem().Interface().(T)
}
