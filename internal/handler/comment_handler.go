package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-project/internal/models"
	"rest-project/internal/services"
	"strconv"
)

type CommentHandler struct {
	Service services.CommentService
}

func NewCommentHandler(service services.CommentService) *CommentHandler {
	return &CommentHandler{Service: service}
}

func (h *CommentHandler) GetCommentsForPost(c *gin.Context) {
	postID := c.Param("post_id")
	postIDUint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	comments, err := h.Service.GetCommentsForPost(uint(postIDUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No comments found"})
		return
	}
	c.JSON(http.StatusOK, comments)
}

func (h *CommentHandler) CreateComment(c *gin.Context) {
	postID := c.Param("post_id")
	postIDUint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var req struct {
		Content string `json:"content"`
		UserID  uint   `json:"user_id"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	comment := models.Comment{
		PostID:  uint(postIDUint), // Явное преобразование
		Content: req.Content,
		UserID:  req.UserID,
	}

	if err := h.Service.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully"})
}
