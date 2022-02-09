package tx

import (
	"github.com/stretchr/testify/mock"
)

type mockOperation struct {
	mock.Mock
}

func (m *mockOperation) handleRollback() {
	m.Called()
}

func (m *mockOperation) execute() {
	m.Called()
}
