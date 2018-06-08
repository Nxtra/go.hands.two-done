package main

// insult represent an insult
type insult struct {
	ID    int    `json:"id"`
	Text  string `json:"text"`
	Score int    `json:"score"`
}
