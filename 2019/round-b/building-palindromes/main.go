package main

import (
	"fmt"
	"os"
)

// Testcase contains the blocks and questions.
type testcase struct {
	blocks string
	qs     []qstn
}

// Qstn contains the left and right one-based block numbers
// for the palindrome question about the blocks.
type qstn struct {
	left  int
	right int
}

// Run returns the number of questions which reference valid
// palindrome subsets of blocks.
func (tc *testcase) run() int {
	count := 0
	ps := newPrefixSums(tc.blocks)
	for _, q := range tc.qs {
		if ps.check(q.left-1, q.right) {
			count++
		}
	}
	return count
}

// CharSums contains the count of all characters.
// It is essentially indexed by character with 'A' at 0.
type charSums []int

// NewCharSums creates a charSums with an entry for each letter.
func newCharSums() charSums {
	return make([]int, 26)
}

// Add increments the character count for the given character rune.
func (cs charSums) add(r rune) {
	cs[r-'A']++
}

// PrefixSums contains the charSums for each of a string's prefixes.
// A prefix of string s is s[0:i] where i goes from 0 to the length
// of the string.
type prefixSums []charSums

// NewPrefixSums creates the prefixSums for a given string.
func newPrefixSums(s string) prefixSums {
	ps := prefixSums(make([]charSums, len(s)+1))
	ps[0] = newCharSums()
	for i := 1; i <= len(s); i++ {
		ps[i] = newCharSums()
		copy(ps[i], ps[i-1])
		ps[i].add([]rune(s)[i-1])
	}

	return ps
}

// Check for possible palindrome for blocks[i:j] (length j - i)
func (ps prefixSums) check(i, j int) bool {
	odds := 0
	for csi := range ps[j] {
		if (ps[j][csi]-ps[i][csi])%2 != 0 {
			odds++
			if odds > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	// Get number of testcases
	input := os.Stdin
	T := 0
	fmt.Fscanf(input, "%d\n", &T)

	// Get input for each testcase
	for t := 1; t <= T; t++ {
		T := testcase{}

		// Get number of blocks and number of questions
		blocksLen := 0
		Q := 0
		fmt.Fscanf(input, "%d %d\n", &blocksLen, &Q)

		T.qs = make([]qstn, Q)

		// Get blocks string
		fmt.Scanf("%s\n", &T.blocks)

		// Get input for each question
		for i := range T.qs {
			Qi := qstn{}
			fmt.Fscanf(input, "%d %d\n", &Qi.left, &Qi.right)
			T.qs[i] = Qi
			// fmt.Printf("Question %d: %d %d\n", j+1, q.left, q.right)
		}

		// Run testcase
		fmt.Printf("Case #%d: %d\n", t, T.run())
	}
}
