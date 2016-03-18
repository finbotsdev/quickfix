package undinstrmtcollgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
)

//NoUnderlyings is a repeating group in UndInstrmtCollGrp
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//CollAction is a non-required field for NoUnderlyings.
	CollAction *int `fix:"944"`
}

//UndInstrmtCollGrp is a fix50sp2 Component
type UndInstrmtCollGrp struct {
	//NoUnderlyings is a non-required field for UndInstrmtCollGrp.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
}

func (m *UndInstrmtCollGrp) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }