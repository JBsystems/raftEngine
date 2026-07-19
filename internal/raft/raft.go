package raft

type NodeState int

const (
	StateIdle NodeState = iota
	StateConnected
	StateError
	StateRetrying
)
