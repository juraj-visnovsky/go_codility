package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		inputX, inputY, inputD, output int
	}{
    {10, 20, 2, 5},
    {10, 21, 2, 6},
  }

  for _, c := range cases {
    result := Solution(c.inputX, c.inputY, c.inputD)

    if result != c.output {
      t.Errorf("Solution(%v, %v, %v) == %v, should be %v", c.inputX, c.inputY, c.inputD, result, c.output)
    }
  }
}
