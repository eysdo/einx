package agent

import (
	"sync/atomic"
)

type AgentID uint64
type ProtoTypeID = uint32
type EventType = int

type Agent interface {
	GetID() AgentID
	WriteMsg(msg_id ProtoTypeID, msg interface{}) bool
	Close()
	Run()
	Destroy()
}

type AgentSesMgr interface {
	OnAgentEnter(AgentID, Agent)
	OnAgentExit(AgentID, Agent)
	OnMgrError(EventType, error)
}

var agent_id uint64 = 0

func GenAgentID() AgentID {
	return AgentID(atomic.AddUint64(&agent_id, 1))
}