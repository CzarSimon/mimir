package main

import (
	"sync"
)

// ClusterQueue Contains a mutex to prevent concurrent updating of a cluster
type ClusterQueue struct {
	Length int64
	Mutex  *sync.Mutex
}

// QueueMap Map of ClusterQueues identified by the cluster hash as key
type QueueMap map[string]*ClusterQueue

func newQueueMap() QueueMap {
	return make(QueueMap)
}

// addCluster Adds a cluster in the QueueMap
func (queue QueueMap) addCluster(clusterHash string) {
	if _, present := queue[clusterHash]; present {
		queue[clusterHash].Length++
	} else {
		queue[clusterHash] = &ClusterQueue{
			Length: 0,
			Mutex:  &sync.Mutex{},
		}
	}
}

// removeCluster Removes a cluster from the QueueMap
func (queue QueueMap) removeCluster(clusterHash string) {
	if queue[clusterHash].Length == 1 {
		delete(queue, clusterHash)
	} else {
		queue[clusterHash].Length--
	}
}

// lockCluster Locks the mutex of a given cluster
func (queue QueueMap) lockCluster(clusterHash string) {
	queue[clusterHash].Mutex.Lock()
}

// unlockCluster Unlocks the cluster of a given cluster
func (queue QueueMap) unlockCluster(clusterHash string) {
	queue[clusterHash].Mutex.Unlock()
}
