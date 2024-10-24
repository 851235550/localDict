package services

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	} else {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}
}
