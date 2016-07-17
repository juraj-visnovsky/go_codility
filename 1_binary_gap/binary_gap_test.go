package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input, output int
	}{
    {0, 0},
    {1, 0},
    {2, 0},
    {3, 0},
    {4, 0},
    {5, 1},
    {6, 0},
    {7, 0},
    {8, 0},
    {9, 2},
    {20, 1},
    {529, 4},
    {1041, 5},
    {2548924, 2},
    {1247188000, 4},
    {4115884126333, 5},
  }

  for _, c := range cases {
    result := Solution(c.input)
    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }

}
