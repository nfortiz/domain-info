package utils

var grades map[string]int

func GetGrade(grade string) int {
	grades = map[string]int{
		"A+": 9,
		"A": 8,
		"A-": 7,
		"B": 6,
		"C": 5,
		"D": 4,
		"E": 3,
		"F": 2,
		"M": 1,
		"T": 0,
	}
	resp, ok := grades[grade]
	if !ok {
		return -1
	}
	return resp
}
