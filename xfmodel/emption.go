package xfmodel

type EmotionResp struct {
	Score     float64 `json:"score"`     //分类对应得分，范围 0-1
	Sentiment int     `json:"sentiment"` //情感极性分类结果 0：中性 1：褒义 -1：贬义
}

func (e EmotionResp) SentimentString() string {
	switch e.Sentiment {
	case 0:
		return "中性"
	case 1:
		return "褒义"
	case -1:
		return "贬义"
	default:
		return "不详"
	}

}
