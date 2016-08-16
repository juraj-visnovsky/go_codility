package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    input []int
    output int
	}{
    {[]int{2,1,1,2,3,1}, 3},
    {[]int{9,8,7,6,5,4,3,2,1}, 9},
    {[]int{0,5,4,7,1,0,2,4,4,2,1,5,8,1}, 7},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
