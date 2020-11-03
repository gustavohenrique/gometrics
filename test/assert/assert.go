package assert

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func Equal(t *testing.T, a interface{}, b interface{}, messages ...string) {
	if a == b {
		return
	}
	message := fmt.Sprintf("\n%v !=\n%v", a, b)
	if len(messages) > 0 {
		message = messages[0]
	}
	t.Fatal(message)
}

func NotEqual(t *testing.T, a interface{}, b interface{}, messages ...string) {
	if a != b {
		return
	}
	Equal(t, a, b, messages...)
}

func DeepEqual(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		message := fmt.Sprintf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		t.Fatal(message)
		// t.FailNow()
	}
}

func Nil(t *testing.T, a interface{}, messages ...string) {
	if isNil(a) {
		return
	}
	message := fmt.Sprintf("Expected nil, but got: %#v", a)
	if len(messages) > 0 {
		message = messages[0]
	}
	t.Fatal(message)
}

func NotNil(t *testing.T, a interface{}, messages ...string) {
	if !isNil(a) {
		return
	}
	message := fmt.Sprintf("Expected not nil, but got: %#v", a)
	if len(messages) > 0 {
		message = messages[0]
	}
	t.Fatal(message)
}

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	value := reflect.ValueOf(a)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}
	return false
}

func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}
	return false
}
