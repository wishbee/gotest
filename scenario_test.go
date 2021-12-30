package gotest

import (
	"testing"
)

type SomeDataToBeUnitTested struct {
	id int
}
func (t *SomeDataToBeUnitTested)SetId(i int) {
	t.id = i
}
func (t *SomeDataToBeUnitTested)Id() int {
	return t.id
}

func TestScenario_AssertEqual(t *testing.T) {
	scenario := NewScenario(t)
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 4", func(and And, then Then) {
		v.SetId(4)
		then.AssertEqual(4, v.Id())
		and.I("reset Id as 5", func(and And, then Then) {
			v.SetId(5)
			then.AssertEqual(5, v.Id())
			then.Logln("Some information logging...")
			// ...
			// ...
			then.Logln("Some more informational logging")
		})
		and.I("reset Id again to 0", func(and And, then Then) {
			v.SetId(0)
			then.Logln("Then value of Id should be 0")
			then.AssertEqual(0,v.Id())
		})
	})
}


func TestScenario_AssertNotEqual(t *testing.T) {
	scenario := NewScenario(t)
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("expect the value of Id to be as 4", func(and And, then Then) {
			then.AssertNotEqual(4,v.Id())
		})
	})
}

func TestScenario_AssertTrue(t *testing.T) {
	scenario := NewScenario(t)
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("expect the value of Id to be as 5", func(and And, then Then) {
			then.AssertTrue(5==v.Id())
		})
	})
}

func TestScenario_AssertFalse(t *testing.T) {
	scenario := NewScenario(t)
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("expect the value of Id to be as 4", func(and And, then Then) {
			then.AssertFalse(4==v.Id())
		})
	})
}

func TestScenario_AssertNil(t *testing.T) {
	scenario := NewScenario(t)
	scenario.When("I initialise the pointer object with Nil", func(and And, then Then) {
		var v *SomeDataToBeUnitTested
		v = nil
		and.I("expect the variable as nil", func(and And, then Then) {
			then.AssertNil(v)
		})
	})
}