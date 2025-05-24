package StateMachine

import (
	"errors"
	"fmt"
	"sync"
)

// TransitionError provides structured error information for transition failures.
type TransitionError interface {
	error
	Stage() string
	Unwrap() error
}

type transitionErrorImpl struct {
	stage string
	err   error
}

func (e *transitionErrorImpl) Error() string {
	return fmt.Sprintf("%s: %v", e.stage, e.err)
}

func (e *transitionErrorImpl) Unwrap() error {
	return e.err
}

func (e *transitionErrorImpl) Stage() string {
	return e.stage
}

func newTransitionError(stage string, err error) TransitionError {
	if err == nil {
		return nil
	}
	return &transitionErrorImpl{
		stage: stage,
		err:   err,
	}
}

// StateMachine struct is responsible for actual transition methods
type StateMachine struct {
	stateMap map[string]State
	mutex    sync.Mutex
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		stateMap: make(map[string]State),
		mutex:    sync.Mutex{},
	}
}

func (sm *StateMachine) RegisterState(stage State) error {
	name := stage.Name()
	ok := sm.stateMap[name]
	if ok != nil {
		return errors.New(fmt.Sprintf("state %s already registered", name))
	}
	sm.stateMap[name] = stage
	return nil
}

func (sm *StateMachine) PerformTransition(from, to string, globalData map[string]interface{}) TransitionError {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	startState, ok := sm.stateMap[from]
	if !ok {
		return newTransitionError("validation", errors.New(fmt.Sprintf("state %s not registered", from)))
	}

	targetState, ok := sm.stateMap[to]
	if !ok {
		return newTransitionError("validation", errors.New(fmt.Sprintf("state %s not registered", to)))
	}

	if startState.Order() >= targetState.Order() {
		return newTransitionError("validation", errors.New(fmt.Sprintf(
			"transition from %s to %s failed", startState.Name(), targetState.Name())))
	}

	ctx := &Context{globalData}
	err := targetState.PreTransition(ctx)
	if err != nil {
		return newTransitionError("pre-transition", err)
	}

	err = targetState.Transition(ctx)
	if err != nil {
		return newTransitionError("transition", err)
	}

	err = targetState.PostTransition(ctx)
	if err != nil {
		return newTransitionError("post-transition", err)
	}
	return nil
}
