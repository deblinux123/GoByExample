package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (s ServerState) String() string {
	return stateName[s]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println("Current State:", ns)

	ns2 := transition(ns)
	fmt.Println("Current State:", ns2)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateRetrying
	default:
		panic(fmt.Errorf("Unknown state: %s", s))
	}
}
