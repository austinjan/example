package slice

import "testing"

func addeach(s []int) {
	for i := range s {
		s[i]++
	}
}

// TestSlideAsFunctionParameter tests pass slice as function parameter, after call function, the slice is changed
func TestSlideAsFunctionParameter(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Log("slice :", s)
	// pass a slice to a function, function will add each element by 1
	addeach(s)
	// slice out side the function will be changed
	t.Log("after function executed slice is ", s)
}
