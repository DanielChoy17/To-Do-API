package views

type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequst struct {
	Name string `json:"name"`
	Todo string `json:"todo"`
}