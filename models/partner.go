package models

import "time"

// Partner 人魚漢頓的夥伴（三麗鷗角色，如小百合等）
// 參考：人魚漢頓（ハンギョドン）是三麗鷗 1985 年推出的角色，最好的朋友是章魚女孩小百合
type Partner struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Intro     string    `json:"intro"` // 夥伴簡介
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
