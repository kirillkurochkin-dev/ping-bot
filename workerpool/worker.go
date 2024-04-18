package workerpool

import (
	"net/http"
	"time"
)

type Worker struct {
	client *http.Client
}

func NewWorker(timeout time.Duration) *Worker {
	return &Worker{client: &http.Client{
		Timeout: timeout,
	}}
}

func (w Worker) makeRequest(j Job) Response {
	startTime := time.Now()
	response := Response{}
	response.URL = j.URL

	data, err := w.client.Get(j.URL)
	if err != nil {
		response.Error = err
		return response
	}

	response.StatusCode = data.StatusCode
	response.ResponseTime = time.Since(startTime)

	return response
}
