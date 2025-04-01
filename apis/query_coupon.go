package apis

import "encoding/json"

type QueryCoupon struct {
	Body   *RequestBody
	Path   string
	Method string
}

func NewQueryCoupon(q *RequestBody) *QueryCoupon {
	return &QueryCoupon{
		Body:   q,
		Path:   "/cps_open/common/api/v1/query_coupon",
		Method: "POST",
	}
}

func (q *QueryCoupon) GetMethod() string {
	return q.Method
}

func (q *QueryCoupon) GetPath() string {
	return q.Path
}

func (q *QueryCoupon) BuildBody() []byte {
	if body, err := json.Marshal(q.Body); err != nil {
		return nil
	} else {
		return body
	}
}
