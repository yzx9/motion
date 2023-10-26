package user

/**
 * Value Object
 */
type Avatar interface {
	URL() string
	Valid() bool
}

type avatar struct{}

func (avatar) URL() string   { return "" }
func (a avatar) Valid() bool { return a.URL() != "" }

var emptyAvatar = avatar{}
