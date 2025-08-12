package Perimeter

import "testing"


func TestPerimeter(t *testing.T){

	certainResult := func(t *testing.T, get, want float64) {
		t.Helper()
        if get != want {
            t.Errorf("got '%f' want '%f'", get, want)
        }
	}

	t.Run("test number",func(t *testing.T){
		getWidth := 12
		getHeight := 24
		get := Perimeter(float64(getWidth),float64(getHeight))
		want := 72
		certainResult(t,get,float64(want))
	})

	
}