package Sumarray

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){


    assertCorrectMessage := func(t *testing.T, got, want any) {
        t.Helper()
        differState := reflect.DeepEqual(got,want)
        if !differState {
            t.Errorf("got '%q' want '%q'", got, want)
        }
    }
	

    t.Run("Sum array test", func(t *testing.T) {
		rawArray := [5]int{1,2,3,4,5}
		expect := 15
		input := Sum(rawArray[:])
			assertCorrectMessage(t, input, expect)
		})

    t.Run("SumAll array test",func(t *testing.T){
        get := SumAll([][]int{{3,4},{4,5}})
        expect := []int{7,9}
        assertCorrectMessage(t,get,expect)
    })
}