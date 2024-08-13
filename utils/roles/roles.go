package roles

import "strings"

// ? Enum or Role type

const (
	Admin     string = "admin"
	Moderator string = "moderator"
)

func ToDBString(roles []string) string {
	return strings.Join(roles, ",")
}

func FromDBString(roles string) []string {
	return strings.Split(roles, ",")
}
