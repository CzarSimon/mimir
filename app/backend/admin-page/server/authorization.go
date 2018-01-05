package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type bodyAuth struct {
	Token string
}

func (env *Env) authenticate(res http.ResponseWriter, req *http.Request) error {
	if parseToken(req) == env.auth.token || env.devMode {
		return nil
	} else {
		err := errors.New("user not authenticated")
		http.Error(res, err.Error(), http.StatusUnauthorized)
		return err
	}
}

func parseToken(req *http.Request) string {
	return req.Header.Get("Authorizaton")
}

type Auth struct {
	User, Pwd string
}

type User struct {
	Username, Token string
}

func (env *Env) login(res http.ResponseWriter, req *http.Request) {
	var auth Auth
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&auth); err != io.EOF {
		checkErrNice(err)
	}
	if env.auth.pwdHash == pwdHash(auth.Pwd, env.auth.salt) || env.devMode {
		js, err := json.Marshal(User{auth.User, env.auth.token})
		checkErrNice(err)
		jsonRes(res, js)
	} else {
		jsonStringRes(res, "incorrect password")
	}
}

func generateToken() string {
	randomBytes := make([]byte, 10)
	rand.Seed(time.Now().UnixNano())
	rand.Read(randomBytes)
	return fmt.Sprintf("%x", sha256.Sum256(randomBytes))
}

func pwdHash(pwd, salt string) string {
	byteHash := sha256.Sum256([]byte(pwd + salt))
	return fmt.Sprintf("%x", byteHash)
}
