package io

import (
	"github.com/remogatto/prettytest"
	"os"
	"testing"
)

type testStdinGetterSuite struct {
	prettytest.Suite
}

func getterSuiteSetup() (getter *StdinGetter) {
	getter = NewStdinGetter()
	return
}

func TestStdinGetterRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testStdinGetterSuite),
	)
}

func (t *testStdinGetterSuite) TestNewStdinGetter() {
	getter := NewStdinGetter()
	t.Equal(os.Stdin, getter.in)
}

func (t *testStdinGetterSuite) TestGetString() {
	var input string
	SimulateInput("sample-string", func() {
		getter := getterSuiteSetup()
		input = getter.GetString()
	})

	t.Equal("sample-string", input)
}

func (t *testStdinGetterSuite) TestGetInt() {
	var input int
	SimulateInput("1", func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})

	t.Equal(1, input)
}

func (t *testStdinGetterSuite) TestRepeatedlyAskForInt() {
	var input int
	SimulateMultipleInput([]string{"A", "C", "8"}, func() {
		getter := getterSuiteSetup()
		input = getter.GetInt()
	})
	t.Equal(8, input)
}

func (t *testStdinGetterSuite) TestGetIntWithCallback() {
	var count int
	SimulateMultipleInput([]string{"A", "C", "8"}, func() {
		getter := getterSuiteSetup()
		getter.GetIntWithCallback(func() {
			count += 1
		})
	})

	t.Equal(2, count)
}
