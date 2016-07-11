package apiauth


import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"fmt"
	"net/url"
	"encoding/json"
	"encoding/base64"
)

type GetParamSecurityKEY func(string) string

func APIParamsSecurity(f GetParamSecurityKEY) beego.FilterFunc {
	return func(ctx *context.Context) {
		secretkey := f(ctx.Input.Query("appid"))
		if secretkey == "" {
			ctx.ResponseWriter.WriteHeader(403)
			ctx.WriteString("not exist this appid")
			return
		}
		if ctx.Input.Query("data") == "" {
			ctx.ResponseWriter.WriteHeader(403)
			ctx.WriteString("miss query param: data")
			return
		}
		paramdecryption(secretkey, ctx.Input.Query("data"),ctx.Request.Form)
		return
	}
}

//负责处理解密
func paramdecryption(key string, data string, params url.Values) {
	//fmt.Println("111: "+data)
	uDec, _ := base64.URLEncoding.DecodeString(data)
	//fmt.Println(string(uDec))
	var obj interface{} // var obj map[string]interface{}
	json.Unmarshal([]byte(uDec), &obj)
	m := obj.(map[string]interface{})
	for k,_ := range m {
		params.Add(k,fmt.Sprintf("%v",  m[k]))
		fmt.Println(fmt.Sprintf("%v",  m[k]))
	}
	//fmt.Println("33333:"+params.Encode())
}


