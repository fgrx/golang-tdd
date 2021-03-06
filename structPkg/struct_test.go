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

	// checkArea:=func(t testing.TB, shape Shape, want float64){
	// 	t.Helper()
	// 	got := shape.Area() 
	// 	if got!=want {
	// 		t.Errorf("expected %g got %g",got,want)
	// 	}
	// }

	// t.Run("Rectangles",func(t *testing.T) {
	// 	rectangle := Rectangle{12,6}
	// 	checkArea(t, rectangle,72.0)
	// })


	// t.Run("circles",func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle,314.1592653589793)
	// })

	// refacto
	areaTests:=[]struct {
		name string
		shape Shape
		want float64
	}{
		{name: "Rectangle",shape:Rectangle{Width:12,Height:6}, want:72},
		{name: "Circles", shape:Circle{Radius:10}, want:314.1592653589793},
		{name:"Triangle",shape:Triangle{Base:12, Height:6}, want: 36.0},
	}

	for _, tt:= range areaTests{
		got:= tt.shape.Area()
		if got!=tt.want {
			t.Errorf("expected %g got %g",got,tt.want)
		}
	}
}