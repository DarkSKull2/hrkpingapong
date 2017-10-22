package models

import "encoding/json"

//Activity stuff
type Activity struct {
	Action           string                `json:"action"`
	Attachments      Attachments           `json:"attachments"`
	AttachmentLayout string                `json:"attachmentLayout"`
	ChannelData      json.RawMessage       `json:"channelData"` // map[string]interface{}
	ChannelID        string                `json:"channelId"`
	Conversation     ConversationAccount   `json:"conversation"`
	Entities         json.RawMessage       `json:"entities"` // []map[string]interface{}
	From             ChannelAccount        `json:"from"`
	HistoryDisclosed bool                  `json:"historyDisclosed"`
	ID               string                `json:"id"`
	InputHint        string                `json:"inputHint"`
	Locale           string                `json:"locale"`
	LocalTimestamp   string                `json:"localTimestamp"` //iso8601.Time
	MembersAdded     ChannelAccounts       `json:"membersAdded"`
	MembersRemoved   ChannelAccounts       `json:"membersRemoved"`
	Recipient        ChannelAccount        `json:"recipient"`
	RelatesTo        ConversationReference `json:"relatesTo"`
	ReplyToID        string                `json:"replyToId"`
	ServiceURL       string                `json:"serviceUrl"`
	Speak            string                `json:"speak"`
	SuggestedActions SuggestedActions      `json:"suggestedActions"`
	Summary          string                `json:"summary"`
	Text             string                `json:"text"`
	TextFormat       string                `json:"textFormat"`
	Timestamp        string                `json:"timestamp"` //iso8601.Time
	TopicName        string                `json:"topicName"`
	Type             string                `json:"type"`
}

// Activities things
type Activities []Activity
