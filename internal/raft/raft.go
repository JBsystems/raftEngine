package raft

import "sync"

type NodeState int

const (
	Follower NodeState = iota
	Candidate
	Leader
)

type RaftNode struct {
	mu sync.Mutex

	nodeId int
	peers  []string

	currentTerm int
	votedFor    int
	log         []byte

	commitIndex int
	lastApplied int
	state       NodeState
}
