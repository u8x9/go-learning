package obj_cache

import (
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Create a new object.")
			return 100
		},
	}

	v := pool.Get().(int)
	t.Log(v)
	pool.Put(3)
	//runtime.GC() // GC 会清除 sync.pool 中缓存的对象[版本差异]
	v1, _ := pool.Get().(int)
	t.Log(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Create a new object.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(200)
	pool.Put(300)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			v := pool.Get().(int)
			t.Logf("[#%d] got: %d\n", id, v)
		}(i)
	}
	wg.Wait()
}
