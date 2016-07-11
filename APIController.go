package apiauth

import (
	"github.com/astaxie/beego"
	"fmt"
)

type APIController struct {
	beego.Controller
	err  error
	data interface{}
}
// 函数结束时,组装成json结果返回
func (this *APIController) Finish() {
	r := struct {
		Error interface{} `json:"error"`
		Data  interface{} `json:"data"`
	}{}
	if this.err != nil {
		r.Error = this.err.Error()
	}
	r.Data = this.data
	this.Data["json"] = r
	this.ServeJSON()
}
// 如果请求的参数不存在,就直接 error返回
func (this *APIController) MustString(key string) string {
	v := this.GetString(key)
	if v == "" {
		this.Data["json"] = map[string]string{
			"error": fmt.Sprintf("require filed: %s", key),
			"data":  "orz!!",
		}
		this.ServeJSON()
		this.StopRun()
	}
	return v
}