package observer

import "testing"

func TestObserver(t *testing.T) {

	customerA := &CustomerA{}
	customerB := &CustomerB{}

	office := &NewsOffice{}
	// 客户订阅报纸
	office.addCustomer(customerA)
	office.addCustomer(customerB)
	// 新的报纸
	office.newspaperCome()

}
