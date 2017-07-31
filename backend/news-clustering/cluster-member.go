package main

// ClusterMember represent an article in a cluester and its score
type ClusterMember struct {
	ClusterHash 	string
	UrlHash     	string
	ReferenceScore 	float64
	SubjectScore 	float64
}

// CalcCompoundScore Sums an articles Reference and Subject scorea and returns the result
func CalcCompoundScore(member ClusterMember) float64 {
	return member.ReferenceScore + member.SubjectScore
}
