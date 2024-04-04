package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/ssh"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not read .env file %v", err)
	}

	host := os.Getenv("HOSTNAME")
	port := 22
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		log.Fatalf("Failed to connect to SSH: %v", err)
	}

	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout

	err = session.Run("pihole -up")
	if err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
