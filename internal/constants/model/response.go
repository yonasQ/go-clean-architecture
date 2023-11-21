package model

type Response struct {
	// OK is only true if the request was successful.
	OK bool `json:"ok"`
	// MetaData contains additional data like filtering, pagination, etc.
	MetaData *MetaData `json:"meta_data,omitempty"`
	// Data contains the actual data of the response.
	Data interface{} `json:"data,omitempty"`
	// Error contains the error detail if the request was not successful.
	Error *ErrorResponse `json:"error,omitempty"`
}

type MetaData struct {
	// Page  is used to identify the starting point to return rows from a result set
	Page    int `form:"page" json:"page,omitempty"`
	// PerPage is the number of records to return after filtering
	PerPage int `form:"per_page" json:"per_page,omitempty"`
	// Sort is used to sort the result set in ascending or descending order
	Sort string `form:"sort" json:"sort,omitempty"`
	// Total is the total number of data without pagination
	Total int `json:"total"`
	// Extra contains other response specific data
	Extra interface{} `json:"extra,omitempty"`
}

type ErrorResponse struct {
	// Code is the error code. It is not status code
	Code int `json:"code"`
	// Message is the error message.
	Message string `json:"message,omitempty"`
	// Description is the error description.
	Description string `json:"description,omitempty"`
	// StackTrace is the stack trace of the error.
	// It is only returned for debugging
	StackTrace string `json:"stack_trace,omitempty"`
	// FieldError is the error detail for each field, if available that is.
	FieldError []FieldError `json:"field_error,omitempty"`
}

type FieldError struct {
	// Name is the name of the field that caused the error.
	Name string `json:"name"`
	// Description is the error description for this field.
	Description string `json:"description"`
}
