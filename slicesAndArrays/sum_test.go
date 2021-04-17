package slicesAndArrays

import (
	"reflect"
	"testing"
)



func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers",func(t *testing.T){
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("go %d want %d given : %v",got,want,numbers)
		}
	})
}

func TestSumAll(t *testing.T){
	got:=SumAll([]int{1,2},[]int{0,9})
	want:=[]int{3,9}

	// Ne marche pas entre slice
	// if got!=want{
	// 	t.Errorf("got %v want %v", got, want)
	// }

	if !reflect.DeepEqual(got,want){
		t.Errorf("got %v want %v", got, want)
	}
}