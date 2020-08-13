package singleton

import "sync"

/*
使用懒惰模式的单例模式，使用双重检查加锁保证线程安全
*/

//Singleton 是单例模式类
type Singleton struct {
	Version string
}

var (
	singleton *Singleton
	once      sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{Version: "1.0.0"}
	})
	return singleton
}
