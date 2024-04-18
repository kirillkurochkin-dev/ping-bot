package workerpool

import (
	"fmt"
	"time"
)

type Response struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Error        error
}

func (r *Response) String() string {
	if r.Error != nil {
		return fmt.Sprintf("URL: %s; ERROR, %s;", r.URL, r.Error)
	}

	return fmt.Sprintf("URL: %s; StatusCode: %d; ResponseTie: %s;",
		r.URL, r.StatusCode, r.ResponseTime)
}
