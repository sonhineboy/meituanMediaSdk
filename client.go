package meituanMediaSdk

import (
	"bytes"
	"fmt"
	"github.com/sonhineboy/meituanMediaSdk/apis"
	"net/http"
	"time"
)

const BaseUrl = "https://media.meituan.com"

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	return &Client{
		client: httpClient,
	}
}
func (c *Client) exec(apiBase apis.ApiBase, myHeaders *apis.Headers) (*http.Response, error) {
	request, err := http.NewRequest(apiBase.GetMethod(), fmt.Sprintf("%s%s", BaseUrl, apiBase.GetPath()), bytes.NewReader(apiBase.BuildBody()))
	if err != nil {
		return nil, err
	}
	myHeaders.SetEncodeSign(apiBase.GetMethod(), apiBase.GetPath())
	request.Header = myHeaders.GetHttpHeader()
	return c.client.Do(request)
}
