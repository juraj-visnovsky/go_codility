package codility

// import "fmt"
import "strings"
import "strconv"

func Solution(N int) int {
  binary := strconv.FormatInt(int64(N), 2)
  // fmt.Println(N, binary)

  maxGapLength := 0

  gaps := strings.Split(binary, "1")

  for i := 0; i < len(gaps); i++ {
    currentGapLength := len(gaps[i])

    if (i + 1) < len(gaps) {
      if currentGapLength > maxGapLength {
        maxGapLength = currentGapLength
      }
    }
  }

  return maxGapLength
}
