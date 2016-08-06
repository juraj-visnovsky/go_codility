package codility

func Solution(S string, P []int, Q []int) []int {
  nucleotides := map[rune]int {
    'A': 1,
    'C': 2,
    'G': 3,
    'T': 4,
  }

  last_positions := map[rune][]int {
    'A': make([]int, len(S)),
    'C': make([]int, len(S)),
    'G': make([]int, len(S)),
    'T': make([]int, len(S)),
  }

  for _, array := range last_positions {
    for i := 0; i < len(S); i++ {
      array[i] = -1
    }
  }

  for i, c := range S {
    for key, array := range last_positions {
      if c == key {
        array[i] = i
      } else if i > 0 {
        array[i] = array[i - 1]
      }
    }
  }

  results := make([]int, len(P))

  for i, p := range P {
    q := Q[i]

    if last_positions['A'][q] >= p {
      results[i] = nucleotides['A']
    } else if last_positions['C'][q] >= p {
      results[i] = nucleotides['C']
    } else if last_positions['G'][q] >= p {
      results[i] = nucleotides['G']
    } else if last_positions['T'][q] >= p {
      results[i] = nucleotides['T']
    }
  }

  return results
}
