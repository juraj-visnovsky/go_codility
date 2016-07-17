package codility

func Solution(A []int, K int) []int {
    arrayLength := len(A)

    if arrayLength == 0 {
      return []int{}
    }

    rotation := Rotation(K, arrayLength)

    beginning := A[(arrayLength - rotation):]
    ending := A[:(arrayLength - rotation)]

    rotatedArray := append(beginning, ending...)

    return rotatedArray
}

func Rotation(K int, arrayLength int) int {
  return K % arrayLength
}
