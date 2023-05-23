package http

type Response struct {
	Action string `json:"action"`
	X      string `json:"x"`
	Y      string `json:"y"`
	Answer string `json:"answer"`
	Cashed bool   `json:"cashed"`
}
