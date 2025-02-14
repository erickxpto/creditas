package utils

import "time"

func CalculateAge(birthday time.Time) int {
	//now := time.Now()
	now := time.Date(2025, time.February, 13, 0, 0, 0, 0, time.UTC)
	age := now.Year() - birthday.Year()

	if now.YearDay() < birthday.YearDay() {
		age--
	}
	return age
}
