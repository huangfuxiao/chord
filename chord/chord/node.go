/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Node struct definition and methods defined on it including
 *  initializing a node, joining a node to the Chord ring, and functions
 *  for a node to make calls to other nodes in the Chord ring.
 */

package chord

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/brown-csci1380/s17-bli12-hqian1/cs138"
	"google.golang.org/grpc"
)

// Non-local node representation
type RemoteNode struct {
	Id   []byte
	Addr string
}

// Local node representation
type Node struct {
	Id         []byte           /* Unique Node ID */
	Listener   *net.TCPListener /* Node listener socket */
	Server     *grpc.Server     /* RPC Server */
	Addr       string           /* String of listener address */
	RemoteSelf *RemoteNode      /* Remote node of our self */

	Successor   *RemoteNode  /* This Node's successor */
	Predecessor *RemoteNode  /* This Node's predecessor */
	sLock       sync.RWMutex /* RWLock for successor */
	pLock       sync.RWMutex /* RWLock for predecessor */

	IsShutdown bool         /* Is node in process of shutting down? */
	sdLock     sync.RWMutex /* RWLock for shutdown flag */

	FingerTable []FingerEntry /* Finger table entries */
	FtLock      sync.RWMutex  /* RWLock for finger table */

	dataStore map[string]string /* Local datastore for this node */
	DsLock    sync.RWMutex      /* RWLock for datastore */

	wg sync.WaitGroup /* WaitGroup of concurrent goroutines to sync before exiting */
}

// Initailize a Chord node, start listener, RPC server, and go routines.
func (node *Node) init(parent *RemoteNode, definedId []byte) error {
	if KEY_LENGTH > 128 || KEY_LENGTH%8 != 0 {
		log.Fatal(fmt.Sprintf("KEY_LENGTH of %v is not supported! Must be <= 128 and divisible by 8", KEY_LENGTH))
	}

	listener, _, err := cs138.OpenTCPListener()
	if err != nil {
		return err
	}

	node.Id = HashKey(listener.Addr().String())
	if definedId != nil {
		node.Id = definedId
	}

	node.Listener = listener
	node.Addr = listener.Addr().String()
	node.IsShutdown = false
	node.dataStore = make(map[string]string)

	// Populate RemoteNode that points to self
	node.RemoteSelf = new(RemoteNode)
	node.RemoteSelf.Id = node.Id
	node.RemoteSelf.Addr = node.Addr

	// Join this node to the same chord ring as parent
	err = node.join(parent)
	if err != nil {
		return err
	}

	// Populate finger table
	node.initFingerTable()

	// Thread 1: start RPC server on this connection

	// rpc.RegisterName(node.Addr, node)
	// node.spawn(func() { node.startRpcServer() })
	node.Server = grpc.NewServer()
	RegisterChordRPCServer(node.Server, node)
	go node.Server.Serve(node.Listener)

	// Thread 2: kick off timer to stabilize periodically

	ticker1 := time.NewTicker(time.Millisecond * 100) //freq
	node.spawn(func() { node.stabilize(ticker1) })

	// Thread 3: kick off timer to fix finger table periodically

	ticker2 := time.NewTicker(time.Millisecond * 90) //freq
	node.spawn(func() { node.fixNextFinger(ticker2) })

	return err
}

// Adds a new goroutine to the WaitGroup, spawns the go routine,
// and removes the goroutine from the WaitGroup on exit.
func (node *Node) spawn(fun func()) {
	go func() {
		node.wg.Add(1)
		fun()
		node.wg.Done()
	}()
}

// This node is trying to join an existing ring that a remote node is a part of (i.e., other)
func (node *Node) join(other *RemoteNode) error {
	if other != nil {
		node.Predecessor = nil
		successor, err := other.FindSuccessorRPC(node.Id)
		if err != nil {
			return err
		}
		node.Successor = successor
	} else {
		node.Predecessor = nil
		node.Successor = node.RemoteSelf
	}
	return nil
}

// Thread 2: Psuedocode from figure 7 of chord paper
func (node *Node) stabilize(ticker *time.Ticker) {
	for _ = range ticker.C {
		node.sdLock.RLock()
		sd := node.IsShutdown
		node.sdLock.RUnlock()
		if sd {
			Debug.Printf("[%v-stabilize] Shutting down stabilize timer\n", HashStr(node.Id))
			ticker.Stop()
			return
		}

		node.sLock.RLock()
		succ := node.Successor
		node.sLock.RUnlock()

		x, err := succ.GetPredecessorIdRPC()
		if err != nil {
			Debug.Println("stabilize fail!")
			return
		}
		if x != nil && Between(x.Id, node.Id, succ.Id) {
			node.sLock.Lock()
			node.Successor = x
			succ = node.Successor
			node.sLock.Unlock()
		}
		//Debug.Println("before notify,Successor:", node.Successor.Id)
		succ.NotifyRPC(node.RemoteSelf)
	}
}

// Psuedocode from figure 7 of chord paper
func (node *Node) notify(remoteNode *RemoteNode) {

	if EqualIds(node.Id, remoteNode.Id) {
		return
	}
	node.pLock.RLock()
	pred := node.Predecessor
	node.pLock.RUnlock()
	if pred == nil || Between(remoteNode.Id, pred.Id, node.Id) {
		node.pLock.Lock()
		node.Predecessor = remoteNode
		node.pLock.Unlock()
		if pred == nil {
			node.RemoteSelf.TransferKeysRPC(remoteNode, node.Id)
		} else {
			node.RemoteSelf.TransferKeysRPC(remoteNode, pred.Id)
		}
	}

}

// Psuedocode from figure 4 of chord paper
func (node *Node) findSuccessor(id []byte) (*RemoteNode, error) {

	preNode, err := node.findPredecessor(id)
	if err != nil {
		return &RemoteNode{}, err
	}
	//Debug.Printf("preNode:%v\n", HashStr(preNode.Id))
	remoteNode, err := preNode.GetSuccessorIdRPC()
	if err != nil {
		return &RemoteNode{}, err
	}
	//Debug.Printf("sucNode:%v\n", HashStr(remoteNode.Id))
	return remoteNode, nil
}

// Psuedocode from figure 4 of chord paper
func (node *Node) findPredecessor(id []byte) (*RemoteNode, error) {

	newRemoteNode := node.RemoteSelf
	successor, err := newRemoteNode.GetSuccessorIdRPC()
	//Debug.Printf("succ id:%v\n", HashStr(successor.Id))
	if err != nil {
		return &RemoteNode{}, err
	}
	for !BetweenRightIncl(id, newRemoteNode.Id, successor.Id) {
		newRemoteNode, err = newRemoteNode.ClosestPrecedingFingerRPC(id)
		successor, err = newRemoteNode.GetSuccessorIdRPC()
		if err != nil {
			return &RemoteNode{}, err
		}
	}
	//Debug.Printf("id:%v nodeid:%v\n", HashStr(id), HashStr(newRemoteNode.Id))
	return newRemoteNode, nil
}
