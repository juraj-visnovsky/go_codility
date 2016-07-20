package codility

import "math"

func Solution(A []int) int {
  arrayLength := len(A)
  sums := make([]int, arrayLength)
  sums[0] = A[0]

  for i := 1; i < arrayLength; i++ {
    sums[i] = A[i] + sums[i - 1]
  }

  totalSum := sums[arrayLength-1]

  minDifference := math.MaxInt64

  for _, currentSum := range sums {
    minDifference = Min(minDifference, Abs(currentSum - (totalSum - currentSum)))
  }

  return minDifference
}

func Min(left, right int) int {
  if left < right {
    return left
  }
  return right
}

func Abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
