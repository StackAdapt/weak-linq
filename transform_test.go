package weaklinq

import (
	"iter"
	"testing"
)

//----------------------------------------------------------------------------//
// Transform                                                                  //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func TestGetThese(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).GetThese(
		func(t testStruct) any { return t.Id },
	)
	resultIterator, _ := iter.Pull(result.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok && item != testItems[0].Id {
		t.Errorf("Expected first item but got %v", item)
	}

	if item, ok := resultIterator(); ok && item != testItems[1].Id {
		t.Errorf("Expected second item but got %v", item)
	}

	if item, ok := resultIterator(); ok {
		t.Errorf("Expected no item but got %v", item)
	}
}

////////////////////////////////////////////////////////////////////////////////

func TestGet(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems).Get("Name")
		resultIterator, _ := iter.Pull(result.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0].Name {
			t.Errorf("Expected first item but got %v", item)
		}

		if item, ok := resultIterator(); ok && item != testItems[1].Name {
			t.Errorf("Expected second item but got %v", item)
		}

		if item, ok := resultIterator(); ok {
			t.Errorf("Expected no item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad field name", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		result := From(testItems).Get("BadFieldName")
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

		result := From(testItems).Get("Name")
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

func TestAsAny(t *testing.T) {

	testItems := []testStruct{
		{Id: 1, Name: "Test 1"},
		{Id: 2, Name: "Test 2"},
	}

	result := From(testItems).AsAny()
	resultIterator, _ := iter.Pull(result.Seq)

	if resultIterator == nil {
		t.Errorf("Expected iterator but got nil")
	}

	if item, ok := resultIterator(); ok {
		if item.(testStruct) != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}
	}

	if item, ok := resultIterator(); ok {
		if item.(testStruct) != testItems[1] {
			t.Errorf("Expected second item but got %v", item)
		}
	}

	if item, ok := resultIterator(); ok {
		t.Errorf("Expected no item but got %v", item)
	}
}
