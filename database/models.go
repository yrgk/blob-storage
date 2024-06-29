package database

import "time"

type (
	Bucket struct {
		ID   uint `gorm:"primarykey"`
		Name string
	}

	File struct {
		ID   uint `gorm:"primarykey"`
		Data []byte
	}

	FileMeta struct {
		ID       uint `gorm:"primarykey"`
		Name     string
		FileID   uint
		BucketID uint
	}

	FileLink struct {
		ID         uint `gorm:"primarykey"`
		Link       string
		Expiration time.Time
		FileID     uint
	}
)
