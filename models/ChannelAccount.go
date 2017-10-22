package models

//ChannelAccount type
type ChannelAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//ChannelAccounts type
type ChannelAccounts []ChannelAccount
