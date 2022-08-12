package singleton

var singleton *Singleton

// Singleton 饿汉式
type Singleton struct{}

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
