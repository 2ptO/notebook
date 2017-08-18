package chapter1

import (
	"testing"
)

func TestArePermutations(t *testing.T) {
	type testCase struct {
		str1, str2 string
		expected   bool
	}
	testCases := []testCase{testCase{"hellow", "wohell", true},
		testCase{"jelly", "belly", false},
		testCase{"equal", "unequal", false},
		testCase{"neith92hgp", "jsdjtp1-4", false},
		testCase{"anna", "nana", true},
	}

	for _, tCase := range testCases {
		actual := ArePermutations(tCase.str1, tCase.str2)
		t.Logf("Testing permuations. %v %v\n", tCase.str1, tCase.str2)
		if actual != tCase.expected {
			t.Fatalf("Permutations computation failed for %s and %s, expected %t, actual %t\n",
				tCase.str1, tCase.str2, tCase.expected, actual)

		}
	}

}
