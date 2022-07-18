package context

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/fzkun/xfyun/config"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"strconv"
	"time"
)

// Context struct
type Context struct {
	*config.Config
}

func (ctx *Context) ApiRequest(url string, formData map[string]string, bodyData map[string]string, obj any) (httpResp *resty.Response, err error) {
	var (
		respJson string
	)
	base64Param := base64.StdEncoding.EncodeToString([]byte(jsonutil.StructToJsonString(bodyData)))
	curTime := strconv.FormatInt(time.Now().Unix(), 10)

	//X-CheckSum
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(ctx.ApiKey+curTime+base64Param)))

	req := resty.New().R().EnableTrace()
	req.SetHeaders(map[string]string{
		"X-Appid":    ctx.AppID,
		"X-CurTime":  curTime,
		"X-Param":    base64Param,
		"X-CheckSum": checksum,
	})
	req.SetFormData(formData)
	if httpResp, err = req.Post(url); err != nil {
		return
	}
	respJson = httpResp.String()
	respCode := gjson.Get(respJson, "code").Int()
	if respCode != 0 {
		err = errors.New(fmt.Sprintf("xf_code=%d,xf_err=%s", respCode, gjson.Get(respJson, "desc").String()))
		return
	}
	if obj != nil {
		if err = jsonutil.JsonStrToStruct(gjson.Get(respJson, "data").String(), &obj); err != nil {
			return
		}
	}
	return
}
