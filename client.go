package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

type PublicIPv4AddressGetter interface {
	Get(context.Context) (net.IP, error)
}

type AWSCheckIPClient struct {
	// Modify this field when testing the client
	URL    string
	client *http.Client
}

func NewAWSCheckIPClient() *AWSCheckIPClient {
	return &AWSCheckIPClient{
		URL:    "http://checkip.amazonaws.com",
		client: http.DefaultClient,
	}
}

func (c *AWSCheckIPClient) Get(ctx context.Context) (net.IP, error) {
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str := strings.TrimSpace(string(bytes))
	parsed := net.ParseIP(str)
	if parsed == nil {
		return nil, errors.New(fmt.Sprintf("Failed to parse: %s", str))
	}

	return parsed, nil
}
