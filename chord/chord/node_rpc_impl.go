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

	//TODO students should implement this method
	node.sLock.RLock()
	defer node.sLock.RUnlock()
	if node.Successor == nil {
		return &IdReply{nil, "", false}, nil
	} else {
		return &IdReply{node.Successor.Id, node.Successor.Addr, true}, nil
	}
	return nil, nil
}

func (node *Node) SetPredecessorId(req *UpdateReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.FromId); err != nil {
		return nil, err
	}

	//TODO students should implement this method
	node.pLock.RLock()
	defer node.pLock.RUnlock()
	node.Predecessor = &RemoteNode{req.UpdateId, req.UpdateAddr}

	return &RpcOkay{true}, nil
}

func (node *Node) SetSuccessorId(req *UpdateReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.FromId); err != nil {
		return nil, err
	}

	//TODO students should implement this method

	node.sLock.RLock()
	defer node.sLock.RUnlock()
	node.Successor = &RemoteNode{req.UpdateId, req.UpdateAddr}

	return &RpcOkay{true}, nil
}

func (node *Node) Notify(req *NotifyReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.NodeId); err != nil {
		return nil, err
	}

	//TODO students should implement this method

	return nil, nil
}

func (node *Node) FindSuccessor(query *RemoteQuery) (*IdReply, error) {
	if err := validateRpc(node, query.FromId); err != nil {
		return nil, err
	}

	//TODO students should implement this method

	return nil, nil
}

func (node *Node) ClosestPrecedingFinger(query *RemoteQuery) (*IdReply, error) {
	if err := validateRpc(node, query.FromId); err != nil {
		return nil, err
	}

	//TODO students should implement this method
	for i := KEY_LENGTH; i >= 1; i++ {
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
