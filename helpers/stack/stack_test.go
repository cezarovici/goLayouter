package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPush(t *testing.T) {
	t.Parallel()
	var res = Stack{"folder1"}

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

	res := Stack{"folder1", "folder2"}
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
		input  Stack
		output bool
	}

	testCases := []testCase{
		{
			test:   "empty stack.stack",
			input:  Stack{},
			output: true,
		},
		{
			test:   "non-empty stack.stack",
			input:  Stack{"file"},
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
			test:   "empty stack.stack",
			input:  Stack{},
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
		input  Stack
		output string
	}

	testCases := []testCase{
		{
			test:   "non empty",
			input:  Stack{"file", "folder"},
			output: "file/folder",
		},
		{
			test:   "empty stack.stack",
			input:  Stack{},
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
