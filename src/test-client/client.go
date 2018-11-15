package main

import (
	"bytes"
	"net/http"
	"time"
)

type WebClient struct {
	serverUrl string
}

func (c *WebClient) Write(url string) (int, error) {
	req, err := http.NewRequest("GET", c.serverUrl, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Date", time.Now().UTC().Format(http.TimeFormat))

	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}

func (c *WebClient) Close() {

}
