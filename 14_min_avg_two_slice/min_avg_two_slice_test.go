package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    input []int
    output int
	}{
    {[]int{4,2,2,5,1,5,8}, 1},
    {[]int{5,1,5,4,4,8,1}, 0},
    {[]int{9,8,7,2,5,4,6,4}, 3},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
