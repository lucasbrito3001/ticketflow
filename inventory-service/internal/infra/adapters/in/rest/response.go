package rest

type (
	SuccessResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Error     string `json:"error"`
		ErrorCode string `json:"error_code,omitempty"`
		Details   string `json:"details,omitempty"`
	}
)
