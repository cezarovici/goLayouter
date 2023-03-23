package stack_test

import (
	"testing"

	"github.com/cezarovici/goLayouter/helpers/stack"
	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var res = stack.Stack{"folder1"}

	type testCase struct {
		test   string
		input  string
		output stack.Stack
	}

	testCases := []testCase{
		{
			test:   "Push one element",
			input:  "folder2",
			output: stack.Stack{"folder1", "folder2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			res.Push(tc.input)
			require.Equal(t, tc.output, res)
		})
	}
}

func TestPop(t *testing.T) {
	var res = stack.Stack{"folder1", "folder2"}
	type testCase struct {
		test string

		output stack.Stack
	}

	testCases := []testCase{
		{
			test:   "Pop one element",
			output: stack.Stack{"folder1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			res.Pop()
			require.Equal(t, tc.output, res)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type testCase struct {
		test   string
		input  stack.Stack
		output bool
	}

	testCases := []testCase{
		{
			test:   "empty stack.stack",
			input:  stack.Stack{},
			output: true,
		},
		{
			test:   "non-empty stack.stack",
			input:  stack.Stack{"file"},
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, tc.input.IsEmpty())
		})
	}
}

func TestPeek(t *testing.T) {
	type testCase struct {
		test   string
		input  stack.Stack
		output any
	}

	testCases := []testCase{
		{
			test:   "non empty",
			input:  stack.Stack{"file", "folder"},
			output: "folder",
		},
		{
			test:   "empty stack.stack",
			input:  stack.Stack{},
			output: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, tc.input.Peek())
		})
	}
}

func TestString(t *testing.T) {
	type testCase struct {
		test   string
		input  stack.Stack
		output string
	}

	testCases := []testCase{
		{
			test:   "non empty",
			input:  stack.Stack{"file", "folder"},
			output: "file/folder",
		},
		{
			test:   "empty stack.stack",
			input:  stack.Stack{},
			output: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, tc.input.String())
		})
	}
}
