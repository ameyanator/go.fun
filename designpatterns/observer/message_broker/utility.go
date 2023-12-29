package main

import "reflect"

// Contains checks if a value is present in the given slice.
func Contains(slice interface{}, value interface{}) bool {
	sliceValue := reflect.ValueOf(slice)

	if sliceValue.Kind() != reflect.Slice {
		panic("Contains: non-slice type passed")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		if element == value {
			return true
		}
	}

	return false
}
