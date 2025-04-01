package apis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReferralLink(t *testing.T) {
	link := NewReferralLink(NewRequestBody(WithPlatform(3)))
	assert.Equal(t, "POST", link.Method)
}
