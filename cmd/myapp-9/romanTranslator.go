package main

import (
	"fmt"
)

func romanToArabic(numeral string) int {
	romanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	arabicVals := make([]int, len(numeral)+1)

	for index, digit := range numeral {
		if val, present := romanMap[digit]; present {
			arabicVals[index] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	total := 0
	for index := 0; index < len(numeral); index++ {
		if arabicVals[index] >= arabicVals[index+1] {
			total += arabicVals[index]
		} else {
			total += arabicVals[index+1] - arabicVals[index]
			index++
		}
	}

	return total
}

func main() {
	fmt.Println("MCLX is", romanToArabic("MCLX"))
	fmt.Println("MCLXIII is", romanToArabic("MCLXIII"))
	fmt.Println("MCMXCIX is ", romanToArabic("MCMXCIX"))
	fmt.Println("MCMXCVIII is ", romanToArabic("MCMXCVIII"))
	fmt.Println("MMXVIII is ", romanToArabic("MMXVIII"))
	fmt.Println("MCMZ is", romanToArabic("MCMZ"))
}
