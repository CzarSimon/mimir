package main

// ClusterMember represent an article in a cluester and it score
type ClusterMember struct {
	ClusterHash 	string
	UrlHash     	string
	ReferenceScore 	float64
	SubjectScore 	float64
}
