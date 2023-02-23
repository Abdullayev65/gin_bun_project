package hendler

import (
	"github.com/Abdullayev65/gin_bun_project/pkg/io"
	"github.com/Abdullayev65/gin_bun_project/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) UserAdd(c *gin.Context) {
	var sign io.Sign
	c.Bind(&sign)
	user := models.User{Username: sign.Username, Password: sign.Password}
	err := h.Service.UserAdd(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, &user)
}
func (h *Handler) UserGet(c *gin.Context) {
	userID := h.getUserID(c)
	user, _ := h.Service.UserGet(userID)
	c.JSON(http.StatusOK, user)
}
func (h *Handler) UserAll(c *gin.Context) {
	users := h.Service.UserAll()
	c.JSON(200, &users)
}
func (h *Handler) UserToken(c *gin.Context) {
	var sign io.Sign
	c.Bind(&sign)
	user, err := h.Service.UserByUsername(sign.Username)
	if err != nil || user.Password != sign.Password {
		c.String(http.StatusBadRequest, "username or password wrong")
		return
	}
	token, err := h.TokenJWT.Generate(user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
func (h *Handler) UserUpdate(c *gin.Context) {
	var sign io.Sign
	c.Bind(&sign)

	userID := h.getUserID(c)
	user := models.User{ID: userID, Username: sign.Username, Password: sign.Password}
	err := h.Service.UserUpdate(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, &user)
}
func (h *Handler) UserDelete(c *gin.Context) {
	userID := h.getUserID(c)
	err := h.Service.UserDelete(userID)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, "done")
}

func (h *Handler) ParseTokenAndRequiredSetUserID(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		c.String(http.StatusUnauthorized, "Authorization header is empty")
		c.Abort()
		return
	}
	fields := strings.Fields(header)
	if len(fields) != 2 {
		c.String(http.StatusUnauthorized, "Authorization header is invalid")
		c.Abort()
		return
	}
	userID, err := h.TokenJWT.Parse(fields[1])
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		c.Abort()
		return
	}
	c.Set("userID", userID)
}
func (h *Handler) getUserID(c *gin.Context) int64 {
	userID, _ := c.Get("userID")
	return userID.(int64)
}
