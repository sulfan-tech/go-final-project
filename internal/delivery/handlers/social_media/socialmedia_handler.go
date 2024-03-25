package socialmedia

import (
	"context"
	"fmt"
	"go-final-project/internal/domain/social_media/model"
	"go-final-project/internal/domain/social_media/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	socialMediaService service.SocialMediaServiceImpl
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaServiceImpl) *SocialMediaHandler {
	return &SocialMediaHandler{
		socialMediaService: socialMediaService,
	}
}

func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var socialMedia model.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err)})
		return
	}

	createdSocialMedia, err := h.socialMediaService.CreateSocialMedia(context.Background(), &socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create social media: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, createdSocialMedia)
}

func (h *SocialMediaHandler) GetSocialMedias(c *gin.Context) {
	socialMedias, err := h.socialMediaService.GetSocialMedias(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get social medias: %v", err)})
		return
	}

	c.JSON(http.StatusOK, socialMedias)
}

func (h *SocialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	var socialMedia model.SocialMedia
	if err := c.ShouldBindJSON(&socialMedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid request payload: %v", err)})
		return
	}

	updatedSocialMedia, err := h.socialMediaService.UpdateSocialMedia(context.Background(), &socialMedia)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to update social media: %v", err)})
		return
	}

	c.JSON(http.StatusOK, updatedSocialMedia)
}

func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid ID: %v", err)})
		return
	}

	err = h.socialMediaService.DeleteSocialMedia(context.Background(), uint(idUint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to delete social media: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Social media deleted successfully"})
}
