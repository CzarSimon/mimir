package main

import (
  "time"
)

// Cluster is the object representing the score and members of a cluster
type Cluster struct {
  ClusterHash string
	Title       string
  Ticker      string
  Date        time.Time
	Leader      string         
	Score       float64
  Members     []ClusterMember
}
