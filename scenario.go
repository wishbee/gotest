package gotest

import (
	"fmt"
	"path/filepath"
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
