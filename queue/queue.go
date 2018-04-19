package queue

// DocumentCtx --
type DocumentCtx interface {
}

// Tailer used to read sequentially
type Tailer interface {
	ReadingDocument() DocumentCtx
}

// Appender used to write new excerpts sequentially to the upper.
type Appender interface {
	WrintingDocument() DocumentCtx
}

// Queue --
type Queue interface {
	Tailer() Tailer
	Appender() Appender
}
