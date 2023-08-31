package hrank

import (
	"math"
)

func StringIsValid() bool {
	s := "abcdefghhgfedecba"
	limit := 2
	var temp int
	mapint := make(map[int32]int)
	for _, x := range s {
		if mapint[x] == 0 {
			mapint[x] = 1
		} else {
			mapint[x]++
		}
	}

	for _, occured := range mapint {
		if temp == 0 {
			temp = occured
			limit--
			continue
		} else if temp != occured {
			if math.Abs(float64(temp-occured)) > 1 {
				return false
			}
			limit--
		}
		if limit < 0 {
			return false
		}
	}
	return true
}

func SubstrCount() int64 {
	s := "abaabbabaaaaabaababbbaababbabaaaabaabaaaaaabbabbaaabbabaabbbababbbaaabbbbbbbaaabbaaabbaaaaaaababbbabbbabbabbaaababaaabaabaabbababbaababbababbbabbbaabababbabbabbbbbbaaabbaabaaabbabababbbaaababbabbbabbbaabbabbaaaabbabaabaaabaabaaabbaababbbbababaabbabaabbbaaabababbaababbbaabbbbabbbababbabaabbababababbababbaabababbbbabaabbabbabbaaabbbaababbaaabaababaaaabaabaabbbbabbababbbbabaababbaabababbbbbbbbbbbbbbbabbaaabaaaaaaaababbbaabaabbaababbaabaaaaaaaaabbaabaabaabaaaaaabbbabbaaaaaaabaaaabbbaaaabbbaaabaaabbbbbaaaabaaabaabbbaaaaaabaaaaabbabaabaabbbbaabbbbaaaababbabaaabbaababbabbbbaaaababbbaababaaabaaabbbbababbabababbaaaabaabaaaaaabbaaabbbbaabbabaaabaabaabbabbaaaaaabaabbbbbbbbabbaababbabbaaababbbaaaabaabaabaabaabbbbbaabbbbbabbbaaabbbabaababbbbaabbbbabaabaaabaabbabbbbaaaabbabbabaabaaabbbbbbbbaaaaabaabaabbabbababaaaabbababbaabbabaabbbbabbbaaaaabbbbabbbbbbaaaaabbbbbabbbbaaabbbbbbabaabbbbabaabbaabbabbabbabbaabababbaaaabbababbbbabaabaaaabbbababbbababaaaabbbaaaaaabbbaaabaaaaabaababaaabbaabbaabbababaabbaabbaaabbbbbababbabaa"
	var counter int64
	var str string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if i < j {
				str = s[i:j]
				if len(str) == 1 {
					counter++
				} else if len(str) == 2 {
					if checkdouble(str[:1], str[1:]) {
						counter++
					}
				} else if len(str) == 3 {
					if checkdouble(str[:1], str[2:]) {
						counter++
					}
				} else if checkPalindrom(s[i:j]) && s[i:j] != "" {
					counter++

				}

			}
		}
	}

	return counter
}

func checkPalindrom(s string) (result bool) {
	mididx := len(s) / 2

	if len(s)%2 == 0 {
		return checkstring(s[:mididx] + s[mididx:])

	} else {
		return checkstring(s[:mididx] + s[mididx+1:])
	}

}
func checkdouble(a, b string) bool {
	return a == b
}
func checkstring(a string) bool {
	var temp int32

	for _, i := range a {
		if temp == 0 {
			temp = i
			continue
		}
		if temp != i {
			return false
		}
	}
	return true
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}
