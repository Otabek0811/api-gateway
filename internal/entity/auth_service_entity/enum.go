package auth_service_entity

type UserStatus int

const (
	ACTIVE UserStatus = iota
	INVERIFY
	BLOCKED
)

func (s UserStatus) String() string {
	switch s {
	case ACTIVE:
		return "ACTIVE"
	case INVERIFY:
		return "INVERIFY"
	case BLOCKED:
		return "BLOCKED"
	default:
		return "UNKNOWN"
	}
}
