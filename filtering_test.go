package weaklinq

import (
	"iter"
	"testing"
)

//----------------------------------------------------------------------------//
// Filtering                                                                  //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func TestFilterOnThis(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1", IsActive: true},
		{Id: 2, Name: "Test 2", IsActive: false},
		{Id: 3, Name: "Test 3", IsActive: true},
	}

	result := From(testItems).FilterOnThis(
		func(t testStruct) bool { return t.IsActive },
	)
	resultIterator, _ := iter.Pull(result.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0] {
		t.Errorf("Expected first item but got %v", item)
	}

	if item, ok := resultIterator(); ok && item != testItems[2] {
		t.Errorf("Expected third item but got %v", item)
	}

	if item, ok := resultIterator(); ok {
		t.Errorf("Expected no item but got %v", item)
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestFilterOn(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1", IsActive: true},
			{Id: 2, Name: "Test 2", IsActive: false},
		}

		result := From(testItems).FilterOn("IsActive")
		resultIterator, _ := iter.Pull(result.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if item, ok := resultIterator(); ok {
			t.Errorf("Expected no item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad field name", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1", IsActive: true},
		}

		result := From(testItems).FilterOn("BadFieldName")
		resultIterator, _ := iter.Pull(result.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		resultIterator()
	})

	//----------------------------------------------------------------------------//

	t.Run("bad item type", func(t *testing.T) {

		testItems := []int{1, 2, 3}

		result := From(testItems).FilterOn("Name")
		resultIterator, _ := iter.Pull(result.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		resultIterator()
	})

	//----------------------------------------------------------------------------//
}

////////////////////////////////////////////////////////////////////////////////

func TestDistinct(t *testing.T) {

	testItems := []int{1, 2, 1, 3, 2, 4}

	result := From(testItems).Distinct()
	resultIterator, _ := iter.Pull(result.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != 1 {
		t.Errorf("Expected first item but got %v", item)
	}

	if item, ok := resultIterator(); ok && item != 2 {
		t.Errorf("Expected second item but got %v", item)
	}

	if item, ok := resultIterator(); ok && item != 3 {
		t.Errorf("Expected third item but got %v", item)
	}

	if item, ok := resultIterator(); ok && item != 4 {
		t.Errorf("Expected fourth item but got %v", item)
	}

	if item, ok := resultIterator(); ok {
		t.Errorf("Expected no item but got %v", item)
	}
}
