package services

import (
	"errors"
	"sspanel-metron-go/models"
	"time"
)

// 伪代码示例，实际根据支付网关实现回调验证

func ProcessPaymentCallback(orderID string, status string) error {
	order, err := models.GetPaymentOrderByOrderID(orderID)
	if err != nil {
		return errors.New("订单不存在")
	}

	if order.Status == "paid" {
		return nil // 已支付，避免重复处理
	}

	if status != "success" {
		return errors.New("支付未成功")
	}

	// 标记订单为已支付
	order.Status = "paid"
	order.PaidAt = time.Now()
	if err := models.UpdatePaymentOrder(order); err != nil {
		return err
	}

	// 给用户添加流量或者续期
	user, err := models.GetUserByID(order.UserID)
	if err != nil {
		return err
	}

	// 假设订单包含购买流量
	user.TransferEnable += order.TransferAmount
	if err := models.UpdateUser(user); err != nil {
		return err
	}

	return nil
}
