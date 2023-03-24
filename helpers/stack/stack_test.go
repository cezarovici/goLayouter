package stack_test

import (
	"testing"

	"github.com/cezarovici/goLayouter/helpers/stack"
	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	t.Parallel()
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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			res.Push(currentTestCase.input)
			require.Equal(t, currentTestCase.output, res)
		})
	}
}

func TestPop(t *testing.T) {
	t.Parallel()

	res := stack.Stack{"folder1", "folder2"}
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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			res.Pop()
			require.Equal(t, currentTestCase.output, res)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	t.Parallel()

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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase
		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, currentTestCase.output, currentTestCase.input.IsEmpty())
		})
	}
}

func TestPeek(t *testing.T) {
	t.Parallel()

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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, currentTestCase.output, currentTestCase.input.Peek())
		})
	}
}

func TestString(t *testing.T) {
	t.Parallel()

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

	for _, currentTestCase := range testCases {
		currentTestCase := currentTestCase

		t.Run(currentTestCase.test, func(t *testing.T) {
			t.Parallel()

			require.Equal(t, currentTestCase.output, currentTestCase.input.String())
		})
	}
}
