package stack

import "errors"

// RuneStack corresponds to the stack implementation for
// rune types.
type RuneStack []rune

// New returns a new instand of the stack
func NewRuenStack() *RuneStack {
	return &RuneStack{}
}

// IsEmpty determines if there are elements in the stack.
func (rs *RuneStack) IsEmpty() bool {
	return len(*rs) == 0
}

// Push add a new element to the stack.
func (rs *RuneStack) Push(r rune) {
	*rs = append(*rs, r)
}

// Pop removes the last inserted element from the stack and
// returns it.
func (rs *RuneStack) Pop() (rune, error) {
	// Validate if the stack is empty before proceeding.
	if rs.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	// Get the last value in the array.
	value := (*rs)[len(*rs)-1]

	// Modify the stack to contain only the remaining components.
	*rs = (*rs)[:len(*rs)-1]

	// Return the value after popping it.
	return value, nil
}
