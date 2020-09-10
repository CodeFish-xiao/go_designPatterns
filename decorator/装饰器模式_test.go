package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	laowang := &laowang{} //实例化老王

	jacket := &Jacket{}     //夹克装饰器实现类
	jacket.person = laowang //装饰老王
	jacket.show()

	hat := &Hat{}
	hat.person = jacket
	hat.show()

	fmt.Println("cost:", hat.cost())
}
