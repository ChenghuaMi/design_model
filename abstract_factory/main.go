package main

import "fmt"

type Pay interface {
	PayPage(price int) string
}
type Alipay struct {
}

func (*Alipay) PayPage(price int) string {
	return "alipay"
}

type Wx struct {
}

func (*Wx) PayPage(price int) string {
	return "wx"
}

type Refund interface {
	RefundOrder(orderNo string) string
}
type AliPayRefund struct {
}

func (*AliPayRefund) RefundOrder(orderNo string) string {
	return "alipayRefund order"
}

type WxPayRefund struct {
}

func (*WxPayRefund) RefundOrder(orderNo string) string {
	return "WxPay refund order"
}

type PayFactory interface {
	CreatePay() Pay
	CreateRefund() Refund
}
type AliPayFactory struct {
}

func (*AliPayFactory) CreatePay() Pay {
	return &Alipay{}
}
func (*AliPayFactory) CreateRefund() Refund {
	return &AliPayRefund{}
}

type WxFactory struct {
}

func (*WxFactory) CreatePay() Pay {
	return &Wx{}
}
func (*WxFactory) CreateRefund() Refund {
	return &WxPayRefund{}
}

func main() {
	ali := &AliPayFactory{}
	fmt.Println(ali.CreatePay().PayPage(1))
	fmt.Println(ali.CreateRefund().RefundOrder("1111"))

	wx := &WxFactory{}
	fmt.Println(wx.CreatePay().PayPage(1))
	fmt.Println(wx.CreateRefund().RefundOrder("qqqq"))
}
