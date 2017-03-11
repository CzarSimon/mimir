package main

import (
  "sync"
)

type ClusterQueue struct {
  Length int64
  Mutex *sync.Mutex
}

type QueueMap map[string]*ClusterQueue

func newQueueMap() QueueMap {
  return make(QueueMap)
}

func (queue QueueMap) addCluster(clusterHash string) {
  if _, present := queue[clusterHash]; present {
    queue[clusterHash].Length++
  } else {
    queue[clusterHash] = &ClusterQueue{
      Length: 0,
      Mutex: &sync.Mutex{},
    }
  }
}

func (queue QueueMap) removeCluster(clusterHash string) {
  if queue[clusterHash].Length == 1 {
    delete(queue, clusterHash)
  } else {
    queue[clusterHash].Length--
  }
}

func (queue QueueMap) lockCluster(clusterHash string) {
  queue[clusterHash].Mutex.Lock()
}

func (queue QueueMap) unlockCluster(clusterHash string) {
  queue[clusterHash].Mutex.Unlock()
}
