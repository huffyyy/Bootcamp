package employee

type Info interface {
	ToString() string
	ToJson() (string, error)
}