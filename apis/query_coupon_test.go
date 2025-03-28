package apis

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryCoupon(t *testing.T) {
	c := struct {
		Platform   int  `json:"platform,omitempty"`
		ListTopiId *int `json:"listTopiId,omitempty"` // 选品池榜单ID
	}{
		Platform: 1,
	}
	body := NewRequestBody(WithPlatform(1))
	coupon := NewQueryCoupon(body)

	marshal, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, marshal, coupon.BuildBody())
}
