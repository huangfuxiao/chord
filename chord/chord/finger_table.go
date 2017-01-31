/*
 *  Brown University, CS138, Spring 2017
 *
 *  Purpose: Finger table related functions for a given Chord node.
 */

package chord

import (
	"bytes"
	"fmt"
	"math/big"
	"time"
)

var next = 1

// A single finger table entry.
type FingerEntry struct {
	Start []byte      /* ID hash of (n + 2^i) mod (2^m)  */
	Node  *RemoteNode /* RemoteNode that Start points to */
}

// Create initial finger table that only points to itself (will be fixed later).
func (node *Node) initFingerTable() {

	node.FingerTable = make([]FingerEntry, 0)
	for i := 0; i <= KEY_LENGTH; i++ {
		node.FingerTable = append(node.FingerTable, FingerEntry{node.Id, node.RemoteSelf})
	}

}

// Called periodically (in a separate go routine) to fix entries in our finger table.
func (node *Node) fixNextFinger(ticker *time.Ticker) {

	for _ = range ticker.C {
		node.sdLock.RLock()
		sd := node.IsShutdown
		node.sdLock.RUnlock()
		if sd {
			Debug.Printf("[%v] Shutting down fixNextFinger timer\n", HashStr(node.Id))
			ticker.Stop()
			return
		}

		//		next_hash := fingerMath(node.Id, next, KEY_LENGTH)
		//		//node.FingerTable[next] = node.findSuccessor(next_hash)
		//		next += 1
		//		if next > KEY_LENGTH-1 {
		//			next = 1
		//		}
	}
}

// Calculates: (n + 2^i) mod (2^m).
func fingerMath(n []byte, i int, m int) []byte {

	nInt := big.Int{}
	nInt.SetBytes(n)
	aInt := big.Int{}
	aInt.SetInt64(int64(2 ^ i))
	bInt := big.Int{}
	bInt.SetInt64(int64(2 ^ m))

	sum := big.Int{}
	sum.Add(&nInt, &aInt)
	sum.Mod(&sum, &bInt)

	return sum.Bytes()

}

// Print contents of a node's finger table.
func PrintFingerTable(node *Node) {
	node.FtLock.RLock()
	defer node.FtLock.RUnlock()

	fmt.Printf("[%v] FingerTable:\n", HashStr(node.Id))

	for _, val := range node.FingerTable {
		fmt.Printf("\t{start:%v\tnodeLoc:[%v] %v}\n",
			HashStr(val.Start), HashStr(val.Node.Id), val.Node.Addr)
	}
}

// Returns contents of a node's finger table as a string.
func FingerTableToString(node *Node) string {
	node.FtLock.RLock()
	defer node.FtLock.RUnlock()

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("[%v] FingerTable:\n", HashStr(node.Id)))

	for _, val := range node.FingerTable {
		buffer.WriteString(fmt.Sprintf("\t{start:%v\tnodeLoc:[%v] %v}\n",
			HashStr(val.Start), HashStr(val.Node.Id), val.Node.Addr))
	}

	return buffer.String()
}
