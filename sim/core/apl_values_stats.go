package core

import (
	"github.com/isfir/wowsims-turtle/sim/core/proto"
	"github.com/isfir/wowsims-turtle/sim/core/stats"
)

type APLValueCurrentAttackPower struct {
	DefaultAPLValueImpl
	unit *Unit
}

func (rot *APLRotation) newValueCurrentAttackPower(_ *proto.APLValueCurrentAttackPower) APLValue {
	return &APLValueCurrentAttackPower{
		unit: rot.unit,
	}
}
func (value *APLValueCurrentAttackPower) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeFloat
}
func (value *APLValueCurrentAttackPower) GetFloat(_ *Simulation) float64 {
	return value.unit.GetStat(stats.AttackPower)
}
func (value *APLValueCurrentAttackPower) String() string {
	return "Current Attack Power"
}

type APLValueCurrentCastSpeedMultiplier struct {
	DefaultAPLValueImpl
	unit *Unit
}

func (rot *APLRotation) newValueCurrentCastSpeedMultiplier(_ *proto.APLValueCurrentCastSpeedMultiplier) APLValue {
	return &APLValueCurrentCastSpeedMultiplier{
		unit: rot.unit,
	}
}
func (value *APLValueCurrentCastSpeedMultiplier) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeFloat
}
func (value *APLValueCurrentCastSpeedMultiplier) GetFloat(_ *Simulation) float64 {
	return value.unit.PseudoStats.CastSpeedMultiplier
}
func (value *APLValueCurrentCastSpeedMultiplier) String() string {
	return "Current Cast Speed Multiplier"
}
