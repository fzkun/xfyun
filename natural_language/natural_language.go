package natural_language

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fzkun/xfyun/natural_language/config"
	"github.com/fzkun/xfyun/natural_language/context"
	"github.com/go-resty/resty/v2"
	"io"
	"net/url"
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

func (n *NaturalLanguage) Emotion(text string) (err error) {
	var (
		httpResp *resty.Response
	)
	param := make(map[string]string)
	param["type"] = "dependent"
	tmp, _ := json.Marshal(param)
	base64Param := base64.StdEncoding.EncodeToString(tmp)
	curTime := fmt.Sprintf("%d", time.Now().Unix())
	w := md5.New()
	io.WriteString(w, n.ctx.AppID+curTime+base64Param)
	checksum := fmt.Sprintf("%x", w.Sum(nil))

	req := resty.New().R().EnableTrace()
	req.SetHeaders(map[string]string{
		"X-Appid":    n.ctx.AppID,
		"X-CurTime":  curTime,
		"X-Param":    base64Param,
		"X-CheckSum": checksum,
		//"Content-Type": "application/x-www-form-urlencoded",
	})
	data := url.Values{}
	data.Add("text", text)
	req.SetFormDataFromValues(data)
	if httpResp, err = req.Post("http://ltpapi.xfyun.cn/v2/sa"); err != nil {
		return
	}
	fmt.Println(httpResp.String())
	return
}
