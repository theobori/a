package ql

import (
	"time"

	"github.com/theobori/a/filename"
)

type Context struct {
	FilenameMetadata *filename.Metadata
	Date             time.Time
}
