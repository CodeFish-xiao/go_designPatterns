package singleton

import "sync"

var (
	goInstance *Instance
	once       sync.Once
)




// 使用go 实现单例模式
func GoInstance(name string) *Instance {
	if goInstance == nil {
		once.Do(func() {
			goInstance = &Instance{
				Name: name,
			}
		})
	}
	return goInstance
}

/*
其实就是使用 once.Do 来保证 某个对象只会初始化一次，
有一点要要注意的是 这个 once.Do 只会被运行一次，
哪怕 Do 函数里面的发生了异常，对象初始化失败了，这个 Do 函数也不会被再次执行了
 */