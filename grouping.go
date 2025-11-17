package weaklinq

//----------------------------------------------------------------------------//
// Grouping                                                                   //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

// MapIterable is a specialized iterable that includes a key selector and an
// item selector. These selectors are stored until the collection is iterated,
// and then applied to the items. The Overwrite flag determines whether the
// map value for a given key should be overwritten or appended to. If set to
// true, the result map will be expected to be a map[K]V, and if set to false,
// the result map will be expected to be a map[K][]V.
type MapIterable[T any] struct {
	itemIterable Iterable[T]
	keySelector  func(T) any
	itemSelector func(T) any
	overwrite    bool
}

/////////////////////////////////////////////////////////////////////////////////

// DeferredKeyMapIterable is a MapIterable where the key selector has not yet
// been set. Designed to be used in tandem with the By functions. Has very
// little use outside of that.
type DeferredKeyMapIterable[T any] MapIterable[T]

////////////////////////////////////////////////////////////////////////////////

// defaultMapIterable returns a default MapIterable for the given iterable where
// the key selector and item selector are the identity function.
func defaultMapIterable[T any](iterable Iterable[T]) MapIterable[T] {

	return MapIterable[T]{
		itemIterable: iterable,
		keySelector:  identitySelector[T],
		itemSelector: identitySelector[T],
		overwrite:    true,
	}
}

////////////////////////////////////////////////////////////////////////////////

// GroupThis returns a new MapIterable where the items are grouped by the given
// selector. Items with the same key WILL BE OVERWRITTEN as this iterable is
// iterated. Should be used in tandem with the By functions. Has very little
// use outside of that.
func (iterable Iterable[T]) GroupThis(selector func(T) any) DeferredKeyMapIterable[T] {

	mapIterable := defaultMapIterable(iterable)
	mapIterable.itemSelector = selector

	return DeferredKeyMapIterable[T](mapIterable)

	/*
		linq.From([]T{...}).
			GroupThis(
				func(item T) any {
					return item.ItemField
				},
			).By(
				func(item T) any {
					return item.KeyField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// Group returns a new MapIterable where the items are grouped by the given
// field name. Items with the same key WILL BE OVERWRITTEN as this iterable
// is iterated. Should be used in tandem with the By functions. Has very little
// use outside of that. If T is not a struct, or fieldName is not found, this
// function will panic.
func (iterable Iterable[T]) Group(fieldName string) DeferredKeyMapIterable[T] {

	return iterable.GroupThis(
		getFieldNameFunc[T](fieldName),
	)

	/*
		linq.From([]T{...}).
			Group("ItemField").
			By(
				func(item T) any {
					return item.KeyField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupListsOfThis returns a new MapIterable where the items are grouped by the
// given selector. Items with the same key will NOT be overwritten as this
// iterable is iterated. Should be used in tandem with the By functions. Has
// very little use outside of that.
func (iterable Iterable[T]) GroupListsOfThis(selector func(T) any) DeferredKeyMapIterable[T] {

	mapIterable := defaultMapIterable(iterable)
	mapIterable.itemSelector = selector
	mapIterable.overwrite = false

	return DeferredKeyMapIterable[T](mapIterable)

	/*
		linq.From([]T{...}).
			GroupListsOfThis(
				func(item T) any {
					return item.ItemField
				},
			).By(
				func(item T) any {
					return item.KeyField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupListsOf returns a new MapIterable where the items are grouped by the
// given field name. Items with the same key will NOT be overwritten as this
// iterable is iterated. Should be used in tandem with the By functions. Has
// very little use outside of that. If T is not a struct, or fieldName is not
// found, this function will panic.
func (iterable Iterable[T]) GroupListsOf(fieldName string) DeferredKeyMapIterable[T] {

	return iterable.GroupListsOfThis(
		getFieldNameFunc[T](fieldName),
	)

	/*
		linq.From([]T{...}).
			GroupListsOf("ItemField").
			By(
				func(item T) any {
					return item.KeyField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupByThis returns a new MapIterable where the items are grouped by the
// given key selector. Items with the same key WILL BE OVERWRITTEN as this
// iterable is iterated.
func (iterable Iterable[T]) GroupByThis(selector func(T) any) MapIterable[T] {

	mapIterable := defaultMapIterable(iterable)
	mapIterable.keySelector = selector

	return mapIterable

	/*
		linq.From([]T{...}).
			GroupByThis(
				func(item T) any {
					return item.ItemField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupBy returns a new MapIterable where the items are grouped by the given
// field name. Items with the same key WILL BE OVERWRITTEN as this iterable is
// iterated. If T is not a struct, or fieldName is not found, this function
// will panic.
func (iterable Iterable[T]) GroupBy(fieldName string) MapIterable[T] {

	return iterable.GroupByThis(
		getFieldNameFunc[T](fieldName),
	)

	/*
		linq.From([]T{...}).
			GroupBy("ItemField")
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupListsByThis returns a new MapIterable where the items are grouped by
// the given key selector. Items with the same key will NOT be overwritten as
// this iterable is iterated.
func (iterable Iterable[T]) GroupListsByThis(selector func(T) any) MapIterable[T] {

	mapIterable := defaultMapIterable(iterable)
	mapIterable.keySelector = selector
	mapIterable.overwrite = false

	return mapIterable

	/*
		linq.From([]T{...}).
			GroupListsByThis(
				func(item T) any {
					return item.ItemField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// GroupListsBy returns a new MapIterable where the items are grouped by the
// given field name. Items with the same key will NOT be overwritten as this
// iterable is iterated. If T is not a struct, or fieldName is not found, this
// function will panic.
func (iterable Iterable[T]) GroupListsBy(fieldName string) MapIterable[T] {

	return iterable.GroupListsByThis(
		getFieldNameFunc[T](fieldName),
	)

	/*
		linq.From([]T{...}).
			GroupListsBy("ItemField")
	*/
}

////////////////////////////////////////////////////////////////////////////////

// ByThis returns a new MapIterable where the items are grouped by the given
// key selector. Callable from other MapIterables. Designed to be used in
// tandem with the GroupThis or GroupListsOfThis functions. Items with the
// same key will NOT be overwritten as this iterable is iterated.
func (iterable DeferredKeyMapIterable[T]) ByThis(selector func(T) any) MapIterable[T] {

	iterable.keySelector = selector
	return MapIterable[T](iterable)

	/*
		linq.From([]T{...]).
			GroupThis("ItemField").
			ByThis(
				func(item T) any {
					return item.KeyField
				},
			)
	*/
}

////////////////////////////////////////////////////////////////////////////////

// By returns a new MapIterable where the items are grouped by the given field
// name. Callable from other MapIterables. Designed to be used in tandem with
// the GroupThis or GroupListsOfThis functions. Items with the same key will NOT
// be overwritten as this iterable is iterated.
func (iterable DeferredKeyMapIterable[T]) By(fieldName string) MapIterable[T] {

	return iterable.ByThis(
		getFieldNameFunc[T](fieldName),
	)

	/*
		linq.From([]T{...}).
			GroupListsOfThis("ItemField").
			By("KeyField")
	*/
}
