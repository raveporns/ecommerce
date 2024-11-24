package modal

import "database/sql"

// Comment คือตัวแทนของคอมเมนต์ในระบบ
type Comment struct {
	CommentID   int    `json:"comment_id"`
	CommentText string `json:"comment_text"`
	Rating      int    `json:"rating"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// ScanComment ใช้สำหรับแปลงข้อมูลจากฐานข้อมูลเป็น Comment struct
func ScanComment(rows *sql.Rows) (*Comment, error) {
	var comment Comment

	err := rows.Scan(
		&comment.CommentID,
		&comment.CommentText,
		&comment.Rating,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
