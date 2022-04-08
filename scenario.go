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
	Expect(objectName string, actualValue interface{}) Assert
	Logger
}

type Assert interface {
	ShouldBeEqualTo(expected interface{})
	ShouldNotBeEqualTo(expected interface{})
	ShouldBeTrue()
	ShouldBeFalse()
	ShouldBeNil()
	ShouldNotBeNil()
}

type assert struct {
	objectName string
	actual     interface{}
	s          *scenario
}

type scenario struct {
	t *testing.T
	//scenario string
	//assertions []bool
	continueOnAssertionFailed bool
	depth                     int
	scenario                  string
}

func NewScenario(t *testing.T, scenarioText string) Scenario {
	return &scenario{
		t:                         t,
		continueOnAssertionFailed: true,
		scenario:                  scenarioText,
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

func (s *scenario) When(action string, v func(and And, then Then)) {
	fmt.Println("Scenario:", s.scenario)
	s.depth++
	s.Logln("When", action)
	s.depth++
	v(s, s)
	s.depth--
}

func (s *scenario) I(action string, v func(and And, then Then)) {
	s.Logln("And I", action)
	s.depth++
	v(s, s)
	s.depth--
}

func (s *scenario) Expect(objectName string, actualValue interface{}) Assert {
	return &assert{
		objectName: objectName,
		actual:     actualValue,
		s:          s,
	}
}

// Assert implementations
func (a *assert) ShouldBeEqualTo(expected interface{}) {
	a.s.Logf("Then I expect %s should be equal to %v\n", a.objectName, expected)
	if !reflect.DeepEqual(expected, a.actual) {
		logWithCaller(2, "Assertion failed: expected value %v and actual value %v are not equal\n", expected, a.actual)
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) ShouldNotBeEqualTo(expected interface{}) {
	a.s.Logf("Then I expect %s should not be equal to %v\n", a.objectName, expected)
	if reflect.DeepEqual(expected, a.actual) {
		logWithCaller(2, "Assertion failed: expected value %v and actual value %v are equal\n", expected, a.actual)
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) ShouldBeTrue() {
	a.s.Logf("Then I expect %s should be True\n", a.objectName)
	b, ok := a.actual.(bool)
	if !ok {
		logWithCaller(2, "Assertion failed: the actual value passed is not boolean, but %T\n", a.actual)
	}
	if !b {
		logWithCaller(2, "Assertion failed: expected 'true' but it is 'false'\n")
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) ShouldBeFalse() {
	a.s.Logf("Then I expect %s should be False\n", a.objectName)
	b, ok := a.actual.(bool)
	if !ok {
		logWithCaller(2, "Assertion failed: the actual value passed is not boolean, but %T\n", a.actual)
	}
	if b {
		logWithCaller(2, "Assertion failed: expected 'false' but it is 'true'\n")
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) ShouldBeNil() {
	a.s.Logf("Then I expect %s should be Nil\n", a.objectName)
	if !a.isNil(a.actual) {
		logWithCaller(2, "Assertion failed: expected 'nil' value but it is not 'nil'\n")
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) ShouldNotBeNil() {
	a.s.Logf("Then I expect %s should not be Nil\n", a.objectName)
	if a.isNil(a.actual) {
		logWithCaller(2, "Assertion failed: expected 'not nil' value but it is a 'nil'\n")
		a.s.t.Fail()
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
		return
	}
}

func (a *assert) isNil(v interface{}) bool {
	if v == nil {
		return true
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.Map, reflect.Chan, reflect.Slice, reflect.Ptr, reflect.Array, reflect.Interface, reflect.Func:
		return reflect.ValueOf(v).IsNil()
	default:
		logWithCaller(3, "passed object can not be tested for nil'ness due to its type: %T\n", v)
		if !a.s.continueOnAssertionFailed {
			panic("assertion failed")
		}
	}
	return false
}
