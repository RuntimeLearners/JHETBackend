package models

import "time"

type SearchParams struct {
	Page          int        `json:"page"`
	Size          int        `json:"size"`
	CreaterID     *uint64    `json:"id,omitempty"`
	Keywords      *string    `json:"keywords,omitempty"`
	Status        *string    `json:"status,omitempty"`
	Urgency       *string    `json:"urgency,omitempty"` // 仅用于对应前端接口
	ShowSpams     bool       `json:"show_spams,omitempty"`
	ShowPrivates  bool       `json:"show_privates,omitempty"`
	CreatedBefore *time.Time `json:"created_before,omitempty"`
	CreatedAfter  *time.Time `json:"created_after,omitempty"`
	UpdatedBefore *time.Time `json:"updated_before,omitempty"`
	UpdatedAfter  *time.Time `json:"updated_after,omitempty"`
	MinPrecedence *uint8     // 通过前端给出的Urgency在后端解析
	MaxPrecedence *uint8     // 通过前端给出的Urgency在后端解析
}
