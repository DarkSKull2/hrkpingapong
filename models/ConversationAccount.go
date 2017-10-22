package models

//ConversationAccount description
type ConversationAccount struct {
	ID      string `json:"id"`
	IsGroup bool   `json:"isGroup"`
	Name    string `json:"name"`
}

//ConversationAccounts details
type ConversationAccounts []ConversationAccount
