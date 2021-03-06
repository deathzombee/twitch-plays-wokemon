package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInterpetron(t *testing.T) {
	var aMap = map[string]Command{"A": "B", "C": "D"}

	Convey("constructor returns an interpetron", t, func() {
		dummyInterpetron := new(CommandInterpretron)
		dummyInterpetron.dict = aMap

		So(NewInterpetron(aMap), ShouldResemble, dummyInterpetron)
	})

	Convey("parseCommand takes a string and returns a command", t, func() {
		testInterpetron := NewInterpetron(aMap)
		someCommand, _ := testInterpetron.parseCommand("A")
		So(someCommand, ShouldEqual, Command("B"))
		_, err := testInterpetron.parseCommand("B")
		So(err, ShouldBeError)
	})
}
