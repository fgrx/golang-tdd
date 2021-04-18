package structPkg

import "testing"



func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0,10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if(got !=want){
		t.Errorf("expected %.2f got %.2f",want, got)
	}
}


func TestArea(t *testing.T){
	rectangle := Rectangle{4.0,3.0}

	got := Area(rectangle)
	want:=12.0

	if(got !=want){
		t.Errorf("expected %.2f got %.2f",want, got)
	}
}