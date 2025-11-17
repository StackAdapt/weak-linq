package weaklinq

import "testing"

//----------------------------------------------------------------------------//
// Materialization                                                            //
//----------------------------------------------------------------------------//

////////////////////////////////////////////////////////////////////////////////

func TestAndAssignToMap(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}

		resultMap := make(map[int]string)
		From(testItems).Group("Name").By("Id").AndAssignToMap(&resultMap)

		if resultMap[testItems[0].Id] != testItems[0].Name {
			t.Errorf("Expected first item but got %v", resultMap[testItems[0].Id])
		}

		if resultMap[testItems[1].Id] != testItems[1].Name {
			t.Errorf("Expected second item but got %v", resultMap[testItems[1].Id])
		}

		if len(resultMap) != 2 {
			t.Errorf("Expected 2 items but got %v", len(resultMap))
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad result type", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultMap := make(map[int]string)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Group("Id").By("Name").AndAssignToMap(&resultMap)
	})

	//----------------------------------------------------------------------------//

	t.Run("bad type", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultMap := make([]string, 0)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Group("Id").By("Name").AndAssignToMap(&resultMap)
	})

	//----------------------------------------------------------------------------//

	t.Run("not pointer", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultMap := make(map[int]string)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Group("Id").By("Name").AndAssignToMap(resultMap)
	})

	//----------------------------------------------------------------------------//

	t.Run("with lists", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
			{Id: 1, Name: "Test 3"},
		}

		resultMap := make(map[int][]string)
		From(testItems).GroupListsOf("Name").By("Id").AndAssignToMap(&resultMap)

		if resultMap[testItems[0].Id][0] != testItems[0].Name {
			t.Errorf("Expected first item but got %v", resultMap[testItems[0].Id][0])
		}

		if resultMap[testItems[1].Id][0] != testItems[1].Name {
			t.Errorf("Expected second item but got %v", resultMap[testItems[1].Id][0])
		}

		if resultMap[testItems[2].Id][1] != testItems[2].Name {
			t.Errorf("Expected third item but got %v", resultMap[testItems[2].Id][0])
		}

		if len(resultMap[testItems[0].Id]) != 2 {
			t.Errorf("Expected 2 items but got %v", len(resultMap[testItems[0].Id]))
		}

		if len(resultMap[testItems[1].Id]) != 1 {
			t.Errorf("Expected 1 item but got %v", len(resultMap[testItems[1].Id]))
		}

		if len(resultMap[testItems[2].Id]) != 2 {
			t.Errorf("Expected 2 items but got %v", len(resultMap[testItems[2].Id]))
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("with lists not slice items", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
			{Id: 1, Name: "Test 3"},
		}

		resultMap := make(map[int]string)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).GroupListsOf("Name").By("Id").AndAssignToMap(&resultMap)
	})

	//----------------------------------------------------------------------------//
}

////////////////////////////////////////////////////////////////////////////////

func TestAndAssignToSlice(t *testing.T) {

	//----------------------------------------------------------------------------//

	t.Run("generic", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
			{Id: 2, Name: "Test 2"},
		}
		resultSlice := make([]string, 0)
		From(testItems).Get("Name").AndAssignToSlice(&resultSlice)

		if resultSlice[0] != testItems[0].Name {
			t.Errorf("Expected first item but got %v", resultSlice[0])
		}

		if resultSlice[1] != testItems[1].Name {
			t.Errorf("Expected second item but got %v", resultSlice[1])
		}

		if len(resultSlice) != 2 {
			t.Errorf("Expected 2 items but got %v", len(resultSlice))
		}
	})

	//----------------------------------------------------------------------------//

	t.Run("bad result type", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultSlice := make(map[int]string)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Get("Name").AndAssignToSlice(&resultSlice)
	})

	//----------------------------------------------------------------------------//

	t.Run("bad type", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultSlice := make([]int, 0)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Get("Name").AndAssignToSlice(&resultSlice)
	})

	//----------------------------------------------------------------------------//

	t.Run("not pointer", func(t *testing.T) {

		testItems := []testStruct{
			{Id: 1, Name: "Test 1"},
		}

		resultSlice := make([]string, 0)

		defer func() {
			if err := recover(); err == nil {
				t.Errorf("Expected panic but got %v", err)
			}
		}()

		From(testItems).Get("Name").AndAssignToSlice(resultSlice)
	})

	//----------------------------------------------------------------------------//
}
