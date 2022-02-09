package tx

type transaction struct {
	operations stack
}

func newTransaction() *transaction {
	return &transaction{
		operations: stack{},
	}
}

func (t *transaction) addOperation(op *operation) {
	t.operations.Push(op)
}

func (a *transaction) handleRollback() {
	for {
		elem, hasOp := a.operations.Pop()
		if !hasOp {
			break
		}
		op := elem.(*operation)
		op.handleRollback()
	}
}
