package operate

import (
	"net/http"
)

// func Login(cond *AuthCond) (*Auth, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.AuthPath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &Auth{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func LoginValidate(cond *AuthValidateCond) (*AuthValidate, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.AuthValidatePath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &AuthValidate{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func WalletTransfer(cond *ExchangeCond) (*Exchange, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.ExchangePath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &Exchange{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func Logout(cond *CheckoutCond) (*Checkout, error) {
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: ulg168utils.CheckoutPath,
// 		Cond: cond,
// 	}
// 	res, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, h.Cond.Body())
// 	if err != nil {
// 		return nil, err
// 	}
// 	obj := &Checkout{}
// 	if err := json.Unmarshal(res, obj); err != nil {
// 		return nil, err
// 	}
// 	return obj, nil
// }

// func Settlement(gameID string) error {
// 	tmp := strings.Replace(ulg168utils.SettlementPath, ":game_id", gameID, -1)
// 	h := &ulg168utils.Acion{
// 		Host: ulg168utils.Conf.APIHost,
// 		Path: tmp,
// 		Cond: nil,
// 	}
// 	_, err := httprequest.PostForm(h.URL(), ulg168utils.DefaultTimeout, nil)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func LoginReverse(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.AuthPath
// 		fmt.Println(req.URL)
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func LoginValidateReverse(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.AuthValidatePath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func ExchangeReverse(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.ExchangePath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func LogoutReverse(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.CheckoutPath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func GetMemberInfo(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.MemberInfoPath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

// func ULG168Login(c *gin.Context) {
// 	director := func(req *http.Request) {
// 		urls := strings.Split(ulg168utils.Conf.APIHost, "://")
// 		schema, host := urls[0], urls[1]
// 		req.Host = host
// 		req.Method = "POST"
// 		req.URL.Host = req.Host
// 		req.URL.Scheme = schema
// 		req.URL.Path = ulg168utils.ULG168LoginPath
// 	}
// 	proxy := &httputil.ReverseProxy{Director: director}
// 	proxy.ModifyResponse = addCustomHeader
// 	proxy.ServeHTTP(c.Writer, c.Request)
// }

func addCustomHeader(r *http.Response) error {
	r.Header["Access-Control-Allow-Origin"] = []string{}
	return nil
}
