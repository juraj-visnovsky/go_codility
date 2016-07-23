package codility

func Solution(A []int) int {
  expected_sum := len(A)
  actual_sum := 0

  for i, val := range A {
    expected_sum += (i + 1)
    actual_sum += val
  }

  return (expected_sum - actual_sum + 1)
}
