package netobj

import (
	"github.com/fluofoxxo/outrun/obj"
)

type Chao struct {
	obj.Chao
	Status   int64 `json:"status"` // enums.ChaoStatus*
	Level    int64 `json:"level"`
	Dealing  int64 `json:"setStatus"` // enums.ChaoDealing*
	Acquired int64 `json:"acquired"`  // flag
}

func NewNetChao(chao obj.Chao, status, level, dealing, acquired int64) Chao {
	return Chao{
		chao,
		status,
		level,
		dealing,
		acquired,
	}
}
