package singleton

import "sync"

var lazySingleton *Singleton
var once sync.Once

// SingletonLazy 饿汉式

// 懒汉式
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}

	return lazySingleton
}
