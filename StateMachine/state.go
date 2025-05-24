package StateMachine

type State interface {
	Order() int
	Name() string

	PreTransition(ctx *Context) error
	PostTransition(ctx *Context) error
	Transition(ctx *Context) error
}
