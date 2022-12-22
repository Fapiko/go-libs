package utils

import (
	"io"
)

func NoErrCloseWrapper(closer io.Closer) {
	_ = closer.Close()
}
