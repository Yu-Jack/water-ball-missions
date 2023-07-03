package domain

type Patient struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Gender string  `json:"gender"`
	Age    int     `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Cases  []Case  `json:"cases"`
}
