package singleton
import "sync"

var (
	instance *Instance
	lock     sync.Mutex
)

type Instance struct {
	Name string
}

// 双重检查
func GetInstance(name string) *Instance {
	if instance == nil {//为空加锁赋值
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Instance{Name: name}
		}
	}
	return instance
}
