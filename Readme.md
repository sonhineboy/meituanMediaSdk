### 美团联盟Sdk
由于官方没有提供golang版的所以自己不得不造轮子，本着“大家好才是真的好”精神，所以分享出来
#### 功能

- [x] headers 签名
- [x] 商品查询接口
- [ ] 订单回调用验签
- [ ] 获取推广链接接口
- [ ] 查询订单接口

> 其实只要完成一个接口其它的都一样，欢迎pr
### 使用方法

**安装**
```go
go get -u github.com/sonhineboy/meituanMediaSdk
```
**使用**
```go
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
		fmt.Error(err)
		return
	}

	all, err := io.ReadAll(res.Body)
	if err != nil {
        fmt.Error(err)
		return
	}
	fmt.Println(string(all))
```