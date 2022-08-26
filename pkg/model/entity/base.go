package entity

import (
	"gorm.io/plugin/optimisticlock"
	"time"
)

type BaseDO struct {
	Version          optimisticlock.Version `gorm:"version"`             // 乐观锁
	Id               uint                   `gorm:"primarykey"`          // 主键ID
	CreatedTime      time.Time              `gorm:"created_time"`        // 创建时间
	CreatedBy        string                 `gorm:"created_by"`          // 创建人
	LastModifiedTime time.Time              `gorm:"last_modified_time"`  // 最后更新时间
	LastModifiedBy   string                 `gorm:"last_modified_by"`    // 最后更新人
	Deleted          bool                   `gorm:"is_deleted" json:"-"` // 删除时间
}
