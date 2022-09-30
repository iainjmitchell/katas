package app

import (
	"sort"
	"strings"
)

func ToRoman(mappings map[int]string, arabic int) string {
	var romanNumeral string
	for _, arabicMapping := range highestArabicFirst(mappings) {
		occurrences := arabic / arabicMapping
		if occurrences >= 1 {
			romanNumeral += strings.Repeat(mappings[arabicMapping], occurrences)
			arabic -= arabicMapping * occurrences
		}
	}
	return romanNumeral
}

func highestArabicFirst(mappings map[int]string) []int {
	highestArabicFirst := make([]int, 0, len(mappings))
	for key := range mappings {
		highestArabicFirst = append(highestArabicFirst, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(highestArabicFirst)))
	return highestArabicFirst
}
