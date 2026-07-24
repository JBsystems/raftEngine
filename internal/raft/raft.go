package raft

import (
	"sync"
	"time"
)

type nodeState int

const (
	Follower nodeState = iota
	Candidate
	Leader
)

type raftNode struct {
	mu sync.Mutex

	nodeId int
	peers  []string

	currentTerm int
	votedFor    int
	log         []byte

	commitIndex int
	lastApplied int
	state       nodeState

	timer *time.Timer

	ticker *time.Ticker
}
