package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	var stack = Stack{"folder1"}

	type testCase struct {
		test   string
		input  string
		output Stack
	}

	testCases := []testCase{
		{
			test:   "Push one element",
			input:  "folder2",
			output: Stack{"folder1", "folder2"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			stack.Push(tc.input)
			require.Equal(t, tc.output, stack)
		})
	}
}

func TestPop(t *testing.T) {
	var stack = Stack{"folder1", "folder2"}
	type testCase struct {
		test string

		output Stack
	}

	testCases := []testCase{
		{
			test:   "Pop one element",
			output: Stack{"folder1"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			stack.Pop()
			require.Equal(t, tc.output, stack)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type testCase struct {
		test   string
		input  Stack
		output bool
	}

	testCases := []testCase{
		{
			test:   "empty stack",
			input:  Stack{},
			output: true,
		},
		{
			test:   "non-empty stack",
			input:  Stack{"file"},
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
		input  Stack
		output any
	}

	testCases := []testCase{
		{
			test:   "non empty",
			input:  Stack{"file", "folder"},
			output: "folder",
		},
		{
			test:   "empty stack",
			input:  Stack{},
			output: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, tc.input.Peek())
		})
	}
}
