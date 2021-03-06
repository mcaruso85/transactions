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
			} else {
				txHandler.HandleSet(args[1], args[2])
			}
		case "UNSET":
			if len(args) != 2 {
				fmt.Println("Wrong args for UNSET. Use UNSET varName")
			} else {
				txHandler.HandleUnSet(args[1])
			}
		case "GET":
			if len(args) != 2 {
				fmt.Println("Wrong args for GET. Use GET varName")
			} else {
				txHandler.HandleGet(args[1])
			}
		case "BEGIN":
			txHandler.HandleBegin()
		case "ROLLBACK":
			txHandler.HandleRollback()
		case "COMMIT":
			txHandler.HandleCommit()
		case "EXIT":
			os.Exit(0)
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
