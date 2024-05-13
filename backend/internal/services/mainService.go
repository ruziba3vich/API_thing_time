package services

import (
	"errors"
	"fmt"

	"github.com/time_project/backend/internal/models"
)

func GetResult(time models.RFC3339) (string, error) {
	// var result string
	months := map[string]string{
		"January":   "01",
		"February":  "02",
		"March":     "03",
		"April":     "04",
		"May":       "05",
		"June":      "06",
		"July":      "07",
		"August":    "08",
		"September": "09",
		"October":   "10",
		"November":  "11",
		"December":  "12",
	}

	if len(time.DayOfWeek) > 0 &&
		time.DayOfMonth != nil &&
		len(time.Month) > 0 &&
		time.Year != nil &&
		time.Hour != nil &&
		time.Minute != nil &&
		time.Second != nil {

		return fmt.Sprintf("%d-%s-%dT%d:%d:%d-05:00",
			*time.Year,
			months[time.Month],
			*time.DayOfMonth,
			*time.Hour,
			*time.Minute,
			*time.Second), nil
	}
	return "", errors.New("invalid data")
}

/// # 2024-05-13T15:30:00-07:00

/*
type RFC3339 struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth *int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       *int    `json:"year"`
	Hour       *int    `json:"hour"`
	Minute     *int    `json:"minute"`
	Second     *int    `json:"second"`
}
*/
