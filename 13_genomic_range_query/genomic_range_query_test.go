package codility

import "testing"
import "reflect"

func TestSolution(t *testing.T) {
  cases := []struct {
    inputS string
    inputP, inputQ, output []int
	}{
    {"CAGCCTA", []int{2,5,0}, []int{4,5,6}, []int{2,4,1}},
  }

  for _, c := range cases {
    result := Solution(c.inputS, c.inputP, c.inputQ)

    if !reflect.DeepEqual(result, c.output) {
      t.Errorf("Solution(%v, %v, %v) == %v, should be %v", c.inputS, c.inputP, c.inputQ, result, c.output)
    }
  }
}
