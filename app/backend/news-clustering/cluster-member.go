package main

import (
	"database/sql"
)

// ClusterMember represent an article in a cluester and its score
type ClusterMember struct {
	ClusterHash    string
	URLHash        string
	ReferenceScore float64
	SubjectScore   float64
}

// CalcCompoundScore Sums an articles Reference and Subject scorea and returns the result
func CalcCompoundScore(member ClusterMember) float64 {
	return member.ReferenceScore + member.SubjectScore
}

// GetMembers Retrives a list of members other than the supplied member
func GetMembers(tx *sql.Tx, member ClusterMember) ([]ClusterMember, error) {
	members := []ClusterMember{member}
	query := `SELECT CLUSTER_HASH, URL_HASH, REFERENCE_SCORE, SUBJECT_SCORE
						FROM CLUSTER_MEMBER WHERE CLUSTER_HASH=$1 AND URL_HASH!=$2`
	rows, err := tx.Query(query, member.ClusterHash, member.URLHash)
	defer rows.Close()
	if err != nil {
		return members, err
	}
	var m ClusterMember
	for rows.Next() {
		err = rows.Scan(&m.ClusterHash, &m.URLHash, &m.ReferenceScore, &m.SubjectScore)
		if err != nil {
			return members, err
		}
		members = append(members, m)
	}
	return members, nil
}

// StoreMember Stores new or updates member if existing
func StoreMember(tx *sql.Tx, member ClusterMember) error {
	existingMember, err := memberExists(tx, member)
	if err != nil {
		return err
	}
	if existingMember {
		err = UpdateMember(tx, member)
		if err != nil {
			return err
		}
	} else {
		err = StoreNewMember(tx, member)
		if err != nil {
			return err
		}
	}
	return nil
}

// UpdateMember Updates an existing member
func UpdateMember(tx *sql.Tx, member ClusterMember) error {
	query := `UPDATE CLUSTER_MEMBER SET REFERENCE_SCORE=$1, SUBJECT_SCORE=$2
						WHERE CLUSTER_HASH=$3 AND URL_HASH=$4`
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(member.ReferenceScore, member.SubjectScore, member.ClusterHash, member.URLHash)
	if err != nil {
		return err
	}
	return nil
}

// StoreNewMember Inserts a new member in the database
func StoreNewMember(tx *sql.Tx, member ClusterMember) error {
	query := `INSERT INTO CLUSTER_MEMBER
						(CLUSTER_HASH, URL_HASH, REFERENCE_SCORE, SUBJECT_SCORE)
						VALUES($1,$2,$3,$4)`
	stmt, err := tx.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(member.ClusterHash, member.URLHash, member.ReferenceScore, member.SubjectScore)
	if err != nil {
		return err
	}
	return nil
}

// memberExists Checks if member is already stored
func memberExists(tx *sql.Tx, member ClusterMember) (bool, error) {
	query := "SELECT URL_HASH FROM CLUSTER_MEMBER WHERE CLUSTER_HASH=$1 AND URL_HASH=$2"
	var urlHash string
	err := tx.QueryRow(query, member.ClusterHash, member.URLHash).Scan(urlHash)
	if err == nil {
		return true, nil
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	return false, err
}
