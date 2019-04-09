package main

import (
	"comma-rpc-services/services"
	"comma-rpc-services/thrift/schedule"
	"github.com/astaxie/beego/config"
	"comma-rpc-services/dao"
	"fmt"
	"time"
	"github.com/astaxie/beego/logs"
	"syscall"
	"github.com/judwhite/go-svc/svc"
	"comma-rpc-services/helper"
	"github.com/opentracing/opentracing-go"
)

type program struct {
	Thrift *services.Thrift
}

func main() {
	prg := &program{}
	if err := svc.Run(prg, syscall.SIGINT, syscall.SIGTERM); err != nil {
		fmt.Println(err)
	}
}

func (p *program) Init(env svc.Environment) (err error) {
	cnf, err := config.NewConfig("ini", "conf/conf.ini")
	if err != nil {
		fmt.Printf("config.NewConfig err, path:conf/conf.ini, err: %s\n", err)
		return
		logs.Info("")
	}
	if _, err = time.LoadLocation(cnf.String("TimeLocation")); err != nil {
		fmt.Printf("time.LoadLocation err, TimeLocation: %s, err: %s\n", cnf.String("TimeLocation"), err)
		return
	}
	if err = dao.Init(cnf.String("DbUrl")); err != nil {
		fmt.Printf("dao.Init err, DbUrl: %s, err: %s\n", cnf.String("DbUrl"), err)
		return
	}

	trace := helper.NewJaegerTracer(
		cnf.String("OpentracingDomain"),
		cnf.String("OpentracingUsername"),
		cnf.String("OpentracingPassword"),
		cnf.String("OpentracingName"),
	)
	opentracing.SetGlobalTracer(trace)

	handler := &services.ScheduleService{}
	processor := schedule.NewScheduleProcessor(handler)
	p.Thrift, err = services.NewThrift(cnf.String("ListenAddr"), processor)
	if err != nil {
		fmt.Printf("new thrift err: %s\n", err)
		return
	}

	p.Thrift.Log.SetLogger(logs.AdapterFile, cnf.String("FileLog"))
	p.Thrift.Log.SetLevel(cnf.DefaultInt("FileLogLevel", 7))
	return
}

func (p *program) Start() error {
	p.Thrift.Run()
	return nil
}

func (p *program) Stop() error {
	if p.Thrift != nil {
		p.Thrift.Exit()
	}
	return nil
}
