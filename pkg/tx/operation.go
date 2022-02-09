package tx

type operation struct {
	state    map[string]string
	varName  string
	newValue string
	oldValue string
}

func newOperation(state map[string]string, varName string, newValue string) *operation {
	return &operation{
		state:    state,
		varName:  varName,
		newValue: newValue,
	}
}

func (o *operation) execute() {
	o.oldValue = o.state[o.varName]
	if o.newValue == "" {
		delete(o.state, o.varName)
	} else {
		o.state[o.varName] = o.newValue
	}
}

func (o *operation) handleRollback() {
	if o.oldValue == "" {
		delete(o.state, o.varName)
	} else {
		o.state[o.varName] = o.oldValue
	}
}
