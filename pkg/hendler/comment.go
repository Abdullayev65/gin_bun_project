package hendler

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/io"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) CommentAdd(c *gin.Context) {
	userID := h.getUserID(c)
	var commentInput io.CommentInput
	c.Bind(&commentInput)
	comment := &models.Comment{UserID: userID, Text: commentInput.Text, PostID: commentInput.PostID}
	err := h.Service.CommentAdd(comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, comment)
}
func (h *Handler) CommentAll(c *gin.Context) {
	comments := h.Service.CommentAll()
	c.JSON(200, &comments)
}
func (h *Handler) CommentGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "id of comment is invalid")
		return
	}
	comment, err := h.Service.CommentGet(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, comment)
}
func (h *Handler) CommentUpdate(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "id of comment is invalid")
		return
	}
	var commentInput io.CommentInput
	c.Bind(&commentInput)
	comment := &models.Comment{ID: id, Text: commentInput.Text}
	err = h.Service.CommentUpdate(comment)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, comment)
}
func (h *Handler) CommentDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "id of comment is invalid")
		return
	}
	err = h.Service.CommentDelete(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, "done")
}
