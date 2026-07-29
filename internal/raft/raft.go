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

func NewRaftNode(nodeID int, peers []string) *RaftNode {
	return &RaftNode{
		nodeID:      nodeID,
		peers:       peers,
		votedFor:    -1,
		currentTerm: 0,
		commitIndex: 0,
		lastApplied: 0,
		state:       Follower,
	}
}
