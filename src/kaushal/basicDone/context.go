package main

import (
	"context"
	"fmt"
)

type TraceId string

type TraceIdKey int

func main() {
	traceId := TraceId("iuyio-giu87-gkjhiuh-987hkl")

	traceIdKey := TraceIdKey(0)

	ctx := context.WithValue(context.Background(), traceIdKey, traceId)

	if uuid, ok := ctx.Value(traceIdKey).(TraceId); ok {
		fmt.Println("TraceId : ", uuid)
	}

}
