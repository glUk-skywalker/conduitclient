package objects

import (
	"net/url"
	"strconv"
)

// URLParamAppendable is the interface for structs that can be appended as url.Params
type URLParamAppendable interface {
	AppendTo(params *url.Values, transactionPrefix string)
}

// ColumnTransaction is the structure for status transactions
type ColumnTransaction struct {
	Value []string
}

// AppendTo appends itself to the passed url.Values with the passed prefix
func (t ColumnTransaction) AppendTo(params *url.Values, transactionPrefix string) {
	params.Add(transactionPrefix+"[type]", "column")
	for i, v := range t.Value {
		params.Add(transactionPrefix+"[value]["+strconv.Itoa(i)+"]", v)
	}
}

// StatusTransaction is the structure for status transactions
type StatusTransaction struct {
	Value string
}

// AppendTo appends itself to the passed url.Values with the passed prefix
func (t StatusTransaction) AppendTo(params *url.Values, transactionPrefix string) {
	params.Add(transactionPrefix+"[type]", "status")
	params.Add(transactionPrefix+"[value]", t.Value)
}
