package jeager

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

func initJeager(service string) io.Closer {
	cfg := jaegercfg.Configuration{
		// 将采样频率设置为1，每一个span都记录，方便查看测试结果
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
			// 将span发往jaeger-collector的服务地址
			CollectorEndpoint: "http://localhost:14268/api/traces",
		},
	}
	closer, err := cfg.InitGlobalTracer(service, jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return closer
}

func Jeager() gin.HandlerFunc {
	return func(c *gin.Context) {
		closer := initJeager("test")
		defer closer.Close()
		// 获取jaeger tracer
		t := opentracing.GlobalTracer()
		// 创建root span
		sp := t.StartSpan("in-process-service")
		// main执行完结束这个span
		defer sp.Finish()
		// 将span传递给Foo
		opentracing.ContextWithSpan(c, sp)
	}
}
