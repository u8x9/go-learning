package obj_pool

import (
	"errors"
	"time"
)

type reusableObject struct{}
type objectPool struct {
	bufChan chan *reusableObject
}

func newObjectPool(numOfObj int) *objectPool {
	pool := objectPool{}
	pool.bufChan = make(chan *reusableObject, numOfObj)
	for i := 0; i < numOfObj; i++ {
		pool.bufChan <- &reusableObject{}
	}
	return &pool
}

func (p *objectPool) get(timeout time.Duration) (*reusableObject, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
	return nil, errors.New("can not get object from pool")
}
func (p *objectPool) release(obj *reusableObject) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
	return errors.New("can not release object to pool")
}
