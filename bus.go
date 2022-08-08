package bus

import (
	"errors"
	"sync"
	"time"

	"github.com/chefsgo/bus"
	"github.com/chefsgo/util"
)

var (
	errRunning = errors.New("Bus is running")
)

type (
	defaultDriver  struct{}
	defaultConnect struct {
		mutex   sync.RWMutex
		running bool
		actives int64

		name     string
		config   bus.Config
		delegate bus.Delegate

		runner *util.Runner
	}
)

// 连接
func (driver *defaultDriver) Connect(name string, config bus.Config) (bus.Connect, error) {
	return &defaultConnect{
		name: name, config: config, runner: util.NewRunner(),
	}, nil
}

// 打开连接
func (connect *defaultConnect) Open() error {
	return nil
}
func (connect *defaultConnect) Health() (bus.Health, error) {
	connect.mutex.RLock()
	defer connect.mutex.RUnlock()
	return bus.Health{Workload: connect.actives}, nil
}

// 关闭连接
func (connect *defaultConnect) Close() error {
	connect.runner.End()
	return nil
}

func (connect *defaultConnect) Accept(delegate bus.Delegate) error {
	connect.mutex.Lock()
	defer connect.mutex.Unlock()

	connect.delegate = delegate

	return nil
}

// 单机版 不用做处理，因为根本就调用不到这边
func (connect *defaultConnect) Register(name string) error {
	connect.mutex.Lock()
	defer connect.mutex.Unlock()

	return nil
}

// 开始订阅者
func (connect *defaultConnect) Start() error {
	if connect.running {
		return errRunning
	}

	connect.running = true
	return nil
}

// 本来不会执行到这
func (connect *defaultConnect) Request(name string, data []byte, timeout time.Duration) ([]byte, error) {
	return nil, nil
}
