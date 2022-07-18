package xfmodel

type EmotionResp struct {
	Score     float64 `json:"score"`
	Sentiment int     `json:"sentiment"`
}
