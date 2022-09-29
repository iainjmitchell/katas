package app

func ToRoman(mappings map[int]string, arabic int) string {
	return mappings[arabic]
}
