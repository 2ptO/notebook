package chapter1

// Given two strings, write a method to find whether one string
// is a permutation of the other.

// ArePermutations implementation
func ArePermutations(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	chars := make(map[rune]int)
	for _, ch := range str1 {
		chars[ch]++
	}

	for _, ch := range str2 {
		if _, ok := chars[ch]; ok {
			chars[ch]--
			if chars[ch] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// type text struct {
// 	val string
// }

// func (t text) isPermutationOf(otherText string) bool {
// 	if len(t.val) != len(otherText) {
// 		return false
// 	}

// 	s1 := strings.Split(t.val, "")
// 	sort.Strings(s1)
// 	s2 := strings.Split(otherText, "")
// 	sort.Strings(s2)

// 	return strings.Join(s1, "") == strings.Join(s2, "")
// }

// func main() {
// 	t1 := text{"hellow"}
// 	fmt.Println(t1.isPermutationOf("wohell"))
// 	fmt.Println(t1.isPermutationOf("wohelowng"))
// }
