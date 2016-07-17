package codility

import "testing"
import "reflect"

func TestSolution(t *testing.T) {
  cases := []struct {
		inputA []int
    inputK int
    output []int
	}{
    {[]int{}, 0, []int{}},
    {[]int{}, 1, []int{}},
    {[]int{3, 8, 9, 7, 6}, 0, []int{3, 8, 9, 7, 6}},
    {[]int{3, 8, 9, 7, 6}, 1, []int{6, 3, 8, 9, 7}},
    {[]int{3, 8, 9, 7, 6}, 2, []int{7, 6, 3, 8, 9}},
    {[]int{3, 8, 9, 7, 6}, 3, []int{9, 7, 6, 3, 8}},
    {[]int{3, 8, 9, 7, 6}, 4, []int{8, 9, 7, 6, 3}},
    {[]int{3, 8, 9, 7, 6}, 5, []int{3, 8, 9, 7, 6}},
  }

  for _, c := range cases {
    result := Solution(c.inputA, c.inputK)
    if !reflect.DeepEqual(result, c.output) {
      t.Errorf("Solution(%v, %v) == %v, should be %v", c.inputA, c.inputK, result, c.output)
    }
  }
}

func TestRotation(t *testing.T) {
  cases := []struct {
		inputK, inputLength, output int
	}{
    {0, 5, 0},
    {1, 5, 1},
    {2, 5, 2},
    {3, 5, 3},
    {4, 5, 4},
    {5, 5, 0},
    {6, 5, 1},
    {7, 5, 2},
    {8, 5, 3},
    {9, 5, 4},
    {10, 5, 0},
  }

  for _, c := range cases {
    result := Rotation(c.inputK, c.inputLength)
    if result != c.output {
      t.Errorf("Rotation(%v, %v) == %v, should be %v", c.inputK, c.inputLength, result, c.output)
    }
  }
}
