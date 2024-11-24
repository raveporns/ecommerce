package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raveporns/ecommerce/modals"
)

// ReviewController struct holds the DB connection.
type ReviewController struct {
	DB *sql.DB
}

// getCommentByQuery is a helper function to query comments based on a SQL query and parameters.
func (rc *ReviewController) getCommentByQuery(query string, args ...interface{}) ([]modal.Comment, error) {
	rows, err := rc.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []modal.Comment

	// Loop through the rows and scan each record
	for rows.Next() {
		var comment modal.Comment
		if err := rows.Scan(&comment.CommentID, &comment.CommentText, &comment.Rating, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return comments, nil
}

// GetCommentAll returns all comments.
func (rc *ReviewController) GetCommentAll(c *gin.Context) {
	query := `SELECT comment_id, comment_text,rating, created_at, updated_at FROM comments`

	comments, err := rc.getCommentByQuery(query)
	if err != nil {
		log.Printf("Error fetching comments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// AddComment adds a new comment.
func (rc *ReviewController) AddComment(c *gin.Context) {
	var comment modal.Comment

	// Bind JSON request to comment object
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate required fields
	if comment.CommentText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CommentText is required"})
		return
	}

	// Prepare the SQL query to insert the new comment
	query := `INSERT INTO comments (comment_text, rating, created_at, updated_at) VALUES ($1, $2, $3, $4)`

	_, err := rc.DB.Exec(query, comment.CommentText,comment.Rating, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error adding comment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding comment"})
		return
	}

	// Return success response with the added comment details
	c.JSON(http.StatusOK, gin.H{
		"message":      "Comment added successfully",
		"comment_text": comment.CommentText,
	})
}

// UpdateComment updates the comment text for a given comment ID.
func (rc *ReviewController) UpdateComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	var comment modal.Comment

	// Bind JSON request to comment object
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Validate required fields
	if comment.CommentText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "CommentText is required"})
		return
	}

	// Update the comment text and updated_at timestamp
	query := `UPDATE comments SET comment_text = $1 ,rating = $2, updated_at = $3 WHERE comment_id = $4 `
	_, err := rc.DB.Exec(query, comment.CommentText, time.Now(), commentID)
	if err != nil {
		log.Printf("Error updating comment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated"})
}

// RemoveComment removes a comment by its ID.
func (rc *ReviewController) RemoveComment(c *gin.Context) {
	commentID := c.Param("comment_id")

	// Delete the comment from the database
	query := `DELETE FROM comments WHERE comment_id = $1`
	_, err := rc.DB.Exec(query, commentID)
	if err != nil {
		log.Printf("Error removing comment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error removing comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment removed"})
}
