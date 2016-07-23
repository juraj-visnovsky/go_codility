package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    inputX int
		inputA []int
    output int
	}{
    {5, []int{1,3,1,4,2,3,5,4}, 6},
    {7, []int{1,3,1,4,2,3,5,4,2,4,2,7,1,2,1,4,6,3,2}, 16},
    {8, []int{8,3,2,3,1,2,3,5,4,3,6,8,7,5,4}, 12},
    {3, []int{1,2,1}, -1},
    {1, []int{2}, -1},
  }

  for _, c := range cases {
    result := Solution(c.inputX, c.inputA)

    if result != c.output {
      t.Errorf("Solution(%v, %v) == %v, should be %v", c.inputX, c.inputA, result, c.output)
    }
  }
}
