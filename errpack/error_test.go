package errpack_test

import (
	"errors"
	"testing"

	"github.com/238Studio/child-nodes-error-manager/errpack"
)

func TestError(t *testing.T) {
	err := errpack.NewError(errpack.NotException, errpack.Assist, errors.New("test"))

	t.Log(errpack.ErrToStruct(err))
	t.Log(errpack.ExtractErrorLevel(err))
	t.Log(errpack.ExtractErrorModule(err))
	t.Log(errpack.ExtractErrorMessage(err))
	t.Log(errpack.ExtractStackTrace(err))
	t.Log(errpack.ExtractErrorUUID(err))
}
