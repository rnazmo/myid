package myid

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Generate generates ID.
//
// ID Spec:
//   UUID v4
//
func Generate() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", errors.Wrap(err, "failed to uuid.NewRandom()")
	}
	return id.String(), nil
}
