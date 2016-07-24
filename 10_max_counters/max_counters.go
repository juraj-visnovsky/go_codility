package codility

func Solution(N int, A []int) []int {
  // Initialize counters
  counters := make([]int, N)

  for i := range counters {
    counters[i] = 0
  }

  maxCounter := 0
  currentIncrease := 0

  // Operations
  for _, operation := range A {
    if operation == N + 1 {
      currentIncrease = maxCounter
    } else if operation >= 1 {
      maxCounter = Max(maxCounter, IncreaseCounter(&counters, operation, currentIncrease))
    }
  }

  for i := range counters {
    counters[i] = Max(counters[i], currentIncrease)
  }

  return counters
}

func IncreaseCounter(A *[]int, N int, currentIncrease int) int {
  if N >= 1 && N <= len(*A) {
    if (*A)[N - 1] < currentIncrease {
      (*A)[N - 1] = currentIncrease
    }
    (*A)[N - 1]++
    return (*A)[N - 1]
  }
  return 0
}

func Max(left int, right int) int {
  if left > right {
    return left
  }
  return right
}
