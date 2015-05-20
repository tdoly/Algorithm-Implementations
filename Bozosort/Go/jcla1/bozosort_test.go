package bozosort

import "testing"

func TestBozosort(t *testing.T) {
    arrs := [][]int{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, []int{2, 3, 7, -5, -1, 4, -10}}

    var result []int

    for _, arr := range arrs {
        result = Bozosort(arr)
        if !isSorted(result) {
            t.Errorf("Bozosort(%v) = %v", arr, result)
        }
    }
}