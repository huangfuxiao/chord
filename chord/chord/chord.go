/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Defines global constants and variables and provides
 *  functionality to create and shutdown nodes.
 */

package chord

import "time"

// Number of bits (i.e. m value). Assumes <= 128 and divisible by 8
const KEY_LENGTH = 16

// Timeout of RPC Calls
const RPC_TIMEOUT = 5000 * time.Millisecond

// Creates a Chord node with a pre-defined ID (useful for testing).
func CreateDefinedNode(parent *RemoteNode, definedId []byte) (*Node, error) {
	node := new(Node)
	err := node.init(parent, definedId)
	if err != nil {
		return nil, err
	}
	return node, err
}

// Create Chord node with random ID based on listener address.
func CreateNode(parent *RemoteNode) (*Node, error) {
	node := new(Node)
	err := node.init(parent, nil)
	if err != nil {
		return nil, err
	}
	return node, err
}

// Gracefully shutdown a specified Chord node.
func ShutdownNode(node *Node) {
	node.sdLock.Lock()
	node.IsShutdown = true
	node.sdLock.Unlock()

	node.sLock.RLock()
	succ := node.Successor
	node.sLock.RUnlock()

	node.pLock.RLock()
	pred := node.Predecessor
	node.pLock.RUnlock()

	//Debug.Printf("pred id:%v succ id:%v", pred.Id, succ.Id)
	if pred != nil {
		node.RemoteSelf.TransferKeysRPC(succ, pred.Id)
		pred.SetSuccessorIdRPC(succ)
		if EqualIds(pred.Id, succ.Id) {
			succ.SetPredecessorIdRPC(nil)
			predid, _ := succ.GetPredecessorIdRPC()
			Debug.Printf("succ id:%v pred id:%v", succ, predid)
		} else {
			succ.SetPredecessorIdRPC(pred)
		}
		//newpred, _ := succ.GetPredecessorIdRPC()
		//if newpred == nil {
		//	Debug.Println("nil")
		//} else {
		//	Debug.Printf("succ id:%v pred id:%v", pred.Id, newpred.Id)
		//}

	}
	// Wait for all go routines to exit.
	node.wg.Wait()
	node.Server.GracefulStop()
	node.Listener.Close()
}
