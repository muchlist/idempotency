package main

import (
	"idempotency/pkg/global"
	"idempotency/pkg/mlogger"
	"idempotency/pkg/web"
)

func main() {
	// init log
	contextField := map[string]any{
		"request_id": global.RequestIDKey,
		// "trace_id":   global.TraceIDKey,
	}

	log := mlogger.New(mlogger.Options{
		Level:        mlogger.LevelInfo,
		Output:       "stdout",
		ContextField: contextField,
	})

	// create and start api server
	webApi := web.New(log, 8080, "stg", "idempotency")
	err := webApi.Serve(mapRoutes(log))
	if err != nil {
		log.Error("serve web api", err)
	}
}
