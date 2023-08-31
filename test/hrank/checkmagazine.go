package hrank

import "fmt"

func CheckMagazine(magazine []string, note []string) {
	// Write your code here
	yesno := true
	flag := make(map[string]int)
	for _, y := range note {
		_, ok := flag[y]
		if !ok {
			flag[y] = 0

		}
		for _, x := range magazine {
			if x == y {
				flag[y]++
			}
		}
	}
	for _, y := range flag {
		if y == 0 || y > 1 {
			yesno = false
		}
	}
	fmt.Println(flag)
	if yesno {
		fmt.Println("YES")

	} else {
		fmt.Println("NO")

	}

}
