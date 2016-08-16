package codility

import "sort"

func Solution(A []int) int {
  if len(A) < 3 {
    return 0
  }

  sort.Ints(A)

  first_product := A[0] * A[1] * A[len(A) - 1]
  last_product := A[len(A) - 1] * A[len(A) - 2] * A[len(A) - 3]

  if first_product > last_product {
    return first_product
  } else {
    return last_product
  }
}
