package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input []int
    output int
	}{
    {[]int{1,2,1}, 2},
    {[]int{1,1,2}, 2},
    {[]int{2,1,1}, 2},
    {[]int{1,4,7,7,8,9,6,8,1,4,6}, 9},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
