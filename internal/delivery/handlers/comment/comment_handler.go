package commenthandler

import (
	"context"
	"encoding/json"
	"go-final-project/internal/domain/comment/model"
	commentservice "go-final-project/internal/domain/comment/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	commentService commentservice.CommentServiceImpl
}

// NewCommentHandler creates a new instance of CommentHandler.
func NewCommentHandler(commentService commentservice.CommentServiceImpl) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// CreateComment handles the creation of a new comment.
func (ch *CommentHandler) CreateComment(c *gin.Context) {
	var commentReq CommentRequest
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	comment := model.Comment{
		Message: commentReq.Message,
		PhotoID: uint(commentReq.PhotoID),
	}

	createdComment, err := ch.commentService.CreateComment(context.Background(), int(comment.Photo.ID), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdComment)
}

func (ch *CommentHandler) GetComments(c *gin.Context) {
	comment, err := ch.commentService.GetComments(context.Background())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// UpdateComment updates an existing comment.
func (ch *CommentHandler) UpdateComment(c *gin.Context) {
	var comment model.Comment
	if err := json.NewDecoder(c.Request.Body).Decode(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedComment, err := ch.commentService.UpdateComment(context.Background(), &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedComment)
}

// DeleteComment deletes a comment by its ID.
func (ch *CommentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	err = ch.commentService.DeleteComment(context.Background(), uint(idInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
