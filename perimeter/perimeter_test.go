package Perimeter

import (
	"testing"
)


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

	t.Run("test Area",func(t *testing.T){
		width := 12
		height := 2 
		get := Area(float64(width),float64(height))
		expect :=24
		certainResult(t,get,float64(expect))
	})

	
}