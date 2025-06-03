

package error_message

import "errors"

var (
	ErrWrongPassword = errors.New("incorrect password entered")
	ErrNoResultSet   = errors.New("no rows in result set")
)
