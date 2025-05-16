package models

import "gorm.io/gorm"

type Node struct {
	gorm.Model
	Name       string
	Address    string
	Port       int
	Status     string // 节点状态，在线/离线
	Protocol   string // 支持协议类型，例如 vmess, trojan, ss
}

func CreateNode(node *Node) error {
	return db.Create(node).Error
}

func GetAllNodes() ([]Node, error) {
	var nodes []Node
	err := db.Find(&nodes).Error
	return nodes, err
}

func UpdateNodeByID(id int, node Node) error {
	return db.Model(&Node{}).Where("id = ?", id).Updates(node).Error
}

func DeleteNodeByID(id int) error {
	return db.Delete(&Node{}, id).Error
}
