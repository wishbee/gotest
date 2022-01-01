# gotest

Intuitive and simple golang testing framework which helps in writing unit tests in a way which improves the readability of the test.

Here is an example unit test which demonstrate the easy and intuitive way of writing unit test.
``` 
package test

import (
	. "github.com/wishbee/gotest"
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
```

Below is the output from above unit test.
```
=== RUN   TestScenario_AssertEqual
Scenario: Test ShouldBeEqualTo
    When I set the Id as 4
        Then I expect the value should be equal to 4
        And I reset Id as 5
            Then I expect the value should be equal to 5
            Some information logging...
            Some more informational logging
        And I reset Id again to 0
            Then I expect the value should be equal to 0
--- PASS: TestScenario_AssertEqual (0.00s)
PASS

Process finished with exit code 0
```