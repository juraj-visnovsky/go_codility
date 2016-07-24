package codility

import "testing"

func TestSolution(t *testing.T) {
  cases := []struct {
		input []int
    output int
	}{
    {[]int{1,2,3}, 0},
    {[]int{3,2,1,4,3}, 1},
    {[]int{2,12,1}, 11},
    {[]int{0,100}, 100},
  }

  for _, c := range cases {
    result := Solution(c.input)

    if result != c.output {
      t.Errorf("Solution(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}

func TestMin(t *testing.T) {
  cases := []struct {
    inputL, inputR, output int
	}{
    {1,2,1},
    {2,1,1},
    {2,2,2},
  }

  for _, c := range cases {
    result := Min(c.inputL, c.inputR)

    if result != c.output {
      t.Errorf("Min(%v, %v) == %v, should be %v", c.inputL, c.inputR, result, c.output)
    }
  }
}

func TestAbs(t *testing.T) {
  cases := []struct {
    input, output int
	}{
    {1,1},
    {-1,1},
    {0,0},
  }

  for _, c := range cases {
    result := Abs(c.input)

    if result != c.output {
      t.Errorf("Abs(%v) == %v, should be %v", c.input, result, c.output)
    }
  }
}
