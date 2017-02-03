/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Implementation of RPC_API functions, these are the functions
 *  that actually get executed locally on a destination Chord node when
 *  an RPC() function is called.
 */

package chord

import (
	"bytes"
	"errors"
	"fmt"
)

// Validate that we're executing this RPC on the intended node.
func validateRpc(node *Node, reqId []byte) error {
	if !bytes.Equal(node.Id, reqId) {
		errStr := fmt.Sprintf("Node ids do not match %v, %v", node.Id, reqId)
		return errors.New(errStr)
	}
	return nil
}

func (node *Node) GetPredecessorId(req *RemoteId) (*IdReply, error) {
	if err := validateRpc(node, req.Id); err != nil {
		return nil, err
	}
	// Predecessor may be nil, which is okay.
	node.pLock.RLock()
	defer node.pLock.RUnlock()
	if node.Predecessor == nil {
		return &IdReply{nil, "", false}, nil
	} else {
		return &IdReply{node.Predecessor.Id, node.Predecessor.Addr, true}, nil
	}
}

func (node *Node) GetSuccessorId(req *RemoteId) (*IdReply, error) {
	if err := validateRpc(node, req.Id); err != nil {
		return nil, err
	}

	node.sLock.RLock()
	defer node.sLock.RUnlock()
	return &IdReply{node.Successor.Id, node.Successor.Addr, true}, nil
}

func (node *Node) SetPredecessorId(req *UpdateReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.FromId); err != nil {
		return nil, err
	}

	node.pLock.Lock()
	defer node.pLock.Unlock()
	node.Predecessor = &RemoteNode{req.UpdateId, req.UpdateAddr}

	return &RpcOkay{true}, nil
}

func (node *Node) SetSuccessorId(req *UpdateReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.FromId); err != nil {
		return nil, err
	}

	node.sLock.Lock()
	defer node.sLock.Unlock()
	node.Successor = &RemoteNode{req.UpdateId, req.UpdateAddr}

	return &RpcOkay{true}, nil
}

func (node *Node) Notify(req *NotifyReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.NodeId); err != nil {
		return nil, err
	}
	node.notify(&RemoteNode{req.UpdateId, req.UpdateAddr})

	return &RpcOkay{true}, nil
}

func (node *Node) FindSuccessor(query *RemoteQuery) (*IdReply, error) {
	if err := validateRpc(node, query.FromId); err != nil {
		return nil, err
	}

	node.sLock.RLock()
	defer node.sLock.RUnlock()
	successor, err := node.findSuccessor(query.Id)
	if err != nil {
		return &IdReply{}, err
	}
	return &IdReply{successor.Id, successor.Addr, true}, nil
}

func (node *Node) ClosestPrecedingFinger(query *RemoteQuery) (*IdReply, error) {
	if err := validateRpc(node, query.FromId); err != nil {
		return nil, err
	}
	node.FtLock.RLock()
	defer node.FtLock.RUnlock()
	for i := KEY_LENGTH - 1; i >= 0; i = i - 1 {
		if Between(node.FingerTable[i].Node.Id, node.Id, query.Id) {
			return &IdReply{node.FingerTable[i].Node.Id, node.FingerTable[i].Node.Addr, true}, nil
		}
	}

	return &IdReply{node.Id, node.Addr, true}, nil
}

// Check if node is alive
func nodeIsAlive(remoteNode *RemoteNode) bool {
	_, err := remoteNode.GetSuccessorIdRPC()
	return err == nil
}
