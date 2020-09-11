package proxy

import "fmt"

//猪八戒去找高翠兰结果是孙悟空变的，可以这样理解：把高翠兰的外貌抽象出来，
//高翠兰本人和孙悟空都实现了这个接口，猪八戒访问高翠兰的时候看不出来这个是孙悟空，所以说孙悟空是高翠兰代理类。

type Seller interface {
	call()
}

// 高翠兰
type Gaocuilan struct {
	charm int //魅力值
}

func (g *Gaocuilan) call() {
	fmt.Printf("魅力值：%d八戒哥哥，不要嘛~", g.charm) //作者老闷骚怪了

}

// 孙悟空
type Sunwukong struct {
	gaocuilan *Gaocuilan // 持有一个高翠兰对象##NTR警告
}

func (s *Sunwukong) call() {
	fmt.Printf("魅力值：%d八戒哥哥，不要嘛~", s.gaocuilan.charm) //作者老闷骚怪了
}
