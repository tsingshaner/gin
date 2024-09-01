package shared

import (
	"time"

	"github.com/tsingshaner/go-pkg/util"
	"gorm.io/gorm"
)

type (
	ID = uint64

	Model struct {
		IDModel
		TimeModel
		SoftDeletedModel
	}

	IDModel struct {
		ID ID `json:"id" gorm:"type:bigint;primarykey;autoIncrement"`
	}

	TimeModel struct {
		CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
		UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	}

	SoftDeletedModel struct {
		DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	}
)

func NewModel(opts ...util.WithFn[Model]) *Model {
	return util.BuildWithOpts(&Model{}, opts...)
}

func WithModelID(id ID) util.WithFn[Model] {
	return func(m *Model) {
		m.ID = id
	}
}
