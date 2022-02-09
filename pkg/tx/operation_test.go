package tx

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	mockVarName := "varName"
	mockNewValue := "newValue"
	mockState := make(map[string]string)

	op := newChangeOperation(mockState, mockVarName, mockNewValue)

	op.execute()

	newValue := mockState[mockVarName]
	assert.Equal(t, mockNewValue, newValue, "new value do not match")
}

func TestHandleRollback(t *testing.T) {
	mockVarName := "varName"
	mockNewValue := "value"
	mockState1 := make(map[string]string)
	mockOldValue := "oldValue"
	mockState1[mockVarName] = mockNewValue

	op1 := newChangeOperation(mockState1, mockVarName, mockNewValue)

	op1.oldValue = mockOldValue
	op1.handleRollback()

	valueAfterRollback := mockState1[mockVarName]
	assert.Equal(t, mockOldValue, valueAfterRollback, "after rollback value do not match")

	mockState2 := make(map[string]string)
	mockOldValue = ""
	mockState2[mockVarName] = mockNewValue
	op2 := newChangeOperation(mockState2, mockVarName, mockNewValue)
	op2.newValue = mockNewValue
	op2.oldValue = mockOldValue
	op2.handleRollback()

	v := reflect.ValueOf(mockState2).MapIndex(reflect.ValueOf(mockVarName))
	isPresent := v != reflect.Value{}
	assert.True(t, !isPresent, "old value is empty and must delete the variable from map")

}
