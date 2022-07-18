package xfmodel

type KeysResp struct {
	Ke []KeysInfo `json:"ke"`
}
type KeysInfo struct {
	Score string `json:"score"` //关键词
	Word  string `json:"word"`  //候选词成为关键词的概率 (float)
}
