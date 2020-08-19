package engine

type Config struct {
	AutoStart         bool
	DisableEncryption bool
	DownloadDirectory string
	DownloadType      string
	EnableUpload      bool
	EnableSeeding     bool
	IncomingPort      int
}
