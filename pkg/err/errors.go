package err

import (
	"github.com/cockroachdb/errors"
)

var (
	ErrNotFoundEnv = errors.New("Not Found Env")
)
