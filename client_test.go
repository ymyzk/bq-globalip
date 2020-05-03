package main

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAWSCheckIPClientGetWorks(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`8.8.8.8`))
	}))
	defer server.Close()

	client := NewAWSCheckIPClient()
	client.URL = server.URL

	addr, err := client.Get(ctx)
	assert.Equal(t, addr, net.ParseIP("8.8.8.8"))
	assert.Nil(t, err)
}

func TestAWSCheckIPClientGetWrongFormat(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"address": 8.8.8.8}`))
	}))
	defer server.Close()

	client := NewAWSCheckIPClient()
	client.URL = server.URL

	addr, err := client.Get(ctx)
	assert.Nil(t, addr)
	assert.NotNil(t, err)
}

func TestAWSCheckIPClientGetHTTPError(t *testing.T) {
	ctx := context.Background()

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		http.Error(rw, "", http.StatusBadRequest)
	}))
	defer server.Close()

	client := NewAWSCheckIPClient()
	client.URL = server.URL

	addr, err := client.Get(ctx)
	assert.Nil(t, addr)
	assert.NotNil(t, err)
}
