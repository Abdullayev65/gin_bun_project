package hendler

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/io"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) PostAdd(c *gin.Context) {
	userID := h.getUserID(c)
	var postInput io.PostInput
	c.Bind(&postInput)
	post := &models.Post{UserID: userID, Description: postInput.Description}
	err := h.Service.PostAdd(post)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, post)
}
func (h *Handler) PostAll(c *gin.Context) {
	posts := h.Service.PostAll()
	c.JSON(200, &posts)
}
func (h *Handler) PostGet(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id of post is invalid")
		return
	}
	post, err := h.Service.PostGet(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, post)
}
func (h *Handler) PostUpdate(c *gin.Context) {
	var postInput io.PostInput
	c.Bind(&postInput)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id of post is invalid")
		return
	}
	post := &models.Post{ID: id, Description: postInput.Description}
	err = h.Service.PostUpdate(post)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, post)
}
func (h *Handler) PostDelete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id of post is invalid")
		return
	}
	err = h.Service.PostDelete(id)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, "done")
}
