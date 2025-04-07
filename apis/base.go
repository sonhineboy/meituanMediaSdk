package apis

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ApiBase interface {
	BuildBody() []byte
	GetMethod() string
	GetPath() string
}

type Headers struct {
	App                  string // S-Ca-App (必填)
	Signature            string // S-Ca-Signature (必填)
	Timestamp            int64  // S-Ca-Timestamp (必填)
	ContentMD5           []byte // Content-MD5 (必填)
	enCodeContentMD5     string
	SignatureHeadersList []string // S-Ca-Signature-Headers (必填)
	Secret               string
}

func NewHeaders(app, secret string, ContentMD5 []byte) *Headers {
	return &Headers{
		App:                  app,
		Timestamp:            time.Now().UnixMilli(),
		ContentMD5:           ContentMD5,
		SignatureHeadersList: []string{"S-Ca-Timestamp", "S-Ca-App"},
		Secret:               secret,
	}
}

func (h *Headers) setEncodeContentMD5() {
	md5Content := md5.Sum(h.ContentMD5)
	h.enCodeContentMD5 = base64.StdEncoding.EncodeToString(md5Content[:])
}

func (h *Headers) SetEncodeSign(method, path string) {
	if h.enCodeContentMD5 == "" {
		h.setEncodeContentMD5()
	}
	rawData := fmt.Sprintf("%s\n%s\n%s\n%s", method, h.enCodeContentMD5, h.getSignHeaders(), path)
	myHash := hmac.New(sha256.New, []byte(h.Secret))
	myHash.Write([]byte(rawData))
	h.Signature = base64.StdEncoding.EncodeToString(myHash.Sum(nil))
}

func (h *Headers) getSignHeaders() string {
	return fmt.Sprintf("S-Ca-App:%s\nS-Ca-Timestamp:%d", h.App, h.Timestamp)
}

func (h *Headers) GetHttpHeader() http.Header {
	httpHeaders := make(http.Header)
	httpHeaders.Set("S-Ca-App", h.App)
	httpHeaders.Set("S-Ca-AppSecret", h.Secret)
	httpHeaders.Set("S-Ca-Timestamp", strconv.FormatInt(h.Timestamp, 10))
	httpHeaders.Set("S-Ca-Signature", h.Signature)
	httpHeaders.Set("Content-MD5", h.enCodeContentMD5)
	httpHeaders.Set("S-Ca-Signature-Headers", strings.Join(h.SignatureHeadersList, ","))
	httpHeaders.Add("Content-Type", "application/json")
	return httpHeaders
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
	ActID           string   `json:"actId,omitempty"`           // 活动物料ID
	SkuViewID       string   `json:"skuViewId,omitempty"`       // 商品ID
	SID             string   `json:"sid,omitempty"`             // 二级媒体身份标识
	LinkType        *int     `json:"linkType,omitempty"`        // 链接类型（指针类型） 链接类型，枚举值：1 H5长链接；2 H5短链接；3 deeplink(唤起)链接；4 微信小程序唤起路径；5 团口令；6 小程序码
	LinkTypeList    []int    `json:"linkTypeList,omitempty"`    // 链接类型列表 (linkType和linkTypeList必传一个，linkType和linkTypeList都传时，只处理linkTypeList)。枚举值：1 H5长链接；2 H5短链接；3 deeplink(唤起)链接；4 微信小程序唤起路径；5 团口令；6 小程序码
	Text            string   `json:"text,omitempty"`            // 活动链接文本
	CityId          string   `json:"cityId,omitempty"`          //城市编码，榜单场景、多业务供给场景、搜索场景生效。城市ID下载：https://s3plus.meituan.net/media-public/%E5%9F%8E%E5%B8%82%E5%AD%97%E5%85%B82025.xlsx
	BusinessAreaId  *int     `json:"businessAreaId,omitempty"`  //商圈编码，榜单场景、多业务供给场景、搜索场景生效。商圈下载：https://s3plus.meituan.net/media-public/%E5%95%86%E5%9C%88%E5%9F%8E%E5%B8%82%E6%98%A0%E5%B0%84%E5%AD%97%E5%85%B82025.xlsx
}

func NewRequestBody(options ...Option) *RequestBody {
	re := new(RequestBody)
	for i, _ := range options {
		options[i](re)
	}
	return re
}

type Option func(*RequestBody)

func WithBusinessAreaId(businessAreaId int) Option {
	return func(r *RequestBody) {
		r.BusinessAreaId = &businessAreaId
	}
}

func WithCityId(cityId string) Option {
	return func(r *RequestBody) {
		r.CityId = cityId
	}
}

func WithActID(key string) Option {
	return func(r *RequestBody) {
		r.ActID = key
	}
}

func WithSkuViewID(skuViewID string) Option {
	return func(r *RequestBody) {
		r.SkuViewID = skuViewID
	}
}

func WithSID(sid string) Option {
	return func(r *RequestBody) {
		r.SID = sid
	}
}

func WithLinkType(linkType *int) Option {
	return func(r *RequestBody) {
		r.LinkType = linkType
	}
}
func WithLinkTypeList(linkTypeList []int) Option {
	return func(r *RequestBody) {
		r.LinkTypeList = linkTypeList
	}
}

func WithText(txt string) Option {
	return func(r *RequestBody) {
		r.Text = txt
	}
}

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
