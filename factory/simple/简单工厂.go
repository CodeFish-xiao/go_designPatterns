package simple

import "fmt"

//简单工厂模式
//宠物
type Pet interface {
	call()
}

//狗
type Dog struct {
}

func (*Dog) call() {
	fmt.Println("汪汪汪——————")
}

//猫
type Cat struct {
}

func (*Cat) call() {
	fmt.Println("喵喵喵————")
}

type PetFactory struct {
	//宠物工厂
}

func (*PetFactory) CreatePet(like string) Pet {
	if like == "cat" {
		return &Cat{}
	} else if like == "dog" {
		return &Dog{}
	}
	return nil
}
