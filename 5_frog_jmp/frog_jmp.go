package codility

func Solution(X int, Y int, D int) int {
  length := Y - X

  if length % D == 0 {
    return (length / D)
  }

  return (length / D) + 1
}
