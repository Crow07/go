package simplemath

import "testing"

func Testadd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3{
		t.Errorf("Add(1,2) failed.Got %d, excperted 3.", r)
	}
}