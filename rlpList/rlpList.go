package rlpList

type RlpList struct {
	rlpTypes [](interface{})
}

func NewRlpList(rlpTypes [](interface{})) *RlpList {
	return &RlpList{rlpTypes: rlpTypes}
}

func (r *RlpList) GetValue() [](interface{}) {
	return r.rlpTypes
}

func (r *RlpList) SetValue(rlpTypes [](interface{})) [](interface{}) {
	r.rlpTypes = rlpTypes
}
