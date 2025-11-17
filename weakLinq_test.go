package weaklinq

import (
	"iter"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////

type testStruct struct {
	Id       int
	Name     string
	IsActive bool
}

//----------------------------------------------------------------------------//
// Constructors                                                               //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func TestFrom(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		result := From(testItems)
		resultIterator, _ := iter.Pull(result.Seq)
		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok && item != testItems[0] {
			t.Errorf("Expected first item but got %v", item)
		}

		if item, ok := resultIterator(); ok && item != testItems[1] {
			t.Errorf("Expected second item but got %v", item)
		}

		if item, ok := resultIterator(); ok {
			t.Errorf("Expected no item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("empty", func(t *testing.T) {

		var testItems []testStruct
		result := From(testItems)
		resultIterator, _ := iter.Pull(result.Seq)

		if resultIterator == nil {
			t.Errorf("Expected iterator but got nil")
		}

		if item, ok := resultIterator(); ok {
			t.Errorf("Expected no item but got %v", item)
		}
	})

	//----------------------------------------------------------------------------//
}
