package assert

import (
	"reflect"
	"testing"
)

func Equals(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("%+v is not equal to %+v", expected, actual)
	}
}
