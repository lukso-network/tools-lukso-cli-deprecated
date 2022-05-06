package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type HttpClient struct {
	BaseUrl string
	Client  *http.Client
}

func NewHttpClient(baseUrl string) HttpClient {
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	return HttpClient{
		BaseUrl: baseUrl,
		Client:  client,
	}
}

func (c *HttpClient) getUrl(path string) string {
	return fmt.Sprintf("http://%v/%v", c.BaseUrl, path)
}

func (c *HttpClient) Post(body interface{}, response interface{}, token string, path string) (int, error) {
	buffer, err := json.Marshal(body)

	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("POST", c.getUrl(path), bytes.NewBuffer(buffer))

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	if err != nil {
		return -1, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return -1, err
	}

	log.Println(string(b))

	if res.StatusCode >= 300 {
		return res.StatusCode, fmt.Errorf("api responded with %v: ", res.StatusCode)
	}

	err = json.Unmarshal(b, response)

	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil
}

func (c *HttpClient) Put(body interface{}, response interface{}, token string, path string) (int, error) {
	buffer, err := json.Marshal(body)

	if err != nil {
		return -1, err
	}

	req, err := http.NewRequest("PUT", c.getUrl(path), bytes.NewBuffer(buffer))

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	if err != nil {
		return -1, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return -1, err
	}

	if res.StatusCode >= 300 {
		return res.StatusCode, fmt.Errorf("api responded with %v: ", res.StatusCode)
	}

	err = json.Unmarshal(b, response)

	if err != nil {
		return -1, err
	}

	return res.StatusCode, nil
}

func (c *HttpClient) Get(response interface{}, token string, path string) (int, error) {
	req, err := http.NewRequest("GET", c.getUrl(path), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))

	if err != nil {
		return -1, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return -1, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return -1, err
	}

	if res.StatusCode >= 300 {
		return res.StatusCode, fmt.Errorf("api responded with %v: %v", res.StatusCode, string(b))
	}

	if response != nil {
		err = json.Unmarshal(b, response)
		if err != nil {
			return -1, err
		}
	}

	return res.StatusCode, nil
}
