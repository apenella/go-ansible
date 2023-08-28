package enumtipe

import "database/sql/driver"

type CostomBool string

const (
	CostomBoolTrue  CostomBool = "true"
	CostomBoolFalse CostomBool = "false"
)

var bool_values = []string{
	CostomBoolTrue.String(),
	CostomBoolFalse.String(),
}

func (CostomBool) Values() (cbs []string) {
	cbs = append(cbs, bool_values...)
	return
}

func (cb CostomBool) Value() (driver.Value, error) {
	return cb.String(), nil
}

func (cb CostomBool) String() string {
	return string(cb)
}
