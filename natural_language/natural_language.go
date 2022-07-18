package natural_language

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fzkun/goutil/jsonutil"
	"github.com/fzkun/xfyun/natural_language/config"
	"github.com/fzkun/xfyun/natural_language/context"
	"github.com/fzkun/xfyun/xfmodel"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"net/url"
	"strconv"
	"time"
)

// NaturalLanguage 自然语言处理
type NaturalLanguage struct {
	ctx *context.Context
}

func NewNaturalLanguage(cfg *config.Config) *NaturalLanguage {
	return &NaturalLanguage{
		ctx: &context.Context{
			Config: cfg,
		},
	}
}

// Emotion 情感分析 API
//https://www.xfyun.cn/doc/nlp/emotion-analysis/API.html
func (n *NaturalLanguage) Emotion(text string) (data xfmodel.EmotionResp, err error) {
	var (
		httpResp *resty.Response
		respJson string
	)
	param := make(map[string]string)
	param["type"] = "dependent"
	tmp, _ := json.Marshal(param)
	base64Param := base64.StdEncoding.EncodeToString(tmp)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)

	//X-CheckSum
	checksum := fmt.Sprintf("%x", md5.Sum([]byte(n.ctx.ApiKey+curTime+base64Param)))

	req := resty.New().R().EnableTrace()
	req.SetHeaders(map[string]string{
		"X-Appid":    n.ctx.AppID,
		"X-CurTime":  curTime,
		"X-Param":    base64Param,
		"X-CheckSum": checksum,
	})
	values := url.Values{}
	values.Add("text", text)
	req.SetFormDataFromValues(values)
	if httpResp, err = req.Post("http://ltpapi.xfyun.cn/v2/sa"); err != nil {
		return
	}
	respJson = httpResp.String()
	respCode := gjson.Get(respJson, "code").Int()
	if respCode != 0 {
		err = errors.New(fmt.Sprintf("xf_code=%d,xf_err=%s", respCode, gjson.Get(respJson, "desc").String()))
		return
	}
	if err = jsonutil.JsonStrToStruct(gjson.Get(respJson, "data").String(), &data); err != nil {
		return
	}
	return
}
