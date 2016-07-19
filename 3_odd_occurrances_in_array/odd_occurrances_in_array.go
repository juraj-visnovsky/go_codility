package codility

func Solution(A []int) int {
  single := 0

  for _, number := range A {
    single ^= number
  }

  return single
}
