package websock

const (
	DefaultReadDeadline     = 60 * 3  // 3 minutes
	DefaultReadLimit        = 1 << 21 //  2048K
	DefaultReadBufferSize   = 4096
	DefaultWriteBufferSize  = 4096
	DefaultAuthReadDeadline = 10 // 5 seconds
)
