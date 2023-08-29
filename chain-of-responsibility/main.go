package main

import "fmt"

type OrderHandler func(*Order) error

type Order struct {
	// 订单信息
	// ...
}

type OrderProcessingChain struct {
	handlers []OrderHandler
}

func NewOrderProcessingChain() *OrderProcessingChain {
	return &OrderProcessingChain{}
}

func (c *OrderProcessingChain) AddHandler(handler OrderHandler) {
	c.handlers = append(c.handlers, handler)
}

func (c *OrderProcessingChain) ProcessOrder(order *Order) error {
	for _, handler := range c.handlers {
		if err := handler(order); err != nil {
			return err
		}
	}
	return nil
}

func VerifyIdentity(order *Order) error {
	// 身份验证逻辑
	// ...
	return nil
}

func CheckInventory(order *Order) error {
	// 库存检查逻辑
	// ...
	return nil
}

func VerifyPayment(order *Order) error {
	// 支付验证逻辑
	// ...
	return nil
}

func main() {
	order := &Order{ /* 订单信息 */ }

	chain := NewOrderProcessingChain()
	chain.AddHandler(VerifyIdentity)
	chain.AddHandler(CheckInventory)
	chain.AddHandler(VerifyPayment)

	if err := chain.ProcessOrder(order); err != nil {
		fmt.Println("订单处理失败:", err)
		return
	}

	fmt.Println("订单处理成功")
}
