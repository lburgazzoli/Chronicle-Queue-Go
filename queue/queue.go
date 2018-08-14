package queue

import "github.com/OpenHFT/Chronicle-Queue-Go/core"

// Document --
type Document interface {
}

// Tailer used to read sequentially
type Tailer interface {
	core.Disposable

	ReadingDocument() Document
}

// Appender used to write new excerpts sequentially to the upper.
type Appender interface {
	core.Disposable

	WrintingDocument() Document
}

// Queue --
type Queue interface {
	core.Disposable

	Tailer() Tailer
	Appender() Appender
}
