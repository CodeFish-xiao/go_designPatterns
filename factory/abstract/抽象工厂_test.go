package abstract

import (
	"testing"
)

func TestAbstract(t *testing.T) {
	girlFactoryStore := GirlFactoryStore{
		factory: ChineseGirlFactory{},
	}
	girlFactoryStore.factory.CreateGirl("thin").weight()
}
