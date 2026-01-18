package sim

import (
	_ "github.com/isfir/wowsims-turtle/sim/common"
	"github.com/isfir/wowsims-turtle/sim/druid/balance"
	"github.com/isfir/wowsims-turtle/sim/paladin/retribution"
	dpsrogue "github.com/isfir/wowsims-turtle/sim/rogue/dps_rogue"
	"github.com/isfir/wowsims-turtle/sim/shaman/elemental"
	"github.com/isfir/wowsims-turtle/sim/shaman/enhancement"
	"github.com/isfir/wowsims-turtle/sim/shaman/warden"

	"github.com/isfir/wowsims-turtle/sim/druid/feral"
	// restoDruid "github.com/isfir/wowsims-turtle/sim/druid/restoration"
	// feralTank "github.com/isfir/wowsims-turtle/sim/druid/tank"
	_ "github.com/isfir/wowsims-turtle/sim/encounters"
	"github.com/isfir/wowsims-turtle/sim/hunter"
	"github.com/isfir/wowsims-turtle/sim/mage"

	// holyPaladin "github.com/isfir/wowsims-turtle/sim/paladin/holy"
	"github.com/isfir/wowsims-turtle/sim/paladin/protection"
	// "github.com/isfir/wowsims-turtle/sim/paladin/retribution"
	// healingPriest "github.com/isfir/wowsims-turtle/sim/priest/healing"
	"github.com/isfir/wowsims-turtle/sim/priest/shadow"

	// restoShaman "github.com/isfir/wowsims-turtle/sim/shaman/restoration"
	dpsWarlock "github.com/isfir/wowsims-turtle/sim/warlock/dps"
	dpsWarrior "github.com/isfir/wowsims-turtle/sim/warrior/dps_warrior"
	tankWarrior "github.com/isfir/wowsims-turtle/sim/warrior/tank_warrior"
)

var registered = false

func RegisterAll() {
	if registered {
		return
	}
	registered = true

	balance.RegisterBalanceDruid()
	feral.RegisterFeralDruid()
	// feralTank.RegisterFeralTankDruid()
	// restoDruid.RegisterRestorationDruid()
	elemental.RegisterElementalShaman()
	enhancement.RegisterEnhancementShaman()
	warden.RegisterWardenShaman()
	// restoShaman.RegisterRestorationShaman()
	hunter.RegisterHunter()
	mage.RegisterMage()
	// healingPriest.RegisterHealingPriest()
	shadow.RegisterShadowPriest()
	dpsrogue.RegisterDpsRogue()
	dpsWarrior.RegisterDpsWarrior()
	tankWarrior.RegisterTankWarrior()
	// holyPaladin.RegisterHolyPaladin()
	protection.RegisterProtectionPaladin()
	retribution.RegisterRetributionPaladin()
	dpsWarlock.RegisterDpsWarlock()
}
