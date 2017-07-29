// Given two strings, write a method to find whether one string
// is a permutation of the other.

package main

import (
	"fmt"
	"sort"
	"strings"
)

type text struct {
	val string
}

func (t text) isPermutationOf(otherText string) bool {
	if len(t.val) != len(otherText) {
		return false
	}

	s1 := strings.Split(t.val, "")
	sort.Strings(s1)
	s2 := strings.Split(otherText, "")
	sort.Strings(s2)

	return strings.Join(s1, "") == strings.Join(s2, "")
}

func main() {
	t1 := text{"hellow"}
	fmt.Println(t1.isPermutationOf("wohell"))
	fmt.Println(t1.isPermutationOf("wohelowng"))
}
