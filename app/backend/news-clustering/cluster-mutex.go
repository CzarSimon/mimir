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

// newQueueMap Returns a new, empty queue map
func newQueueMap() QueueMap {
	return make(QueueMap)
}

// AddAndLockCluster Add a cluster to the queue map and locks the cluster mutex
func (env *Env) AddAndLockCluster(clusterHash string) {
	env.queue.addCluster(clusterHash)
	env.queue.lockCluster(clusterHash)
}

// RemoveAndUnlockCluster Removes a cluster from the queue map an unlocks the cluster mutex
func (env *Env) RemoveAndUnlockCluster(clusterHash string) {
	env.queue.unlockCluster(clusterHash)
	env.queue.removeCluster(clusterHash)
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
