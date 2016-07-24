package codility

import "testing"
import "reflect"

func TestSolution(t *testing.T) {
  cases := []struct {
    inputN int
    inputA []int
    output []int
	}{
    {5, []int{3,4,4,6,1,4,4}, []int{3,2,2,4,2}},
  }

  for _, c := range cases {
    result := Solution(c.inputN, c.inputA)

    if !reflect.DeepEqual(result, c.output) {
      t.Errorf("Solution(%v, %v) == %v, should be %v", c.inputN, c.inputA, result, c.output)
    }
  }
}

func TestIncreaseCounter(t *testing.T) {
  cases := []struct {
    inputA []int
    inputN, inputMax, output int
	}{
    {[]int{0,0,0}, 4, 0, 0},
    {[]int{0,0,0}, 4, 4, 0},
    {[]int{0,0,0}, 1, 0, 1},
    {[]int{0,0,0}, 1, 3, 4},
    {[]int{0,0,0}, 3, 0, 1},
    {[]int{0,0,0}, 3, 3, 4},
  }

  for _, c := range cases {
    inputCopy := c.inputA
    result := IncreaseCounter(&inputCopy, c.inputN, c.inputMax)

    if result != c.output {
      t.Errorf("IncreaseCounter(%v, %v, %v) == %v, should be %v", c.inputA, c.inputN, c.inputMax, result, c.output)
    }
  }
}

func TestMax(t *testing.T) {
  cases := []struct {
    inputL, inputR, output int
	}{
    {1,2,2},
    {2,1,2},
    {2,2,2},
  }

  for _, c := range cases {
    result := Max(c.inputL, c.inputR)

    if result != c.output {
      t.Errorf("Max(%v, %v) == %v, should be %v", c.inputL, c.inputR, result, c.output)
    }
  }
}
