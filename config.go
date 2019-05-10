package loom

type Config interface {
	DPOS() DPOS
}

type DPOS interface {
	FreeFloor() int64
}
