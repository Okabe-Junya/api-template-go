package models

import "time"

// Response represents the standard API response format
type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Success bool        `json:"success"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(err interface{}) *Response {
	return &Response{
		Success: false,
		Error:   err,
	}
}

// PaginationResponse represents a paginated response format
type PaginationResponse struct {
	Items      interface{} `json:"items"`
	TotalCount int64       `json:"totalCount"`
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalPages int         `json:"totalPages"`
}

// NewPaginationResponse creates a new paginated response
func NewPaginationResponse(items interface{}, totalCount int64, page, pageSize int) *PaginationResponse {
	totalPages := int(totalCount) / pageSize
	if int(totalCount)%pageSize > 0 {
		totalPages++
	}

	return &PaginationResponse{
		Items:      items,
		TotalCount: totalCount,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}
}

// MetaData represents metadata for API responses
type MetaData struct {
	Timestamp   time.Time `json:"timestamp"`
	RequestID   string    `json:"requestId,omitempty"`
	ServiceName string    `json:"serviceName,omitempty"`
	Version     string    `json:"version,omitempty"`
}

// ResponseWithMeta represents a response format with metadata
type ResponseWithMeta struct {
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    *MetaData   `json:"meta"`
	Success bool        `json:"success"`
}

// NewResponseWithMeta creates a new response with metadata
func NewResponseWithMeta(success bool, data interface{}, err interface{}, requestID string) *ResponseWithMeta {
	return &ResponseWithMeta{
		Success: success,
		Data:    data,
		Error:   err,
		Meta: &MetaData{
			Timestamp:   time.Now().UTC(),
			RequestID:   requestID,
			ServiceName: "api-service", // TODO: Get from environment variable
			Version:     "v1.0.0",      // TODO: Inject during build
		},
	}
}
