package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
    inputA, inputB, inputK, output int
	}{
    {6, 11, 2, 3},
    {6, 11, 12, 0},
    {6, 12, 2, 4},
    {6, 6, 2, 1},
    {5, 6, 2, 1},
  }

  for _, c := range cases {
    result := Solution(c.inputA, c.inputB, c.inputK)

    if result != c.output {
      t.Errorf("Solution(%v, %v, %v) == %v, should be %v", c.inputA, c.inputB, c.inputK, result, c.output)
    }
  }
}
