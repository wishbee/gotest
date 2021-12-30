package gotest

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

const (
	Pass = true
	Fail = false
)

func logWithCaller(depth int, format string, a ...interface{}) string {
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		fmt.Printf("%s:%d: ", filepath.Base(file), line)
		fmt.Printf(format, a...)
	}
	return ""
}

type Logger interface {
	Logf(format string, a ...interface{})
	Logln(a ...interface{})
}

type Scenario interface {
	When(condition string, v func(and And, then Then))
}

type When interface {
	When(condition string, v func(and And, then Then))
}

type And interface {
	I(condition string, v func(and And, then Then))
}

type Then interface {
	AssertEqual(expected, actual interface{})
	AssertNotEqual(expected, actual interface{})
	AssertTrue(expected bool)
	AssertFalse(expected bool)
	AssertNil(expected interface{})
	AssertNotNil(expected interface{})
	Logger
}

type scenario struct {
	t *testing.T
	//scenario string
	//assertions []bool
	continueOnAssertionFailed bool
	depth                     int
}

func NewScenario(t *testing.T) Scenario {
	return &scenario{
		t:                         t,
		continueOnAssertionFailed: false,
	}
}

func (s *scenario) Logf(format string, a ...interface{}) {
	fmt.Printf("%s", strings.Repeat(" ", 4*s.depth))
	fmt.Printf(format, a...)

}

func (s *scenario) Logln(a ...interface{}) {
	fmt.Printf("%s", strings.Repeat(" ", 4*s.depth))
	fmt.Println(a...)

}

func (s *scenario) When(condition string, v func(and And, then Then)) {
	s.Logln("When", condition)
	s.depth++
	v(s, s)
	s.depth--
}

func (s *scenario) I(condition string, v func(and And, then Then)) {
	s.Logln("And I", condition)
	s.depth++
	v(s, s)
	s.depth--
}

// Assert implementations
func (s *scenario) AssertEqual(expected, actual interface{}) {
	s.Logf("Then I expect the value should be equal to %v\n", expected)
	if !reflect.DeepEqual(expected, actual) {
		logWithCaller(2, "Assertion failed: expected value %v and actual value %v are not equal\n", expected, actual)
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (s *scenario) AssertNotEqual(expected, actual interface{}) {
	s.Logf("Then I expect the value should not be equal to %v\n", expected)
	if reflect.DeepEqual(expected, actual) {
		logWithCaller(2, "Assertion failed: expected value %v and actual value %v are equal\n", expected, actual)
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (s *scenario) AssertTrue(expected bool) {
	s.Logln("Then I expect the value should be True")
	if !expected {
		logWithCaller(2, "Assertion failed: expected 'true' but it is 'false'\n")
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (s *scenario) AssertFalse(expected bool) {
	s.Logln("Then I expect the value should be False")
	if expected {
		logWithCaller(2, "Assertion failed: expected 'false' but it is 'true'\n")
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (s *scenario) AssertNil(expected interface{}) {
	s.Logln("Then I expect the value should be Nil")

	if !isNil(expected) {
		logWithCaller(2, "Assertion failed: expected 'nil' value but it is not 'nil'\n")
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (s *scenario) AssertNotNil(expected interface{}) {
	s.Logln("Then I expect the value should not be Nil")
	if isNil(expected) {
		logWithCaller(2, "Assertion failed: expected 'not nil' value but it is a 'nil'\n")
		s.t.Fail()
		if !s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Map, reflect.Chan, reflect.Slice, reflect.Ptr, reflect.Array:
		return reflect.ValueOf(v).IsNil()
	}
	return false
}
