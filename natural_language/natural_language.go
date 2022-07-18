package natural_language

import (
	"github.com/fzkun/xfyun/context"
	"github.com/fzkun/xfyun/xfmodel"
)

// NaturalLanguage 自然语言处理
type NaturalLanguage struct {
	ctx *context.Context
}

func NewNaturalLanguage(ctx *context.Context) *NaturalLanguage {
	return &NaturalLanguage{
		ctx: ctx,
	}
}

// Keys 关键词提取 API
//https://www.xfyun.cn/doc/nlp/keyword-extraction/API.html#%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E
func (n *NaturalLanguage) Keys(text string) (data xfmodel.KeysResp, err error) {
	//var httpResp *resty.Response
	_, err = n.ctx.ApiRequest("https://ltpapi.xfyun.cn/v1/ke",
		map[string]string{
			"text": text,
		}, map[string]string{
			"type": "dependent",
		}, &data)
	return
}

// Emotion 情感分析 API
//https://www.xfyun.cn/doc/nlp/emotion-analysis/API.html
func (n *NaturalLanguage) Emotion(text string) (data xfmodel.EmotionResp, err error) {
	_, err = n.ctx.ApiRequest("https://ltpapi.xfyun.cn/v2/sa",
		map[string]string{
			"text": text,
		}, map[string]string{
			"type": "dependent",
		}, &data)
	return
}
