package wrapper

import (
	"encoding/json"
	"net/http"
)

// Meta model

// HTTPResponse format
type HTTPResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

// NewHTTPResponse for create common response
func NewHTTPResponse(code int, message string, data interface{}, err error) *HTTPResponse {
	resp := new(HTTPResponse)
	resp.Data = data
	resp.Code = code
	resp.Message = message
	if err == nil && code < http.StatusBadRequest {
		resp.Success = true
	}
	return resp
}

// JSON for set http JSON response (Content-Type: application/json) with parameter is http response writer
func (resp *HTTPResponse) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}
