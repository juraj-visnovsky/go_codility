package codility

func Solution(A []int, K int) []int {
    arrayLength := len(A)

    if arrayLength == 0 {
      return []int{}
    }

    rotation := Rotation(K, arrayLength)

    begin := A[(arrayLength - rotation):]
    end := A[:(arrayLength - rotation)]

    rotatedArray := append(begin, end...)

    return rotatedArray
}

func Rotation(K int, arrayLength int) int {
  return K % arrayLength
}
