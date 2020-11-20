package micro_kernel

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type DemoCollector struct {
	evtReceiver EventReceiver
	agtCtx      context.Context
	stopChan    chan struct{}
	name        string
	content     string
}

func NewCollect(name string, content string) *DemoCollector {
	return &DemoCollector{
		stopChan: make(chan struct{}),
		name:     name,
		content:  content,
	}
}

func (c *DemoCollector) Init(evtReceiver EventReceiver) error {
	fmt.Println("initialize collector", c.name)
	c.evtReceiver = evtReceiver
	return nil
}

func (c *DemoCollector) Start(agtCtx context.Context) error {
	fmt.Println("start collector", c.name)
	for {
		select {
		case <-agtCtx.Done():
			c.stopChan <- struct{}{}
			break
		default:
			time.Sleep(time.Millisecond * 50)
			c.evtReceiver.OnEvent(Event{c.name, c.content})
		}
	}
	return nil
}

func (c *DemoCollector) Stop() error {
	fmt.Println("stop collector", c.name)
	select {
	case <-c.stopChan:
		return nil
	case <-time.After(time.Second):
		return errors.New("failed to stop for timeout")
	}
	return nil
}

func (c *DemoCollector) Destory() error {
	fmt.Println(c.name, "released resources.")
	return nil
}

func TestAgent(t *testing.T) {
	agt := NewAgent(100)
	c1 := NewCollect("foo", "锄和日当午")
	c2 := NewCollect("bar", "汗滴禾下土")
	agt.RegisterCollector(c1.name, c1)
	agt.RegisterCollector(c2.name, c2)
	if err := agt.Start(); err != nil {
		t.Fatal("Start error", err)
	}
	t.Log(agt.Start())
	time.Sleep(time.Second)
	agt.Stop()
	agt.Destory()
}
