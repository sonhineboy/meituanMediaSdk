package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestJson(t *testing.T) {

	type diyJson struct {
		A int    `json:"a,omitempty"`
		B *int   `json:"b,omitempty"`
		C int64  `json:"c,omitempty"`
		D *int64 `json:"d,omitempty"`
	}

	var b int = 0
	var d int64 = 0
	myJson, err := json.Marshal(diyJson{
		A: 0,
		B: &b,
		C: 0,
		D: &d,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(string(myJson))
}

func TestUrl(t *testing.T) {
	parse, err := url.Parse("http://localhost:8080/asdf/asdfasd")
	if err != nil {
		return
	}
	assert.Equal(t, "/asdf/asdfasd", parse.Path)
}
