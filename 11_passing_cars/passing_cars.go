package codility

func Solution(A []int) int {
  passing_cars_count := 0
  opposite_counter := 0

  for _, val := range A {
    if val == 0 {
      opposite_counter++
    } else {
      passing_cars_count += opposite_counter
    }

    if passing_cars_count > 1000000000 {
			return -1
    }
  }

  return passing_cars_count
}
