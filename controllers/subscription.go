package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"sspanel-metron-go/models"
	"sspanel-metron-go/services"
)

func GetSubscription(c *gin.Context) {
	userID := c.GetUint("userID")

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户不存在"})
		return
	}

	if user.ExpiredAt.Before(time.Now()) {
		c.JSON(http.StatusForbidden, gin.H{"error": "账户已过期"})
		return
	}

	subscriptionLink, err := services.GenerateSubscriptionLink(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成订阅链接失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscription": subscriptionLink})
}
