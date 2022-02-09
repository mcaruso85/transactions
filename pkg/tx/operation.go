package tx

type operation interface {
	execute()
	handleRollback()
}

type changeOperation struct {
	state    map[string]string
	varName  string
	newValue string
	oldValue string
}

func newChangeOperation(state map[string]string, varName string, newValue string) *changeOperation {
	return &changeOperation{
		state:    state,
		varName:  varName,
		newValue: newValue,
	}
}

func (o *changeOperation) execute() {
	o.oldValue = o.state[o.varName]
	if o.newValue == "" {
		delete(o.state, o.varName)
	} else {
		o.state[o.varName] = o.newValue
	}
}

func (o *changeOperation) handleRollback() {
	if o.oldValue == "" {
		delete(o.state, o.varName)
	} else {
		o.state[o.varName] = o.oldValue
	}
}
