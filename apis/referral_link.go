package apis

import "encoding/json"

type ReferralLink struct {
	Body   *RequestBody
	Path   string
	Method string
}

func NewReferralLink(q *RequestBody) *ReferralLink {
	return &ReferralLink{
		Body:   q,
		Path:   "/cps_open/common/api/v1/get_referral_link",
		Method: "POST",
	}
}

func (q *ReferralLink) GetMethod() string {
	return q.Method
}

func (q *ReferralLink) GetPath() string {
	return q.Path
}

func (q *ReferralLink) BuildBody() []byte {
	if body, err := json.Marshal(q.Body); err != nil {
		return nil
	} else {
		return body
	}
}
