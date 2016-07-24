package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input []int
    output int
	}{
    {[]int{1,3,1}, 2},
    {[]int{1}, 2},
    {[]int{1,3,6,4,1,2}, 5},
    {[]int{100}, 1},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
