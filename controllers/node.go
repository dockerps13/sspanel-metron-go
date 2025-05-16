package controllers

import (
	"net/http"
	"strconv"
	"sspanel-metron-go/models"
	"sspanel-metron-go/services"

	"github.com/gin-gonic/gin"
)

func ListNodes(c *gin.Context) {
	nodes, err := services.GetAllNodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取节点失败"})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

func CreateNode(c *gin.Context) {
	var node models.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if err := services.AddNode(node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "新增节点失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "新增节点成功"})
}

func UpdateNode(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var node models.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}
	if err := services.UpdateNode(id, node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新节点失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新节点成功"})
}

func DeleteNode(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteNode(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除节点失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除节点成功"})
}
