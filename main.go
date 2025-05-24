package main

import (
	"github.com/AbhigyaShridhar/go-state-machine/StateMachine"
)

type (
	Context         = StateMachine.Context
	State           = StateMachine.State
	StateMachineAPI = StateMachine.StateMachine
	TransitionError = StateMachine.TransitionError
)

var (
	NewStateMachine = StateMachine.NewStateMachine
)
