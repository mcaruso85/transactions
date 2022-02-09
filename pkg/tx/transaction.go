package tx

type transaction struct {
	operations stack
}

func newTransaction() *transaction {
	return &transaction{
		operations: stack{},
	}
}

func (t *transaction) addOperation(op operation) {
	t.operations.push(op)
}

func (a *transaction) handleRollback() {
	/* Iterate doing a pop in the operations stack of the tx
	 * Execute rollback on every transaction
	 */
	for {
		elem, hasOp := a.operations.pop()
		if !hasOp {
			break
		}
		op := elem.(operation)
		op.handleRollback()
	}
}
