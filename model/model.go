package model

// person data save same structure
type Person_ST struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name"`
	City    string  `json:"city,omitempty"`
	Phone   int     `json:"phone,omitempty"`
	Height  float32 `json:"height,omitempty"`
	Married bool    `json:"married"`
}

// person' data search via any one field or more
type SearchField struct {
	ID      int     `json:"id,omitempty"`
	Name    string  `json:"name,omitempty"`
	City    string  `json:"city,omitempty"`
	Phone   int     `json:"phone,omitempty"`
	Height  float32 `json:"height,omitempty"`
	Married bool    `json:"married,omitempty"`
}
