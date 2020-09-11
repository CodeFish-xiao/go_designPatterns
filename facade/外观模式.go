package facade

import "fmt"

// CPU
type CPU struct {
}

func (*CPU) start() {
	fmt.Println("启动CPU。。。")
}

// 内存
type Memory struct {
}

func (*Memory) start() {
	fmt.Println("启动内存管理。。。")
}

// 硬盘
type Disk struct {
}

func (*Disk) start() {
	fmt.Println("启动硬盘。。。")
}

// 开机键
type StartBtn struct {
}

func (*StartBtn) start() {
	cpu := &CPU{}
	cpu.start() //cpu启动
	memory := &Memory{}
	memory.start() //内存启动
	disk := &Disk{}
	disk.start() //硬盘启动
}
