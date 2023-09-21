package helpers

type HttpResponse struct {
	Status interface{} `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func WebResponse(status interface{}, data interface{}) *HttpResponse {
	res := &HttpResponse{
		Status: status,
		Data:   data,
	}
	return res
}
