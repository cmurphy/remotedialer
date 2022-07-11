package remotedialer

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func log(format string, a ...interface{}) {
	logrus.Debugf("[remotedialer] %s", fmt.Sprintf(format, a...))
}
