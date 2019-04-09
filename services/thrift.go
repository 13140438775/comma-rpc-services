package services

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/astaxie/beego/logs"
)

type Thrift struct {
	listenAddr string

	Log *logs.BeeLogger

	server *thrift.TSimpleServer
}

func NewThrift(listenAddr string, processor thrift.TProcessor) (*Thrift, error) {
	t := &Thrift{
		listenAddr: listenAddr,
		Log:        logs.NewLogger(),
	}

	transport, err := thrift.NewTServerSocketTimeout(t.listenAddr, 0)
	if err != nil {
		return nil, err
	}
	t.server = thrift.NewTSimpleServer2(processor, transport)

	return t, nil
}

func (t *Thrift) Run() {
	t.Log.Info("comma-rpc-service run...")
	t.server.Serve()
}

func (t *Thrift) Exit() {
	t.Log.Info("comma-rpc-service stop...")
	t.server.Stop()
}
