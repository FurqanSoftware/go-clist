package clist

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	DefaultBaseURL = "https://clist.by/api/v1"
)

type Client struct {
	BaseURL string

	Username string
	APIKey   string
}

func (c Client) endpoint(path string) string {
	baseurl := c.BaseURL
	if baseurl == "" {
		baseurl = DefaultBaseURL
	}
	return baseurl + path
}

func (c Client) authorize(data url.Values) {
	if c.Username == "" && c.APIKey == "" {
		return
	}
	data.Set("username", c.Username)
	data.Set("api_key", c.APIKey)
}

func (c Client) ListContests(input ListContestsInput) (*ListContestsResult, error) {
	data := url.Values{}
	c.authorize(data)
	if input.Limit != 0 {
		data.Set("limit", strconv.Itoa(input.Limit))
	}
	if input.Offset != 0 {
		data.Set("offset", strconv.Itoa(input.Offset))
	}
	if input.StartGte != "" {
		data.Set("start__gte", input.StartGte)
	}
	if input.OrderBy != "" {
		data.Set("order_by", input.OrderBy)
	}
	resp, err := http.Get(c.endpoint("/contest/?") + data.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, Error{StatusCode: resp.StatusCode}
	}

	body, _ := ioutil.ReadAll(resp.Body)
	result := ListContestsResult{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
