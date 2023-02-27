package hendler

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/Abdullayev65/gin_bun_project/pkg/utill"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AttachmentAdd(c *gin.Context) {
	multipartForm, err := c.MultipartForm()
	var attachs []models.Attachment
	if err != nil {
		c.String(400, err.Error())
		return
	}
	attachs = utill.UploadFiles(multipartForm.File)
	err = h.Service.AttachmentAddAll(attachs)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(200, attachs)
}
func (h *Handler) AttachmentGet(c *gin.Context) {
	id := c.GetInt("id")
	attach, err := h.Service.AttachmentGet(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.File(attach.Path)
}
