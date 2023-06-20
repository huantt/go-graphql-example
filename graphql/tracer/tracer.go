package tracer

import (
	"context"
	"github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/huantt/go-graphql-sample/pkg/log"
	"time"
)

type CustomTracer struct {
}

func NewCustomTracer() CustomTracer {
	return CustomTracer{}
}

func (t CustomTracer) TraceQuery(ctx context.Context, queryString string, operationName string, variables map[string]interface{}, varTypes map[string]*introspection.Type) (context.Context, func([]*errors.QueryError)) {
	return ctx, func(errs []*errors.QueryError) {}
}

func (t CustomTracer) TraceField(ctx context.Context, label, typeName, fieldName string, trivial bool, args map[string]interface{}) (context.Context, func(*errors.QueryError)) {
	start := time.Now()
	return ctx, func(err *errors.QueryError) {
		log.Tracef("[%s] resolved in %s", label, time.Since(start))
		if err != nil {
			log.Errorf("[%s] Err: %s", label, err.Error())
		}
	}
}

func (t CustomTracer) TraceValidation(context.Context) func([]*errors.QueryError) {
	return func(errs []*errors.QueryError) {}
}
