package rpc

type RequestVoteArgs struct {
	CurrentTerm  int
	NodeVote     int
	LastLogIndex int
	LastLogTerm  int
}

type RequestVoteReply struct {
	OwnCurrTerm int
	VoteGranted bool
}

type AppendEntriesArgs struct {
	LeadCurrentTerm   int
	LeadNodeId        int
	PrevLogIndex      int
	PrevLogTerm       int
	RepEntries        []byte
	LeadHiComLogIndex int
}

type AppendEntriesReply struct {
	FollowerCurrTerm int
	LogCheck         bool
}
