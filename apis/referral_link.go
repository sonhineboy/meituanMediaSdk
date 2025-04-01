package apis

import "encoding/json"

type ReferralLink struct {
	Body   *ReferralLinkRequestBody
	Path   string
	Method string
}

func NewReferralLink(q *ReferralLinkRequestBody) *ReferralLink {
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

type ReferralOption func(*ReferralLinkRequestBody)
type ReferralLinkRequestBody struct {
	Platform     int    `json:"platform,omitempty"`     // 商品业务一级分类
	BizLine      *int   `json:"bizLine,omitempty"`      // 商品业务二级分类（指针类型区分零值）
	ActID        string `json:"actId,omitempty"`        // 活动物料ID
	SkuViewID    string `json:"skuViewId,omitempty"`    // 商品ID
	SID          string `json:"sid,omitempty"`          // 二级媒体身份标识
	LinkType     *int   `json:"linkType,omitempty"`     // 链接类型（指针类型） 链接类型，枚举值：1 H5长链接；2 H5短链接；3 deeplink(唤起)链接；4 微信小程序唤起路径；5 团口令；6 小程序码
	LinkTypeList []int  `json:"linkTypeList,omitempty"` // 链接类型列表 (linkType和linkTypeList必传一个，linkType和linkTypeList都传时，只处理linkTypeList)。枚举值：1 H5长链接；2 H5短链接；3 deeplink(唤起)链接；4 微信小程序唤起路径；5 团口令；6 小程序码
	Text         string `json:"text,omitempty"`         // 活动链接文本
}

// NewReferralLinkRequestBody 创建请求体并进行校验
func NewReferralLinkRequestBody(opts ...ReferralOption) *ReferralLinkRequestBody {
	req := &ReferralLinkRequestBody{}
	for _, opt := range opts {
		opt(req)
	}
	return req
}
