package services

type Info interface {
	ToString() string
	ToJson() (string, error)
}