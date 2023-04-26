package models

type AccessLog struct {
	Id      string  `json:"id"`
	URL     string  `json:"url"`
	Path    string  `json:"path"`
	Ua      string  `json:"ua"`
	Status  int     `json:"status"`
	Elapsed float64 `json:"Elapsed"`
}
