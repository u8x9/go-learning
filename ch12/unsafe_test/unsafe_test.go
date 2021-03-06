package unsafe_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 危险的行为
func TestUnsafe(t *testing.T) {
	i := 10
	f := *((*float64)(unsafe.Pointer(&i)))
	t.Log(f)
}

type MyInt int

// 合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

// 原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		t.Log(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
		}()
	}
	wg.Wait()
}
