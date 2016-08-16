package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    input []int
    output int
	}{
    {[]int{-3,1,2,-2,5,6}, 60},
    {[]int{-5,1,2,-3,5,6}, 90},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
