package codility

func Solution(A []int) int {
  result := -1
  min_average := float64(10001)

  for i := 0; i < (len(A) - 1); i++ {
    average := float64(A[i] + A[i + 1]) / 2.0

    if average < min_average {
      min_average = average
      result = i
    }

    if i < (len(A) - 2) {
      average = float64(A[i] + A[i + 1] + A[i + 2]) / 3.0
      
      if average < min_average {
        min_average = average
        result = i
      }
    }
  }

  return result
}
