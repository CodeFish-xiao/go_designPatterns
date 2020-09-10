package observer

import "fmt"

//又称发布订阅模式
// 客户抽象类（观察者）
type Customer interface {
	update()
}

type CustomerA struct {
}

func (*CustomerA) update() {
	fmt.Println("我是客户A, 我收到报纸了")
}

type CustomerB struct {
}

func (*CustomerB) update() {
	fmt.Println("我是客户B, 我收到报纸了")
}

// 报社 （被观察者)
type NewsOffice struct {
	customers []Customer
}

func (n *NewsOffice) addCustomer(customer Customer) { //客户注册
	n.customers = append(n.customers, customer)
}

func (n *NewsOffice) newspaperCome() {
	// 通知所有客户
	n.notifyAllCustomer()
}

func (n *NewsOffice) notifyAllCustomer() {
	for _, customer := range n.customers {
		customer.update()
	}
}
