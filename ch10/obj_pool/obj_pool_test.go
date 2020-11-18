package obj_pool

import (
	"testing"
	"time"
)

func TestObjectPool(t *testing.T) {
	pool := newObjectPool(10)
	// // 尝试放置超出池大小的对象
	// if err := pool.release(&reusableObject{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 1; i <= 11; i++ {
		if obj, err := pool.get(time.Second); err != nil {
			t.Error(err)
		} else {
			t.Logf("#%d, %p\n", i, obj)
			// 注释掉以下释放对象的代码，将在第11次 get 时引发超时
			if err := pool.release(obj); err != nil {
				t.Error(err)
			}
		}
	}
	t.Log("Done")
}
