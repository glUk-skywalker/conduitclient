package objects

import (
	"strconv"

	"github.com/gluk-skywalker/conduitclient/urlvalues"
)

// URLParamAppendable is the interface for structs that can be appended as url.Params
type URLParamAppendable interface {
	AppendTo(params *urlvalues.URLValues, transactionPrefix string)
}

// ColumnTransaction is the structure for status transactions
type ColumnTransaction struct {
	Value []string
}

// AppendTo appends itself to the passed urlvalues.URLValues with the passed prefix
func (t ColumnTransaction) AppendTo(params *urlvalues.URLValues, transactionPrefix string) {
	params.Add(transactionPrefix+"[type]", "column")
	for i, v := range t.Value {
		params.Add(transactionPrefix+"[value]["+strconv.Itoa(i)+"]", v)
	}
}

// StatusTransaction is the structure for status transactions
type StatusTransaction struct {
	Value string
}

// AppendTo appends itself to the passed urlvalues.URLValues with the passed prefix
func (t StatusTransaction) AppendTo(params *urlvalues.URLValues, transactionPrefix string) {
	params.Add(transactionPrefix+"[type]", "status")
	params.Add(transactionPrefix+"[value]", t.Value)
}

// AddTask is the structure for add transactions
type AddTask struct {
	Value []string
}

// AppendTo appends itself to the passed urlvalues.URLValues with the passed prefix
func (t AddTask) AppendTo(params *urlvalues.URLValues, transactionPrefix string) {
	params.Add(transactionPrefix+"[type]", "tasks.add")
	for i, v := range t.Value {
		params.Add(transactionPrefix+"[value]["+strconv.Itoa(i)+"]", v)
	}
}
