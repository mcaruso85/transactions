package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mcaruso85/transactions/pkg/tx"
)

func main() {
	txHandler := tx.NewTxHandler()
	for {
		args := getInputArgs()
		switch args[0] {
		case "SET":
			if len(args) != 3 {
				fmt.Println("Wrong args for SET. Use SET varName varValue")
				return
			}
			txHandler.HandleSet(args[1], args[2])
		case "UNSET":
			if len(args) != 2 {
				fmt.Println("Wrong args for UNSET. Use UNSET varName")
				return
			}
			txHandler.HandleUnSet(args[1])
		case "GET":
			if len(args) != 2 {
				fmt.Println("Wrong args for GET. Use GET varName")
				return
			}
			txHandler.HandleGet(args[1])
		case "BEGIN":
			txHandler.HandleBegin()
		case "ROLLBACK":
			txHandler.HandleRollback()
		case "COMMIT":
			txHandler.HandleCommit()
		default:
			fmt.Println("Wrong command. Please enter a valid command: SET, UNSET, GET, BEGIN, ROLLBACK, COMMIT")
		}
	}
}

func getInputArgs() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return strings.Split(line, " ")
}
