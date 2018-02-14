package procotol

import (
	"proxy-web/util"
)

type Protocol interface {
	GetCommand(data *util.Parameter) (string, error)
}
