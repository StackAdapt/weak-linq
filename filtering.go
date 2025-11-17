package weaklinq

//----------------------------------------------------------------------------//
// Filtering                                                                  //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

// FilterOnThis returns a new Iterable where the items are filtered by the given
// predicate.
func (iterable Iterable[T]) FilterOnThis(predicate func(T) bool) Iterable[T] {

	return Iterable[T]{
		Seq: func(yield func(T) bool) {
			iterable.Seq(func(item T) bool {
				if predicate(item) {
					return yield(item)
				}
				return true
			})
		},
	}

	/*
		linq.From([]T{...}).
			FilterOnThis(
				func(item T) bool {
					return boolExpression
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// FilterOn returns a new Iterable where the items are filtered by the given
// fieldName. If T is not a struct, or fieldName is not found, this function
// will panic.
func (iterable Iterable[T]) FilterOn(fieldName string) Iterable[T] {

	nameFunc := getFieldNameFunc[T](fieldName)

	return iterable.FilterOnThis(
		func(item T) bool {
			return nameFunc(item).(bool)
		},
	)

	/*
		linq.From([]T{...}).
			FilterOn("BoolFieldName")
	*/
}

////////////////////////////////////////////////////////////////////////////////

// Distinct returns a new Iterable where the items are distinct.
func (iterable Iterable[T]) Distinct() Iterable[T] {

	seen := make(map[any]bool)

	return Iterable[T]{
		Seq: func(yield func(T) bool) {
			iterable.Seq(func(item T) bool {
				if _, ok := seen[item]; !ok {
					seen[item] = true
					return yield(item)
				}
				return true
			})
		},
	}

	/*
		linq.From([]T{...}).
			Distinct()
	*/
}
