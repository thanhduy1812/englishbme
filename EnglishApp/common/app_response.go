package common

type GTDError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GTDResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Error  *GTDError   `json:"error,omitempty"`
}

func SuccessResponse(data, paging, filter interface{}) *GTDResponse {
	return &GTDResponse{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data interface{}) *GTDResponse {
	return &GTDResponse{
		Data:   data,
		Paging: nil,
		Filter: nil,
	}
}
