package countdiffbit

import "crypto/sha256"

func CountDiffBit(s1, s2 string) int {
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	count := 0

	for i := 0; i < 32; i++ {
		if c1[i] != c2[i] {
			count++
		}
	}

	return count
}
