package httpx

import (
	"io"
	"net/http"
	"strings"
)

type Client struct {
	http *http.Client
}

func NewClient(httpClient *http.Client) Client {
	return Client{
		http: httpClient,
	}
}

func (c Client) MakeRequest(reqMethod, endpoint, reqBody string, headers map[string]string) (string, error) {
	var body io.Reader
	if reqBody != "" {
		body = strings.NewReader(reqBody)
	}

	req, err := http.NewRequest(reqMethod, endpoint, body)
	if err != nil {
		return "", err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respByte), nil
}
