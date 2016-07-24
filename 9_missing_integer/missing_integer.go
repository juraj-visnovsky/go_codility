package codility

func Solution(A []int) int {
  missing := make([]bool, len(A))

  for i := range missing {
    missing[i] = true
  }

  for _, val := range A {
    if val >= 1 && val <= len(A) {
      missing[val - 1] = false
    }
  }

  for i, val := range missing {
    if val == true {
      return i + 1
    }
  }

  return len(A) + 1
}
