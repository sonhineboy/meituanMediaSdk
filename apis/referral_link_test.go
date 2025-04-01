package apis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReferralLink(t *testing.T) {
	link := NewReferralLink(NewReferralLinkRequestBody(func(body *ReferralLinkRequestBody) {
		body.Platform = 1
	}))
	assert.Equal(t, "POST", link.Method)
}
