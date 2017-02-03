/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: API and interal functions to interact with the Key-Value store
 *  that the Chord ring is providing.
 */

package chord

import (
	"fmt"
	"log"
)

/*
 * External API Into Datastore
 */

// Get a value in the datastore, given an abitrary node in the ring.
func Get(node *Node, key string) (string, error) {

	//TODO students should implement this method
	remoteNode, err := node.locate(key)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return remoteNode.GetRPC(key)
}

// Put a key/value in the datastore, given an abitrary node in the ring.
func Put(node *Node, key string, value string) error {

	//TODO students should implement this method
	remoteNode, err := node.locate(key)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//Debug.Println("key:%s locate")
	return remoteNode.PutRPC(key, value)
}

// Internal helper method to find the appropriate node in the ring based on a key.
func (node *Node) locate(key string) (*RemoteNode, error) {

	//TODO students should implement this method
	id := HashKey(key)
	//Debug.Printf("key:%s hash:%v\n", key, HashStr(id))
	return node.findSuccessor(id)
}

/*
 * RPCs to assist with interfacing with the datastore ring
 */

func (node *Node) GetLocal(req *KeyValueReq) (*KeyValueReply, error) {
	if err := validateRpc(node, req.NodeId); err != nil {
		return nil, err
	}

	//TODO students should implement this method
	node.DsLock.RLock()
	defer node.DsLock.RUnlock()
	key := req.Key
	val := node.dataStore[key]
	reply := KeyValueReply{key, val}

	return &reply, nil
}

func (node *Node) PutLocal(req *KeyValueReq) (*KeyValueReply, error) {
	if err := validateRpc(node, req.NodeId); err != nil {
		return nil, err
	}

	//TODO students should implement this method
	node.DsLock.Lock()
	defer node.DsLock.Unlock()
	key := req.Key
	val := req.Value
	node.dataStore[key] = val
	reply := KeyValueReply{key, val}

	return &reply, nil
}

// Find locally stored keys that are between (predId : fromId].
// Any of these nodes should be moved to fromId.
func (node *Node) TransferKeys(req *TransferReq) (*RpcOkay, error) {
	if err := validateRpc(node, req.NodeId); err != nil {
		return nil, err
	}

	//Debug.Println("Transfer Key")
	node.DsLock.Lock()
	defer node.DsLock.Unlock()
	for key, val := range node.dataStore {
		predId := req.PredId
		// if predId == nil {
		// 	predId = node.Id
		// }
		if BetweenRightIncl(HashKey(key), predId, req.FromId) {
			remNode := RemoteNode{req.FromId, req.FromAddr}
			err := remNode.PutRPC(key, val)
			if err != nil {
				reply := RpcOkay{false}
				return &reply, err
			}
			//Delete the data store locally
			delete(node.dataStore, key)
		}
	}
	reply := RpcOkay{true}

	return &reply, nil
}

// Print the contents of a node's datastore.
func PrintDataStore(node *Node) {
	node.DsLock.RLock()
	defer node.DsLock.RUnlock()
	fmt.Printf("Node %v datastore: %v\n", HashStr(node.Id), node.dataStore)
}

// Returns the contents of a node's datastore as a string.
func DataStoreToString(node *Node) string {
	node.DsLock.RLock()
	defer node.DsLock.RUnlock()
	return fmt.Sprintf("Node %v datastore: %v\n", HashStr(node.Id), node.dataStore)
}
