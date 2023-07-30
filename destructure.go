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

func destructure(value reflect.Value, structure reflect.Type) (reflect.Value, bool) {
	// unsafe-pointers are _always_ ignored
	if structure.Kind() == reflect.UnsafePointer {
		return reflect.Value{}, false
	}

	/*
		if value.Type().AssignableTo(structure) {
			newVal := reflect.New(structure).Elem()
			newVal.Set(value)
			return newVal, true
		}
	*/

	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return reflect.Value{}, false
		}
		// Remove indirection for pointer value
		return destructure(value.Elem(), structure)
	}

	if value.Kind() == reflect.Interface {
		if value.IsNil() {
			return reflect.Value{}, false
		}
		// Remove indirection for interface value
		return destructure(value.Elem(), structure)
	}

	// Switch on the next expected kind
	switch structure.Kind() {
	case reflect.String:
		// NB: I believe that with the `AssignableTo` check this will never produce a string value
		v, ok := value.Interface().(string)
		return reflect.ValueOf(v), ok

	case reflect.Bool:
		// NB: I believe that with the `AssignableTo` check this will never produce a bool value
		v, ok := value.Interface().(bool)
		return reflect.ValueOf(v), ok

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !value.CanInt() {
			return reflect.Value{}, false
		}
		newInt := reflect.New(structure).Elem()
		newInt.SetInt(value.Int())
		return newInt, true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if !value.CanUint() {
			return reflect.Value{}, false
		}
		newUint := reflect.New(structure).Elem()
		newUint.SetUint(value.Uint())
		return newUint, true

	case reflect.Float32, reflect.Float64:
		if !value.CanFloat() {
			return reflect.Value{}, false
		}
		newFloat := reflect.New(structure).Elem()
		newFloat.SetFloat(value.Float())
		return newFloat, true

	case reflect.Complex64, reflect.Complex128:
		if !value.CanComplex() {
			return reflect.Value{}, false
		}
		newComplex := reflect.New(structure).Elem()
		newComplex.SetComplex(value.Complex())
		return newComplex, true

	case reflect.Array, reflect.Slice:
		return copySliceArray(value, structure)

	case reflect.Map:
		return copyMap(value, structure)

	case reflect.Struct:
		return copyStruct(value, structure)

	case reflect.Chan:
		chanLen := 0
		if value.Kind() == reflect.Chan {
			chanLen = value.Len()
		}
		newChan := reflect.MakeChan(structure, chanLen)
		return newChan, true

	case reflect.Pointer:
		// TODO test me!
		structureElem := structure.Elem()
		newValue, ok := destructure(value, structureElem)
		if !ok {
			return reflect.Value{}, false
		}
		return newValue.Addr(), true

	case reflect.Interface:
		// TODO test me!
		newInterface := reflect.New(structure)
		structureElem := structure.Elem()
		newValue, ok := destructure(value, structureElem)
		if ok {
			newInterface.Elem().Set(newValue)
		}
		return newInterface, true

	case reflect.UnsafePointer:
		panic("not reachable")

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

	// If we are copying a slice - grow it to the size needed for copying
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

func copyStruct(value reflect.Value, structure reflect.Type) (reflect.Value, bool) {
	newStruct := reflect.New(structure).Elem()

	vType := value.Type()
	for i := 0; i < vType.NumField(); i++ {
		vF := vType.Field(i)
		if sF, ok := structure.FieldByName(vF.Name); ok {
			if fieldValue, ok := destructure(value.Field(i), sF.Type); ok {
				newField := newStruct.FieldByIndex(sF.Index)
				newField.Set(fieldValue)
			}
		}
	}

	return newStruct, true
}

func nilOf[T any]() T {
	newT := (*T)(nil)
	// Grab the reflective type of T
	tType := reflect.TypeOf(newT).Elem()
	return reflect.New(tType).Elem().Interface().(T)
}
