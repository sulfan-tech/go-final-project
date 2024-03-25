package photo

import (
	"context"
	"go-final-project/internal/domain/photo/model"
	"go-final-project/internal/domain/photo/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler struct {
	photoService service.PhotoServiceImpl
}

func NewPhotoHandler(photoService service.PhotoServiceImpl) *PhotoHandler {
	return &PhotoHandler{
		photoService: photoService,
	}
}

func (ph *PhotoHandler) CreatePhoto(c *gin.Context) {
	var photo model.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	createdPhoto, err := ph.photoService.CreatePhoto(context.Background(), photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdPhoto)
}

func (ph *PhotoHandler) GetPhoto(c *gin.Context) {
	photos, err := ph.photoService.GetPhoto()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photos)
}

func (ph *PhotoHandler) GetPhotoByID(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	photo, err := ph.photoService.GetPhotoByID(context.Background(), idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, photo)
}

func (ph *PhotoHandler) UpdatePhoto(c *gin.Context) {
	var photo model.Photo
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	updatedPhoto, err := ph.photoService.UpdatePhoto(context.Background(), &photo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedPhoto)
}

func (ph *PhotoHandler) DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	err := ph.photoService.DeletePhoto(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
