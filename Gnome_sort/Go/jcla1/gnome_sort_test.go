package gnomeSort

import "testing"

func TestGnomeSort(t *testing.T) {
    arrs := [][]int{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, []int{2, 3, 7, -5, -1, 4, -10}}

    var result []int

    for _, arr := range arrs {
        result = GnomeSort(arr)
        if !isSorted(result) {
            t.Errorf("GnomeSort(%v) = %v", arr, result)
        }
    }
}

func isSorted(arr []int) bool {
    for i := 1; i < len(arr); i++ {
        if arr[i-1] > arr[i] {
            return false
        }
    }

    return true
}
