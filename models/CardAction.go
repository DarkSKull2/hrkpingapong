package models

//CardAction descr
type CardAction struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
	Value string `json:"value"`
}

//CardActions def
type CardActions []CardAction
