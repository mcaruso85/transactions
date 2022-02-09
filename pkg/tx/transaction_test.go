package tx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOperation(t *testing.T) {
	mockVarName := "varName"
	mockNewValue := "newValue"
	mockState := make(map[string]string)
	op1 := newChangeOperation(mockState, mockVarName, mockNewValue)
	op2 := newChangeOperation(mockState, mockVarName, mockNewValue)

	tx := newTransaction()

	tx.addOperation(op1)
	tx.addOperation(op2)

	assert.Equal(t, 2, tx.operations.len(), "new value do not match")
}

func TestTransactionHandleRollback(t *testing.T) {

	mockOperation1 := new(mockOperation)
	mockOperation2 := new(mockOperation)

	tx := newTransaction()
	tx.addOperation(mockOperation1)
	tx.addOperation(mockOperation2)

	mockOperation1.On("handleRollback").Return()
	mockOperation2.On("handleRollback").Return()

	tx.handleRollback()

	mockOperation1.AssertExpectations(t)
	mockOperation2.AssertExpectations(t)

	mockOperation1.AssertNumberOfCalls(t, "handleRollback", 1)
	mockOperation2.AssertNumberOfCalls(t, "handleRollback", 1)
}
