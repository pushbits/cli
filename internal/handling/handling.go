package handling

import (
	"io"

	log "github.com/sirupsen/logrus"
)

func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Warn(err)
	}
}
