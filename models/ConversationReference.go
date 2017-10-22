package models

// ConversationReference def
type ConversationReference struct {
	ActivityID   string              `json:"activityId"`
	Bot          ChannelAccount      `json:"bot"`
	ChannelID    string              `json:"channelId"`
	Conversation ConversationAccount `json:"conversation"`
	ServiceURL   string              `json:"serviceUrl"`
	User         ChannelAccount      `json:"user"`
}

// ConversationReferences def
type ConversationReferences []ConversationReference
