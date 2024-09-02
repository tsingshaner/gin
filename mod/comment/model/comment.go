package model

import (
	"github.com/tsingshaner/gin/shared"
)

type Comment struct {
	shared.Model
	shared.TimeModel
	shared.SoftDeletedModel
}
