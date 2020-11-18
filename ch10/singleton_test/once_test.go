package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 单例模式（懒汉式，线程安全）

type singleton struct{}

var singletonInstance *singleton

var once sync.Once

func getSingletonObj() *singleton {
	once.Do(func() {
		fmt.Println("创建对象")
		singletonInstance = new(singleton)
	})
	return singletonInstance
}

func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s := getSingletonObj()
			t.Logf("[#%d] address: %p\n", i, unsafe.Pointer(s))
		}(i)
		wg.Wait()
	}
}
