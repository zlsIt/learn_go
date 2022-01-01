package singleton

import "sync"

// 单例
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func getInstance() *Singleton {
	return singleton
}

// 懒汉式（双重检测）
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func getLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
