package server

type LogEntry string

type AppendEntryRequest struct {
	term        int
	leaderId    int
	prevLogIdx  int
	prevLogTerm int
	entries     []LogEntry

	leaderCommitIdx int
}

type AppendEntryResponse struct {
	term    int
	success bool
}

type VoteRequest struct {
	term        int
	candidateId int
	lastLogIdx  int
	lastLogTerm int
}

type VoteResponse struct {
	term         int
	votedGranted bool
}
