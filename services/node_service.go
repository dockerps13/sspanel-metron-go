package services

import "sspanel-metron-go/models"

func GetAllNodes() ([]models.Node, error) {
	return models.GetAllNodes()
}

