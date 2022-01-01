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
	scenario := NewScenario(t, "Test ShouldBeEqualTo")
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 4", func(and And, then Then) {
		v.SetId(4)
		then.Expect(v.Id()).ShouldBeEqualTo(4)
		and.I("reset Id as 5", func(and And, then Then) {
			v.SetId(5)
			then.Expect(v.Id()).ShouldBeEqualTo(5)
			then.Logln("Some information logging...")
			// ...
			// ...
			then.Logln("Some more informational logging")
		})
		and.I("reset Id again to 0", func(and And, then Then) {
			v.SetId(0)
			then.Expect(v.Id()).ShouldBeEqualTo(0)
		})
	})
}


func TestScenario_AssertNotEqual(t *testing.T) {
	scenario := NewScenario(t, "Test ShouldNotBeEqualTo")
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("check the id again", func(and And, then Then) {
			then.Expect(v.Id()).ShouldNotBeEqualTo(4)
		})
	})
}

func TestScenario_AssertTrue(t *testing.T) {
	scenario := NewScenario(t, "Test ShouldBeTrue")
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("check the id is 5", func(and And, then Then) {
			then.Expect(5==v.Id()).ShouldBeTrue()
		})
	})
}

func TestScenario_AssertFalse(t *testing.T) {
	scenario := NewScenario(t, "Test AssertFalse")
	v := &SomeDataToBeUnitTested{}
	scenario.When("I set the Id as 5", func(and And, then Then) {
		v.SetId(5)
		and.I("check the id is 4", func(and And, then Then) {
			then.Expect(4==v.Id()).ShouldBeFalse()
		})
	})
}

func TestScenario_AssertNil(t *testing.T) {
	scenario := NewScenario(t, "Test AssertNil")
	scenario.When("I initialise the pointer object with Nil", func(and And, then Then) {
		var v *SomeDataToBeUnitTested
		v = nil
		and.I("check the pointer object is nil", func(and And, then Then) {
			then.Expect(v).ShouldBeNil()
		})
	})
}

func TestScenario_AssertNotNil(t *testing.T) {
	scenario := NewScenario(t, "Test AssertNotNil")
	scenario.When("I initialise the pointer object with not Nil", func(and And, then Then) {
		v := &SomeDataToBeUnitTested{}
		and.I("check the pointer object is not nil", func(and And, then Then) {
			then.Expect(v).ShouldNotBeNil()
		})
	})
}