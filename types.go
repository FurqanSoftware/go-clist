package clist

type ListContestsInput struct {
	Limit    int
	Offset   int
	StartGte string
	OrderBy  string
}

type ListContestsResult struct {
	Meta    Meta     `json:"meta"`
	Objects []Object `json:"objects"`
}

type Meta struct {
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
	TotalCount int `json:"total_count"`
}

type Object struct {
	ID       int      `json:"id"`
	Event    string   `json:"event"`
	Href     string   `json:"href"`
	Start    string   `json:"start"`
	Duration int      `json:"duration"`
	End      string   `json:"end"`
	Resource Resource `json:"resource"`
}

type Resource struct {
	ID   int    `json:"id"`
	Icon string `json:"icon"`
	Name string `json:"name"`
}
