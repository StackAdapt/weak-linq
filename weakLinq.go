package weaklinq

import (
	"fmt"
	"reflect"
	"slices"
)

//----------------------------------------------------------------------------//
// Common                                                                     //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func identitySelector[T any](item T) any {
	return item
}

////////////////////////////////////////////////////////////////////////////////

// getFieldNameFunc returns a function that returns the value of the given field
// name. If T is not a struct or pointer to struct, or fieldName is not found,
// this function will panic.
func getFieldNameFunc[T any](fieldName string) func(T) any {

	return func(item T) any {

		if reflect.TypeOf(item).Kind() != reflect.Struct && reflect.TypeOf(item).Kind() != reflect.Pointer {
			panic(fmt.Sprintf("item is not a struct or pointer to struct: %T", item))
		}

		var value reflect.Value
		if reflect.TypeOf(item).Kind() == reflect.Pointer {
			value = reflect.ValueOf(item).Elem()
		} else {
			value = reflect.ValueOf(item)
		}

		field := value.FieldByName(fieldName)
		if !field.IsValid() {
			panic(fmt.Sprintf("field name '%s' not found in struct %T", fieldName, item))
		}

		return field.Interface()
	}
}

//----------------------------------------------------------------------------//
// Constructors                                                               //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

// From creates a new Iterable from a slice of items.
func From[T any](items []T) Iterable[T] {

	return Iterable[T]{
		Seq: slices.Values(items),
	}

	/*
		linq.From([]T{...})
	*/
}
