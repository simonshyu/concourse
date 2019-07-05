package transport

import (
	"context"
	"fmt"
	"github.com/concourse/retryhttp"
)

type UnreachableWorkerRetryer struct {
	DelegateRetryer retryhttp.Retryer
}

func (r *UnreachableWorkerRetryer) IsRetryable(err error) bool {

	if err == context.DeadlineExceeded{
		fmt.Println("ctx_log: IsRetryable DeadlineExceeded")
		return false
	}

	if _, ok := err.(WorkerUnreachableError); ok {
		return true
	}

	return r.DelegateRetryer.IsRetryable(err)
}
