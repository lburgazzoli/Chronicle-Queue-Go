package queue

// Document --
type Document interface {
}

// Tailer used to read sequentially
type Tailer interface {
	ReadingDocument() Document
}

// Appender used to write new excerpts sequentially to the upper.
type Appender interface {
	WrintingDocument() Document
}

// Queue --
type Queue interface {
	Tailer() Tailer
	Appender() Appender
}
