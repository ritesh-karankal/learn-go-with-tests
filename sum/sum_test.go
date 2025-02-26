package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int {1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{1, 1, 1})
	want := []int{6, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
	
	
func TestSumAllTails(t *testing.T) {

	checksums := func(t testing.TB, got, want []int)  {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Run sum of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9} 
		checksums(t, got, want)
		
	})

	t.Run("Safely run sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{})
		want := []int{2, 0}
		checksums(t, got, want)
	})
}