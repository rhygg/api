package structures

type SentimentalResult struct {
	Polarity float32 `json:"polarity"`
	Type     string  `json:"type"`
}
type Sentiment struct {
	Result *SentimentalResult `json:"result"`
}
