package set

type Set interface {
	Name() string
	SAdd(string) bool
	SCard() int64
	SisMember(string) bool
	SMembers() []string
	SRem(string)
}
