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
	return "Wx"
}

type PayFactory interface {
	CreatePay() Pay
}
type AliPayFactory struct {
}

func (*AliPayFactory) CreatePay() Pay {
	return &Alipay{}
}

type WxFactory struct {
}

func (*WxFactory) CreatePay() Pay {
	return &Wx{}
}
func main() {
	ali := &AliPayFactory{}
	fmt.Println(ali.CreatePay().PayPage(1))
	wx := &WxFactory{}
	fmt.Println(wx.CreatePay().PayPage(1))
}
