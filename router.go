package be

type Router interface {
	Execute(*Context)
}
