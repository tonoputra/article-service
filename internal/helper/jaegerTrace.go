package helper

import (
	"io"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type JaegerTrace interface {
	Init() (opentracing.Tracer, io.Closer, error)
}

type jaegerTrace struct {
}

func NewJaeger() JaegerTrace {
	return &jaegerTrace{}
}

// Init the jaeger
// func (jaegerTrace) Init() {
// 	sn := os.Getenv("SERVICE_NAME")
// 	if sn == "" {
// 		log.Panicln("ERROR: env variable SERVICE_NAME must be fill")
// 	}
// 	// cfg := config.Configuration{
// 	// 	Throttler: &config.ThrottlerConfig{
// 	// 		HostPort: "16686",
// 	// 	},
// 	// 	// ServiceName: service,
// 	// 	Sampler: &config.SamplerConfig{
// 	// 		Type:  "const",
// 	// 		Param: 1,
// 	// 	},
// 	// 	Reporter: &config.ReporterConfig{
// 	// 		LogSpans:            true,
// 	// 		BufferFlushInterval: 10 * time.Second,
// 	// 	},
// 	// }

// 	// // tracer, closer, err := cfg.NewTracer(
// 	// // 	config.Logger(jaeger.StdLogger),
// 	// // )
// 	// // if err != nil {
// 	// // 	log.Panicln("ERROR: Cannot init jaeger:", err)
// 	// // }

// 	// // opentracing.SetGlobalTracer(tracer)
// 	// // defer closer.Close()

// 	// tracer, closer, err := cfg.New(
// 	// 	"article-service",
// 	// 	config.Logger(jaeger.StdLogger),
// 	// )
// 	// if err != nil {
// 	// 	log.Panicln("ERROR: Cannot init jaeger:", err)
// 	// }

// 	// opentracing.SetGlobalTracer(tracer)
// 	// defer closer.Close()

// 	// Check Tracing
// 	// if tracing.Err == nil {
// 	// 	defer func() {
// 	// 		err := tracing.Closer.Close()
// 	// 		if err != nil {
// 	// 			log.Println("There was error when closing Open tracer")
// 	// 		}
// 	// 	}()
// 	// } else {
// 	// 	log.Println("Tracing error: ", tracing.Err)
// 	// }
// }

//Init initializing global variables in project for tracing, closer and error
func (jaegerTrace) Init() (opentracing.Tracer, io.Closer, error) {
	host := os.Getenv("JAEGER_HOST")
	port := os.Getenv("JAEGER_PORT")
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Panicln("ERROR: env variable SERVICE_NAME must be fill")
	}

	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			// LogSpans:           env.Get.IsDev(),
			LocalAgentHostPort: host + ":" + port,
		},
		Tags: []opentracing.Tag{
			{
				Key:   "go-env",
				Value: os.Getenv("GO_ENV"),
			},
		},
	}

	Tracer, Closer, Err := cfg.NewTracer(config.Logger(jaeger.StdLogger))

	if Err != nil {
		log.Println("ERROR: cannot init Jaeger:", Err)
	}

	return Tracer, Closer, Err
}

func ErrorSpan(s *opentracing.Span, err error) {
	if s != nil && *s != nil {
		(*s).SetTag("mai.status", "error")
		(*s).SetTag("mai.error", err.Error())
		//s.Finish() //Finish should handled by app
	}
}
