package main

import (
	"fmt"
)

var GlobalStore = make(map[string]string)

type Transaction struct {
	store map[string]string
	next  *Transaction
}

type TransactionStack struct {
	top  *Transaction
	size int
}

type RollbackStack struct {
	top  *Transaction
	size int
}

/*PushTransaction creates a new active transaction*/
func (ts *TransactionStack) PushTransaction() {
	temp := Transaction{store: make(map[string]string)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

/*PopTransaction deletes a transaction from stack*/
func (ts *TransactionStack) PopTransaction() {
	if ts.top == nil {
		fmt.Println("ERROR: Stack is empty")
	} else {
		ts.top = ts.top.next
		ts.size--
	}
}

/*Peek returns the active transaction*/
func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}

// /*
// Commit write(SET) changes to the store with TranscationStack scope
// Also write changes to disk/file, if data needs to persist after the shell closes
// */
// func (ts *TransactionStack) Commit() {
// 	activeTransaction := ts.Peek()
// 	if activeTransaction != nil {
// 		for key, value := range activeTransaction.store {
// 			GlobalStore[key] = value
// 			if activeTransaction.next != nil {
// 				activeTransaction.next.store[key] = value
// 			}
// 		}
// 	} else {
// 		fmt.Printf("INFO: Nothing to commit\n")
// 	}
// }

// /*RollBackTransaction clears all keys SET within a transaction*/
// func (ts *TransactionStack) RollbackTransaction() {
// 	if ts.top == nil {
// 		fmt.Printf("ERROR: No Active Transaction\n")
// 	} else {
// 		for key := range ts.top.store {
// 			delete(ts.top.store, key)
// 		}
// 	}
// }

// /*Get value of key from Store*/
// func Get(key string, ts *TransactionStack) {
// 	activeTransaction := ts.Peek()
// 	if activeTransaction == nil {
// 		if val, present := GlobalStore[key]; present {
// 			fmt.Printf("%s\n", val)
// 		} else {
// 			fmt.Printf("%s not set\n", key)
// 		}
// 	} else {
// 		if val, present := activeTransaction.store[key]; present {
// 			fmt.Printf("%s\n", val)
// 		} else if val, present := GlobalStore[key]; present {
// 			fmt.Printf("%s\n", val)
// 		} else {
// 			fmt.Printf("%s not set\n", key)
// 		}
// 	}
// }

// /*Set key to value */
// func Set(key, value string, ts *TransactionStack) {
// 	ActiveTransaction := ts.Peek()
// 	if ActiveTransaction == nil {
// 		GlobalStore[key] = value
// 	} else {
// 		ActiveTransaction.store[key] = value
// 	}
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	items := &TransactionStack{}
// 	for {
// 		fmt.Printf("> ")
// 		text, _ := reader.ReadString('\n')
// 		// split the text into operation strings
// 		operation := strings.Fields(text)
// 		switch operation[0] {
// 		case "BEGIN":
// 			items.PushTransaction()
// 		case "ROLLBACK":
// 			items.RollbackTransaction()
// 		case "COMMIT":
// 			items.Commit()
// 			items.PopTransaction()
// 		case "END":
// 			items.PopTransaction()
// 		case "SET":
// 			Set(operation[1], operation[2], items)
// 		case "GET":
// 			Get(operation[1], items)
// 		// case "DELETE":
// 		// 	Delete(operation[1], items)
// 		// case "COUNT":
// 		// 	Count(operation[1], items)
// 		case "STOP":
// 			os.Exit(0)
// 		default:
// 			fmt.Printf("ERROR: Unrecognised Operation %s\n", operation[0])
// 		}
// 	}
// }
