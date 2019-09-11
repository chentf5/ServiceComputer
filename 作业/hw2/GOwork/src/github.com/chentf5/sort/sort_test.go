package sort

import "testing"
import "reflect"

func TestSort(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1,2,3,4,5,6},[]int{1,2,3,4,5,6}},
		{[]int{67,34,22,111,54,2121,5431,131},[]int{22,34,54,67,111,131,2121,5431}},
		{[]int{},[]int{}},
	}
	for _, c := range cases {
		got := Sort(c.in)
		if ( !reflect.DeepEqual(got,c.want) ) {
			t.Errorf("quicksort(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}