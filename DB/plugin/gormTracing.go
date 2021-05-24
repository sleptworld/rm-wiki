package plugin

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"gorm.io/gorm"
)
const (
	callBackBeforeName = "opentracing:before"
	callBackAfterName  = "opentracing:after"
)

type OpentracingPlugin struct {}

const gormSpanKey = "__gorm_span"

func before(db *gorm.DB){
	span,_ := opentracing.StartSpanFromContext(db.Statement.Context,"gorm")

	db.InstanceSet(gormSpanKey,span)
}

func after(db *gorm.DB){
	_span,isExist := db.InstanceGet(gormSpanKey)

	if !isExist{
		return
	}
	span,ok := _span.(opentracing.Span)

	if !ok{
		return
	}
	defer span.Finish()

	if db.Error != nil{
		span.LogFields(tracerLog.Error(db.Error))
	}
	span.LogFields(tracerLog.String("sql",db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))

	return
}

func (op *OpentracingPlugin) Name() string{
	return "openTracingPlugin"
}

func (op *OpentracingPlugin) Initialize(db *gorm.DB) (err error){
	db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func InitJeager(){
	cfg, err := jaegercfg.FromEnv()

	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		fmt.Println("Could not parse Jaeger env vars: %s", err.Error())
		return
	}

	cfg.ServiceName = "gorm"
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		fmt.Println("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)
}