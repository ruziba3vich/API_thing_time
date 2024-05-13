package models

type RFC3339 struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth *int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       *int    `json:"year"`
	Hour       *int    `json:"hour"`
	Minute     *int    `json:"minute"`
	Second     *int    `json:"second"`
}

/*

{
  "day_of_week": "Monday",
  "day_of_month": 10,
  "month": "April",
  "year": 2023,
  "hour": 20,
  "minute": 15,
  "second": 20
}

*/
