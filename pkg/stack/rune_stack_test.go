package stack

import (
	"testing"
)

// TestRuneStack_IsEmpty validates that is empty works properly
// for empty and non empty scenarios.
func TestRuneStack_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		rs   *RuneStack
		want bool
	}{
		{
			name: "EmptyStack",
			rs:   &RuneStack{},
			want: true,
		},
		{
			name: "NonEmptyStack",
			rs:   &RuneStack{'H', 'E', 'L', 'L', 'O'},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rs.IsEmpty(); got != tt.want {
				t.Errorf("RuneStack.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestRuneStack_Push validates the contents of the
// stack after a push operation
func TestRuneStack_Push(t *testing.T) {

	// Intialize an empty stack.
	rs := NewRuenStack()

	// Push a rune into the stack.
	rs.Push('W')
	rs.Push('X')
	rs.Push('C')

	// Validations.
	if len(*rs) != 3 {
		t.Error("missing elements in stack")
		return
	}
}

func TestRuneStack_HappyPath(t *testing.T) {

	// Intialize an empty stack.
	rs := NewRuenStack()

	// Push a rune into the stack.
	for _, r := range "WXC" {
		rs.Push(r)
	}

	// Validate non empty stack.
	if rs.IsEmpty() {
		t.Error("non-empty stack was expected")
		return
	}

	// Validate pop operations.
	for _, r := range "CXW" {
		val, err := rs.Pop()
		if err != nil || val != r {
			t.Errorf("pop failed, error: %s, val: %v", err.Error(), val)
			return
		}
	}

	// Validate empty stack.
	if !rs.IsEmpty() {
		t.Error("non-empty stack was expected")
		return
	}

	// Validate pop error.
	_, err := rs.Pop()
	if err == nil {
		t.Errorf("pop failed was expected")
		return
	}
}
