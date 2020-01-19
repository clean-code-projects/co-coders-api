package user_test

import (
	"reflect"
	"testing"
)

func TestAssertEqualWorks(t *testing.T) {
	AssertEqual(t, "Hello", "Hello")
}

func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%+v is not equal to %+v", expected, actual)
	}

}
