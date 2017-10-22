package models

//SuggestedActions desc
type SuggestedActions struct {
	To      []string    `json:"to"`
	Actions CardActions `json:"actions"`
}
