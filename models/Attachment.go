package models

//Attachment shut up the linter
type Attachment struct {
	ContentType string `json:"contentType"`
}

//Attachments shut up the linter
type Attachments []Attachment
