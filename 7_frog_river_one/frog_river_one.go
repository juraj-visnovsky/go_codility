package codility

func Solution(X int, A []int) int {
  leaves := make([]int, X)
  for i := range leaves {
    leaves[i] = -1
  }

  for i, val := range A {
    if val <= X && leaves[val - 1] == -1 {
      leaves[val - 1] = i
    }
  }

  min_time := -1

  for i, val := range leaves {
    if leaves[i] == -1 {
      return -1
    }

    if val > min_time {
      min_time = val
    }
  }

  return min_time
}
