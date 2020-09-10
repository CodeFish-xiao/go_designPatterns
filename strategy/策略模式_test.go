package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{}//创建执行者

	operator.setStrategy(&add{})//设置策略
	result := operator.calculate(1, 2)
	fmt.Println("add:", result)

	operator.setStrategy(&reduce{})
	result = operator.calculate(2, 1)
	fmt.Println("reduce:", result)
}
