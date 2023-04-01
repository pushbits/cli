// Package handling provides convenience functions for cleaning up resources.
package handling

import (
	"io"

	log "github.com/sirupsen/logrus"
)

// Close closes an io resource and prints a warning if that fails.
func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Warn(err)
	}
}
