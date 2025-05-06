package utils

type AuthLevel int

const (
	AuthLevelNone AuthLevel = iota
	AuthLevelUser
	AuthLevelAdmin
)
