package apis

import "encoding/json"

const (
	Path   = "/cps_open/common/api/v1/query_coupon"
	Method = "POST"
)

type QueryCoupon struct {
	Body   *RequestBody
	Path   string
	Method string
}

func NewQueryCoupon(q *RequestBody) *QueryCoupon {
	return &QueryCoupon{
		Body:   q,
		Path:   Path,
		Method: Method,
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

type RequestBody struct {
	Platform        int      `json:"platform,omitempty"`        // 1-到家/其他，2-到店（默认1）
	BizLine         *int     `json:"bizLine,omitempty"`         // 使用指针表示可能为null
	Longitude       int64    `json:"longitude,omitempty"`       // 经度*1000000
	Latitude        int64    `json:"latitude,omitempty"`        // 纬度*1000000
	PriceCap        *int     `json:"priceCap,omitempty"`        // 价格上限（单位元）
	PriceFloor      *int     `json:"priceFloor,omitempty"`      // 价格下限（单位元）
	CommissionCap   *int     `json:"commissionCap,omitempty"`   // 佣金上限（单位元）
	CommissionFloor []int    `json:"commissionFloor,omitempty"` // 佣金下限列表（单位元）
	VpSkuViewIds    []string `json:"vpSkuViewIds,omitempty"`    // 商品ID集合（最多20个）
	ListTopiId      *int     `json:"listTopiId,omitempty"`      // 选品池榜单ID
	SearchText      *string  `json:"searchText,omitempty"`      // 搜索关键字（1-100字符）
	SearchId        string   `json:"searchId,omitempty"`        // 搜索分页ID
	PageSize        *int     `json:"pageSize,omitempty"`        // 分页大小（默认20）
	PageNo          *int     `json:"pageNo,omitempty"`          // 页码（默认1）
	SortField       *int     `json:"sortField,omitempty"`       // 排序字段
	AscDescOrder    *int     `json:"ascDescOrder,omitempty"`    // 排序顺序（1升序，2降序）
}

func NewRequestBody(options ...Option) *RequestBody {
	re := new(RequestBody)
	for i, _ := range options {
		options[i](re)
	}
	return re
}

type Option func(*RequestBody)

// WithPlatform 设置 Platform (1-到家/其他，2-到店)
func WithPlatform(p int) Option {
	return func(r *RequestBody) {
		r.Platform = p
	}
}

// WithBizLine 设置 BizLine (指针类型，允许 nil)
func WithBizLine(b int) Option {
	return func(r *RequestBody) {
		r.BizLine = &b
	}
}

// WithLongitude 设置经度（自动处理 *1e6 逻辑）
func WithLongitude(l float64) Option {
	return func(r *RequestBody) {
		r.Longitude = int64(l * 1e6)
	}
}

// WithLatitude 设置纬度（自动处理 *1e6 逻辑）
func WithLatitude(l float64) Option {
	return func(r *RequestBody) {
		r.Latitude = int64(l * 1e6)
	}
}

// WithPriceCap 设置价格上限（指针类型）
func WithPriceCap(pc int) Option {
	return func(r *RequestBody) {
		r.PriceCap = &pc
	}
}

// WithPriceFloor 设置价格下限（指针类型）
func WithPriceFloor(pf int) Option {
	return func(r *RequestBody) {
		r.PriceFloor = &pf
	}
}

// WithCommissionCap 设置佣金上限（指针类型）
func WithCommissionCap(cc int) Option {
	return func(r *RequestBody) {
		r.CommissionCap = &cc
	}
}

// WithCommissionFloor 设置佣金下限列表（支持可变参数）
func WithCommissionFloor(cf ...int) Option {
	return func(r *RequestBody) {
		r.CommissionFloor = cf
	}
}

// WithVpSkuViewIds 设置商品ID集合（最多20个）
func WithVpSkuViewIds(ids ...string) Option {
	return func(r *RequestBody) {
		if len(ids) > 20 {
			ids = ids[:20] // 自动截断
		}
		r.VpSkuViewIds = ids
	}
}

// WithListTopiId 设置选品池榜单ID（指针类型）
func WithListTopiId(lti int) Option {
	return func(r *RequestBody) {
		r.ListTopiId = &lti
	}
}

// WithSearchText 设置搜索关键字（自动校验长度）
func WithSearchText(st string) Option {
	return func(r *RequestBody) {
		if len(st) >= 1 && len(st) <= 100 {
			r.SearchText = &st
		} else {
			panic("searchText长度需在1-100字符之间")
		}
	}
}

// WithSearchId 设置搜索分页ID
func WithSearchId(sid string) Option {
	return func(r *RequestBody) {
		r.SearchId = sid
	}
}

// WithPageSize 设置分页大小（带默认值覆盖）
func WithPageSize(ps int) Option {
	return func(r *RequestBody) {
		r.PageSize = &ps
	}
}

// WithPageNo 设置页码（带默认值覆盖）
func WithPageNo(pn int) Option {
	return func(r *RequestBody) {
		r.PageNo = &pn
	}
}

// WithSortField 设置排序字段
func WithSortField(sf int) Option {
	return func(r *RequestBody) {
		r.SortField = &sf
	}
}

// WithAscDescOrder 设置排序顺序（1升序，2降序）
func WithAscDescOrder(ado int) Option {
	return func(r *RequestBody) {
		r.AscDescOrder = &ado
	}
}
