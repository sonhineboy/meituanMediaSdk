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
