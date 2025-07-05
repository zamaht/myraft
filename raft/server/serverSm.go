package server

import (
	"fmt"
	"net"
)

type localState struct {
	lastAppliedIdx int
	commitIdx      int
	nextIdx        int
	matchIdx       []int
}

type globalState struct {
	currentTerm int
	votedFor    int
	logs        []LogEntry
}

type State struct {
	globalState globalState
	localState  localState
}

type StateMachine struct {
	Server

	Id   int
	Role ServerRole

	running bool

	currentState State

	appendEntry chan AppendEntryRequest
	requestVote chan VoteRequest

	stopRequest chan bool
}

func NewServerSm(id int) *StateMachine {
	sm := new(StateMachine)

	sm.Id = id
	sm.Role = Follower

	return sm
}

func (sm *StateMachine) Run(address string, port int) {
	sm.init()
	listener := sm.listenOn(address, port)
	defer listener.Close()

	go sm.handleChannelsMsg()

	for sm.running {
		connection := manageConnections(listener)

		sm.handleReceivedClientMsg(connection)
	}
}

func (sm *StateMachine) Stop() {
	sm.running = false
}

func (sm *StateMachine) init() {
	fmt.Printf("Init for Server: %d\n", sm.Id)

	sm.Role = Follower
	sm.running = true

	sm.currentState.globalState = globalState{}
	sm.currentState.localState = localState{}

	sm.appendEntry = make(chan AppendEntryRequest)
	sm.requestVote = make(chan VoteRequest)
}

func (sm *StateMachine) applyRules() {
	fmt.Printf("Applying Rules for Server: %d\n", sm.Id)

	if sm.Role == Follower {
		applyFollowerRules()
	} else if sm.Role == Leader {
		applyLeaderRules()
	}
}

func applyFollowerRules() {
	// Wait for heartbeat
	// If timeout intiate election
	// look for RequestVote channel and send vote
}

func applyLeaderRules() {
	// Send heartbeat to every follower every 3 secs
}

func applyCandidateRules() {
	// Send RequestVote to every follower
}

func (sm *StateMachine) handleReceivedClientMsg(connection net.Conn) {
	fmt.Println("Start Reading.")
	buffer := make([]byte, 1024)
	connection.Read(buffer)
	fmt.Println("Read: ", string(buffer))
}

func (sm *StateMachine) listenOn(address string, port int) net.Listener {
	listener, err := net.Listen(SERVER_TYPE, address+":"+fmt.Sprint(port))
	if err != nil {
		fmt.Println("Error while start listening:", err.Error())
		panic(err)
	}

	fmt.Printf("Listening on: %s:%d\n", address, port)

	return listener
}

func (sm *StateMachine) handleChannelsMsg() {
	for {
		select {
		case <-sm.appendEntry:
			sm.handleAppendEntryMsg()
		case <-sm.requestVote:
			sm.handleRequestVoteMsg()
		case stopRequested := <-sm.stopRequest:
			sm.running = !stopRequested
		}
	}
}

func (sm *StateMachine) handleAppendEntryMsg() {

}

func (sm *StateMachine) handleRequestVoteMsg() {

}

func manageConnections(listener net.Listener) net.Conn {
	connection, err := listener.Accept()
	if err != nil {
		fmt.Println("Error while accepting connection:", err.Error())
	}

	return connection
}
