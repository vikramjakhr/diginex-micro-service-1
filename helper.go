package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type Data struct {
	Message string `json:"message"`
}

type DiginexMSClient struct {
	baseURL url.URL
	*http.Client
}

// Reverse a string
func (data *Data) reverse() error {
	u, err := url.Parse(*diginex_micro_sercvice_2_base_url)
	if err != nil {
		return err
	}
	c := DiginexMSClient{
		*u,
		&http.Client{},
	}

	requestBody, _ := json.Marshal(data)

	req, err := c.newRequest("POST", "/reverse", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(response, &data)
}

func (c *DiginexMSClient) newRequest(method, requestPath string, body io.Reader) (*http.Request, error) {
	msURL := c.baseURL
	msURL.Path = path.Join(msURL.Path, requestPath)
	req, err := http.NewRequest(method, msURL.String(), body)
	if err != nil {
		return req, err
	}
	req.Header.Add("Content-Type", "application/json")
	return req, err
}


