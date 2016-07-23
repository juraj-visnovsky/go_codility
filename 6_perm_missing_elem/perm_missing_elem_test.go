package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input []int
    output int
	}{
    {[]int{2,3,4,5,6}, 1},
    {[]int{1,3,4,5,6}, 2},
    {[]int{1,2,4,5,6}, 3},
    {[]int{1,2,3,5,6}, 4},
    {[]int{1,2,3,4,6}, 5},
    {[]int{1,2,3,4,5}, 6},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
