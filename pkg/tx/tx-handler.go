package tx

import (
	"fmt"
)

type TxHandler struct {
	state        map[string]string
	transactions stack
}

func NewTxHandler() *TxHandler {
	return &TxHandler{
		state:        make(map[string]string),
		transactions: stack{},
	}
}

func (t *TxHandler) HandleSet(varName string, newValue string) {
	t.handleOp(varName, newValue)
}

func (t *TxHandler) HandleUnSet(varName string) {
	// Pass empty string, this will remove the variable from state
	t.handleOp(varName, "")
}

func (t *TxHandler) handleOp(varName string, newValue string) {
	op := newChangeOperation(t.state, varName, newValue)
	elem, hasTx := t.transactions.pop()
	/* Check if there is a transaction in progress, doing a pop in the stack.
	 * If yes, I add the operation to the tx, and I add back the tx to the stack.
	 */
	if hasTx {
		tx := elem.(*transaction)
		tx.addOperation(op)
		t.transactions.push(tx)
	}
	op.execute()
}

func (a *TxHandler) HandleGet(varName string) {
	fmt.Println(a.state[varName])
}

func (t *TxHandler) HandleBegin() {
	tx := newTransaction()
	t.transactions.push(tx)
}

func (t *TxHandler) HandleRollback() {
	elem, hasTx := t.transactions.pop()
	if !hasTx {
		fmt.Println("First you need to BEGIN a transaction")
		return
	}
	// To do a rollback, I pop the last transaction.
	tx := elem.(*transaction)
	tx.handleRollback()
}

func (t *TxHandler) HandleCommit() {
	_, hasValue := t.transactions.pop()
	/* Commit only removes the current tx from the stack.
	 * Rollback is not possible anymore in that tx.
	 */
	if !hasValue {
		fmt.Println("First you need to BEGIN a transaction")
		return
	}
}
