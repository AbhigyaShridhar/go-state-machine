package main

import (
	"errors"
	"fmt"
	"github.com/AbhigyaShridhar/go-state-machine/StateMachine"
)

type WaitingState struct{}

func (w *WaitingState) Order() int   { return 1 }
func (w *WaitingState) Name() string { return "waiting" }
func (w *WaitingState) PreTransition(ctx *StateMachine.Context) error {
	fmt.Println("[Waiting] PreTransition: Authenticating user...")
	ctx.Data["authenticated"] = true
	ctx.Data["referenceID"] = "TXN123456"
	return nil
}
func (w *WaitingState) Transition(ctx *StateMachine.Context) error {
	if ctx.Data["authenticated"] == true {
		fmt.Println("[Waiting] User has been authenticated")
	} else {
		return errors.New("could not authenticate user")
	}
	fmt.Printf("[Waiting] Transition: User authenticated. Ref ID created.")
	return nil
}
func (w *WaitingState) PostTransition(ctx *StateMachine.Context) error {
	fmt.Printf("[Waiting] PostTransition: Ready to send withdrawal request for reference ID: %s.",
		ctx.Data["referenceID"])
	return nil
}

type CompletedState struct{}

func (p *CompletedState) Order() int   { return 2 }
func (p *CompletedState) Name() string { return "completed" }
func (p *CompletedState) PreTransition(ctx *StateMachine.Context) error {
	// this is where you would validate that the reference id is approved
	// this would commonly include a database or API call
	fmt.Printf("[Completed] PreTransition: Sending withdrawal request to bank with data %v...",
		ctx.Data)
	return nil
}
func (p *CompletedState) Transition(ctx *StateMachine.Context) error {
	fmt.Println("[Completed] Transition: Bank processing started...")
	ctx.Data["bankApproved"] = true
	return nil
}
func (p *CompletedState) PostTransition(ctx *StateMachine.Context) error {
	// This is where asynchronous updates after a transaction would be implemented
	fmt.Println("[Completed] PostTransition: Bank accepted transaction.")
	fmt.Printf("[Completed] PostTransition: data sent for post transaction processing %s", ctx.Data)
	return nil
}

func main() {
	atmSM := StateMachine.NewStateMachine()

	states := []StateMachine.State{
		&WaitingState{},
		&CompletedState{},
	}

	for _, state := range states {
		err := atmSM.RegisterState(state)
		if err != nil {
			fmt.Printf("Error registering state %s: %v\n", state.Name(), err)
		}
	}

	data := map[string]interface{}{
		"cardNumber": "1234-5678-9012-3456",
		"amount":     500.0,
	}

	err := atmSM.PerformTransition("waiting", "completed", data)
	if err != nil {
		fmt.Printf("Transition failed at stage %s: %v\n", err.Stage(), err.Error())
		return
	}

	fmt.Println("Transition successful.")
	fmt.Printf("Final transaction data: %+v\n", data)
}
