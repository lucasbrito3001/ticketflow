package rest

type (
	SuccessResponse struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	ErrorResponse struct {
		Error   string `json:"error"`
		Details string `json:"details,omitempty"`
	}
)
