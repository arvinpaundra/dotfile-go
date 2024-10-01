package format

import "net/http"

type (
	Meta struct {
		Status  string `json:"status"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	BaseResponse struct {
		Meta   Meta `json:"meta"`
		Data   any  `json:"data"`
		Errors any  `json:"errors,omitempty"`
	}
)

// 200 - OK
func SuccessOK(message string, data any) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "success",
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	}
}

// 201 - Created
func SuccessCreated(message string, data any) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "success",
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	}
}

// 400 - Bad Request
func BadRequest(message string, errors any) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusBadRequest,
			Message: message,
		},
		Errors: errors,
	}
}

// 401 - Unauthorized
func Unauthorized(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusUnauthorized,
			Message: message,
		},
	}
}

// 403 - Forbidden
func Forbidden(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusForbidden,
			Message: message,
		},
	}
}

// 404 - Not Found
func NotFound(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusNotFound,
			Message: message,
		},
	}
}

// 409 - Conflict
func Conflict(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusConflict,
			Message: message,
		},
	}
}

// 422 - Unprocessable Entity
func UnprocessableEntity(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusUnprocessableEntity,
			Message: message,
		},
	}
}

// 500 - Internal Server Error
func InternalServerError(message string) BaseResponse {
	return BaseResponse{
		Meta: Meta{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: message,
		},
	}
}
