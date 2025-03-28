package meituanMediaSdk

import (
	"github.com/sonhineboy/meituanMediaSdk/apis"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestMtClient(t *testing.T) {

	var (
		appKey    = ""
		appSecret = ""
	)

	client := NewClient()
	requestBody := apis.NewRequestBody(apis.WithPlatform(1), apis.WithListTopiId(1))
	queryCoupon := apis.NewQueryCoupon(requestBody)

	bodyJson := queryCoupon.BuildBody()

	res, err := client.exec(queryCoupon, apis.NewHeaders(appKey, appSecret, bodyJson))
	if err != nil {
		t.Error(err)
		return
	}

	all, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(all))

	assert.Equal(t, 200, res.StatusCode)
}
