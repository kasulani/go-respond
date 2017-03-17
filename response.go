package respond

import (
	"encoding/json"
	"net/http"
)

// Response is our HTTP response
type Response struct {
	Writer  http.ResponseWriter
	Headers map[string]string
}

// NewResponse create and return an
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		Writer: w,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
	}
}

// DeleteHeader deletes a single provided header
func (resp *Response) DeleteHeader(key string) *Response {
	resp.Writer.Header().Del(key)
	return resp
}

// AddHeader adds a single header to the response
func (resp *Response) AddHeader(key string, value string) *Response {
	resp.Writer.Header().Add(key, value)
	return resp
}

// WriteResponse writes the HTTP response status, headers and body
func (resp *Response) writeResponse(code int, v interface{}) error {
	resp.writeStatusCode(code)

	if len(resp.Headers) > 0 {
		resp.writeHeaders()
	}

	if v != nil {
		body, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}

		if _, err := resp.Writer.Write(body); err != nil {
			panic(err)
		}
	}
	return nil
}

func (resp *Response) writeHeaders() {
	for key, value := range resp.Headers {
		resp.Writer.Header().Set(key, value)
	}
}

func (resp *Response) writeStatusCode(code int) {
	resp.Writer.WriteHeader(code)
}
