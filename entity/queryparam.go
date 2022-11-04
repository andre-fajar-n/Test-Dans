package entity

import (
	"fmt"
	"net/http"
)

func (e *JobListRequest) GenerateQueryParam(req *http.Request) *http.Request {
	q := req.URL.Query()
	if e.Description != "" {
		q.Set("description", e.Description)
	}

	if e.Location != "" {
		q.Set("location", e.Location)
	}

	if e.FullTime != nil {
		q.Set("full_time", fmt.Sprintf("%v", *e.FullTime))
	}

	if e.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", e.Page))
	}

	req.URL.RawQuery = q.Encode()

	return req
}
