package test_helpers

import (
  "testing"
  "reflect"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func AssertTypeEqual(t *testing.T, a interface{}, b interface{}) {
  a_t := reflect.TypeOf(a)
  b_t := reflect.TypeOf(b)
  if a_t != b_t {
    t.Fatalf("%s not %s", a_t, b_t)
  }
}

func AssertDeepEqual(t *testing.T, a interface{}, b interface{}) {
  eq := reflect.DeepEqual(a, b)
  if !eq {
    t.Fatalf("got: %+v\nexpected: %+v\n", a, b)
  }
}

