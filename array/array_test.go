package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Testing sum of any numbers in array", func(t *testing.T) {
		numbers := []int{1, 4, 5}
		got := Sum(numbers)
		expected := 10

		if got != expected {
			t.Errorf("Expected %d, got %d, for values %v", expected, got, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Testing sum of args array", func(t *testing.T) {
		got := SumAll([]int{1, 4, 5}, []int{5, 5})
		expected := []int{10, 10}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
}

func TestSumTails(t *testing.T) {
	t.Run("Testing sum of non array tails", func(t *testing.T) {
		got := SumTails([]int{0, 9}, []int{2, 5, 6})
		expected := []int{9, 11}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})
	t.Run("Testing sum of empty array tails", func(t *testing.T) {
		got := SumTails([]int{}, []int{0})
		expected := []int{0, 0}
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

}
