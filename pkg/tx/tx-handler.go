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
	t.handleOp(varName, "")
}

func (t *TxHandler) handleOp(varName string, newValue string) {
	op := newOperation(t.state, varName, newValue)
	elem, hasTx := t.transactions.Pop()
	if hasTx {
		tx := elem.(*transaction)
		tx.addOperation(op)
		t.transactions.Push(tx)
	}
	op.execute()
}

func (a *TxHandler) HandleGet(varName string) {
	fmt.Println(a.state[varName])
}

func (t *TxHandler) HandleBegin() {
	tx := newTransaction()
	t.transactions.Push(tx)
}

func (t *TxHandler) HandleRollback() {
	elem, hasTx := t.transactions.Pop()
	if !hasTx {
		fmt.Println("First you need to BEGIN a transaction")
		return
	}
	tx := elem.(*transaction)
	tx.handleRollback()
}

func (t *TxHandler) HandleCommit() {
	_, hasValue := t.transactions.Pop()
	if !hasValue {
		fmt.Println("First you need to BEGIN a transaction")
		return
	}
}
