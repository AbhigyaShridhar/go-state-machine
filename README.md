# go-state-machine

A simple and extensible finite state machine (FSM) library written in Go. Ideal for workflows like transaction 
management systems, or other sequential operations which can be broken down into stages.

## 📦 Features

- Structured state transitions with lifecycle hooks
- Context-based data passing and validation
- Safe concurrent usage with mutex locking
- Developer-friendly error types with structured information
- Easy to extend with existing state management in a codebase


## 🚀 Quick Start

```bash
go get github.com/AbhigyaShridhar/go-state-machine
```

### 🔧 Usage

```go
package main

import (
    "github.com/AbhigyaShridhar/go-state-machine/StateMachine"
)

func main() {
    type WaitingState struct {
        // implement the State interface
    }
    sm := StateMachine.NewStateMachine()
    err := sm.RegisterState(&WaitingState{})
    if err != nil {
        return
    }
    // ...
    err = sm.PerformTransition("waiting", "completed", data)
    if err != nil {
        // Do relevant error handling
    }
}
```
Navigate to the [atm example](./examples/atm/atm.go) for more understanding on the implementation of state machine package.
Run `go run examples/atm/atm.go` to get started.
