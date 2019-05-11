package loom

type Config interface {
	DPOS() DPOSConfig
	GetConfig(string) string
}

type DPOSConfig interface {
	FeeFloor(int64) int64
}
