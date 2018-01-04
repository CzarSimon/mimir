package main

import (
  "crypto/rand"
  "crypto/sha512"
  "encoding/base64"
  "fmt"
  "log"
  "time"

  "github.com/urfave/cli"
)

const (
  MIN_PASSWORD_LEN   = 10
  TOKEN_LENGTH       = 100
  SLEEP_MILLISECONDS = 350
)

// GeneratePassword Command to generate a new password
func GeneratePassword(c *cli.Context) error {
  passphrase := getPassphrase()
  fmt.Printf("Password: %s", createPassword(passphrase))
  return nil
}

// createPassword Creates a password based on a user supplied passphrase
func createPassword(passphrase string) string {
  randomToken := generateRandomToken()
  timestamp := time.Now().UnixNano()
  return hash(passphrase, randomToken, timestamp)
}

// hash Creates a hash of a passwords constituent parts
func hash(passphrase, token string, timestamp int64) string {
  base := fmt.Sprintf("%s-%s-%d", passphrase, token, timestamp)
  return fmt.Sprinf("%x", sha512.Sum512([]byte(base)))
}

// generateRandomToken Creates a random token to be used in a password
func generateRandomToken() string {
  sleep()
  randomBytes := generateRandomBytes()
  return base64.StdEncoding.EncodeToString(randomBytes)
}

// getPassphrase Queries the user for a compliant passphrase
// to be used in password generation
func getPassphrase() string {
  for {
    passphrase := getHiddenInput("Passphrase")
    if len(passphrase) >= MIN_PASSWORD_LEN {
      return passphrase
    }
    fmt.Println("Passphrase to short, min length = %d", MIN_PASSWORD_LEN)
  }
}

// generateRandomBytes Generates an array of random bytes
func generateRandomBytes() []byte {
  bytes := make([]byte, TOKEN_LENGTH)
  rand.Seed(time.Now().UnixNano())
  _, err := rand.Read(bytes)
  if err != nil {
    fmt.Println(err.Error())
    log.Fatal()
  }
  return bytes
}

// sleep Sleeps in order to randomize calls to time.Now().UnixNano()
func sleep() int64 {
  time.Sleep(time.Milliseconds * SLEEP_MILLISECONDS)
}
