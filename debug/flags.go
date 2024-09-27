package debug

type Flags struct {
	ReturnDetailError bool
	DbLog             bool
	RequestLog        bool
	ResponseLog       bool
	RpcLog            bool
	LogLevel          string
}
