package main

import "fmt"

// KeyValueStore represents a simple key-value store with transactions.
type KeyValueStore struct {
	data             map[string]string
	transactionStack []map[string]string
}

// NewKeyValueStore initializes a new key-value store.
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data:             make(map[string]string),
		transactionStack: make([]map[string]string, 0),
	}
}

// Set sets the value of a key.
func (kvs *KeyValueStore) Set(key, value string) {
	if len(kvs.transactionStack) > 0 {
		// If inside a transaction, update the transaction's data
		currentTransaction := kvs.transactionStack[len(kvs.transactionStack)-1]
		currentTransaction[key] = value
	} else {
		// If not inside a transaction, update the main data
		kvs.data[key] = value
	}
}

// Get gets the value of a key.
func (kvs *KeyValueStore) Get(key string) (string, bool) {
	// First, check if the key is in the current transaction
	if len(kvs.transactionStack) > 0 {
		for i := len(kvs.transactionStack) - 1; i >= 0; i-- {
			currentTransaction := kvs.transactionStack[i]
			if val, exists := currentTransaction[key]; exists {
				return val, true
			}
		}
	}

	// If not in any transaction, check the main data
	val, exists := kvs.data[key]
	return val, exists
}

// Delete deletes a key.
func (kvs *KeyValueStore) Delete(key string) {
	delete(kvs.data, key)
}

// Begin starts a new transaction.
func (kvs *KeyValueStore) Begin() {
	// Start a new transaction
	kvs.transactionStack = append(kvs.transactionStack, make(map[string]string))
}

// Commit commits the current transaction.
func (kvs *KeyValueStore) Commit() {
	if len(kvs.transactionStack) > 0 {
		// Merge the transaction data with the main data
		currentTransaction := kvs.transactionStack[len(kvs.transactionStack)-1]
		for key, value := range currentTransaction {
			kvs.data[key] = value
		}

		// Pop the transaction from the stack
		kvs.transactionStack = kvs.transactionStack[:len(kvs.transactionStack)-1]
	}
}

// Rollback rolls back the current transaction.
func (kvs *KeyValueStore) Rollback() {
	if len(kvs.transactionStack) > 0 {
		// Discard the current transaction
		kvs.transactionStack = kvs.transactionStack[:len(kvs.transactionStack)-1]
	}
}

func main() {
	// Example Usage:
	kvs := NewKeyValueStore()

	kvs.Begin()
	kvs.Set("a", "26")
	val, exists := kvs.Get("a")
	fmt.Printf("Value of a: %s, Exists: %t\n", val, exists) // Output: Value of a: 26, Exists: true

	kvs.Begin()
	kvs.Set("a", "27")
	kvs.Commit()

	val, exists = kvs.Get("a")
	fmt.Printf("Value of a: %s, Exists: %t\n", val, exists) // Output: Value of a: 27, Exists: true

	kvs.Rollback()

	val, exists = kvs.Get("a")
	fmt.Printf("Value of a: %s, Exists: %t\n", val, exists) // Output: Value of a: 26, Exists: true
}
