package main

import (
	"testing"
)

func TestShowCompletedStatus(t *testing.T) {
	// Define test cases
	tests := []struct {
		isDone   bool
		expected string
	}{
		{true, "yes"},
		{false, "no"},
	}

	// Iterate over test cases
	for _, test := range tests {
		result := showCompletedStatus(test.isDone)
		if result != test.expected {
			t.Errorf("showCompletedStatus(%v) = %v; want %v", test.isDone, result, test.expected)
		}
	}
}

func TestAddTodo(t *testing.T) {
	allTodos = []Todo{}
	testTitle := "Test Todo"
	addTodo(testTitle)
	if len(allTodos) != 1 {
		t.Errorf("Expected 1 todo, got %d", len(allTodos))
	}
	addedTodo := allTodos[0]
	if addedTodo.Title != testTitle {
		t.Errorf("Expected title %s, got %s", testTitle, addedTodo.Title)
	}

	if addedTodo.IsDone != false {
		t.Errorf("Expected IsDone to be false, got %v", addedTodo.IsDone)
	}

}
