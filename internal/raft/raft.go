package raft

import (
	"sync"
)

type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

type RaftNode struct {
	mu sync.Mutex

	nodeID int
	peers  []string

	currentTerm int
	votedFor    int

	commitIndex int
	lastApplied int
	state       NodeState
}

// func NewRaftNode() {
// }

// func (r *raftNode) Get
