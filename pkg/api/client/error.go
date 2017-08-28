package client

import (
	"fmt"
	"net/http"

	"github.com/caicloud/helm-registry/pkg/errors"
)

type HTTPCodedError interface {
	String() string
	Error() string
	Code() int
	Desc() string
}

type httpCodedError struct {
	code int
	desc string
}

var _ HTTPCodedError = &httpCodedError{}

func errorFromResponse(resp *http.Response) error {
	return &httpCodedError{resp.StatusCode, resp.Status}
}

func errorFromHelmRegistry(err error) error {
	if e, ok := err.(*errors.Error); ok {
		return &httpCodedError{e.Code, e.Message}
	}
	return err
}

func (e *httpCodedError) Error() string {
	return fmt.Sprintf("Code:%q, Desc:%q", e.code, e.desc)
}

func (e *httpCodedError) String() string {
	return fmt.Sprintf("Code:%q, Desc:%q", e.code, e.desc)
}

func (e *httpCodedError) Code() int {
	return e.code
}

func (e *httpCodedError) Desc() string {
	return e.desc
}
