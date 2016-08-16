package codility

import "sort"

func Solution(A []int) int {
  if len(A) == 0 {
    return 0
  }

  result := 1

  sort.Ints(A)

  for i := 1; i < len(A); i++ {
    current := A[i]
    previous := A[i - 1]

    if current != previous {
      result++
    }
  }

  return result
}
