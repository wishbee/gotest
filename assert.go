package gotest

import "reflect"

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

// ShouldBeEqualTo checks if the
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
