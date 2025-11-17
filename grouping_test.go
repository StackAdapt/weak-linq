package weaklinq

import (
	"iter"
	"testing"
)

//----------------------------------------------------------------------------//
// Grouping                                                                   //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func TestGroupThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GroupThis(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if result.itemSelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
	}

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestGroup(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).Group("Name")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if result.itemSelector(testItems[0]) != testItems[0].Name {
			t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
		}

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad field name", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		result := From(testItems).Group("BadFieldName")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		result.itemSelector(testItems[0])
	})

	//----------------------------------------------------------------------------//
}

////////////////////////////////////////////////////////////////////////////////

func TestGroupListsOfThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GroupListsOfThis(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if result.itemSelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
	}

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestGroupListsOf(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).GroupListsOf("Id")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad field name", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		result := From(testItems).GroupListsOf("BadFieldName")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		result.itemSelector(testItems[0])
	})

	//----------------------------------------------------------------------------//
}

////////////////////////////////////////////////////////////////////////////////

func TestGroupByThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GroupByThis(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}

	if result.keySelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestGroupBy(t *testing.T) {

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).GroupBy("Id")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if result.keySelector(testItems[0]) != testItems[0].Id {
			t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("struct pointer", func(t *testing.T) {

		testItems := []*testStruct{
			{Id: 1, Name: "Test 1", IsActive: true},
			{Id: 2, Name: "Test 2", IsActive: true},
		}

		result := From(testItems).GroupBy("Id")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if result.keySelector(testItems[0]) != testItems[0].Id {
			t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
		}

		if item, ok := resultIterator(); ok && item != testItems[1] {
			t.Errorf("Expected second item but got %v", item)
		}

		if result.keySelector(testItems[1]) != testItems[1].Id {
			t.Errorf("Expected second item but got %v", result.keySelector(testItems[1]))
		}

		if item, ok := resultIterator(); ok {
			t.Errorf("Expected no item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

}

////////////////////////////////////////////////////////////////////////////////

func TestGroupListsByThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GroupListsByThis(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}

	if result.keySelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestGroupListsBy(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GroupListsBy("Id")
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}

	if result.keySelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestByThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).Group("Name").ByThis(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.itemIterable.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}

	if result.itemSelector(testItems[0]) != testItems[0].Name {
		t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
	}

	if result.keySelector(testItems[0]) != testItems[0].Id {
		t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestBy(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).Group("Id").By("Name")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if result.itemSelector(testItems[0]) != testItems[0].Id {
			t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
		}

		if result.keySelector(testItems[0]) != testItems[0].Name {
			t.Errorf("Expected first item but got %v", result.keySelector(testItems[0]))
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad field name", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).Group("Id").By("BadFieldName")
		resultIterator, _ := iter.Pull(result.itemIterable.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if result.itemSelector(testItems[0]) != testItems[0].Id {
			t.Errorf("Expected first item but got %v", result.itemSelector(testItems[0]))
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		result.keySelector(testItems[0])
	})

	//----------------------------------------------------------------------------//
}
