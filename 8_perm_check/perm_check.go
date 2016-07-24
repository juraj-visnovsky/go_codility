package codility

func Solution(A []int) int {
  presentList := make([]bool, len(A))

  for i := range presentList {
    presentList[i] = false
  }

  for _, val := range A {
    if val >= 1 && val <= len(A) && presentList[val - 1] == false {
      presentList[val - 1] = true
    } else {
      return 0
    }
  }

  return 1
}
