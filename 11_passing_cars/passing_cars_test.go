package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    input []int
    output int
	}{
    {[]int{0,1,0,1,1}, 5},
    {[]int{0,1,0,0,1}, 4},
    {[]int{1,0,1,0,1}, 3},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
