package beyond

import (
	"strconv"
	"time"

	"imetatroll.com/character.git/lib/base"
	"imetatroll.com/character.git/lib/dnd"
)

func (char *Character) Transfer(userID string) *base.Character {
	now := time.Now().Unix()

	target := dnd.NewCharacter(userID)

	if char.Data.Inspiration {
		target.Top.Set("Inspiration", "true", now)
	} else {
		target.Top.Set("Inspiration", "false", now)
	}

	// top
	// 2do: class dice/count
	// 2do: save modifiers - if i remember correctly these are actually calculated values
	target.Top.Set("CharacterName", char.Data.Name, now)
	target.Top.Set("Race", char.Data.Race.FullName, now) // or BaseName
	target.Top.Set("Background", char.Data.Background.Definition.Name, now)
	target.Top.Set("Class", char.Data.Classes[0].Definition.Name, now)
	target.Top.Set("Alignment", char.GetAlignment(), now)
	target.Top.Set("XP", strconv.Itoa(char.Data.CurrentXp), now)
	target.Top.Set("Level", strconv.Itoa(char.Data.Classes[0].Level), now)

	// abilities
	target.Top.Set("Strength", char.GetAbility("Strength"), now)
	target.Top.Set("Dexterity", char.GetAbility("Dexterity"), now)
	target.Top.Set("Constitution", char.GetAbility("Constitution"), now)
	target.Top.Set("Intelligence", char.GetAbility("Intelligence"), now)
	target.Top.Set("Wisdom", char.GetAbility("Wisdom"), now)
	target.Top.Set("Charisma", char.GetAbility("Charisma"), now)

	// saving throws
	target.Top.Set("StrengthSaveCheck", char.GetClassSaveProficiency("Strength"), now)
	target.Top.Set("DexteritySaveCheck", char.GetClassSaveProficiency("Dexterity"), now)
	target.Top.Set("ConstitutionSaveCheck", char.GetClassSaveProficiency("Constitution"), now)
	target.Top.Set("IntelligenceSaveCheck", char.GetClassSaveProficiency("Intelligence"), now)
	target.Top.Set("WisdomSaveCheck", char.GetClassSaveProficiency("Wisdom"), now)
	target.Top.Set("CharismaSaveCheck", char.GetClassSaveProficiency("Charisma"), now)

	// skills
	target.Top.Set("AcrobaticsCheck", char.GetProficiency("Acrobatics"), now)
	target.Top.Set("InsightCheck", char.GetProficiency("Insight"), now)
	target.Top.Set("PerformanceCheck", char.GetProficiency("Performance"), now)
	target.Top.Set("AnimalHandlingCheck", char.GetProficiency("AnimalHandling"), now)
	target.Top.Set("IntimidationCheck", char.GetProficiency("Intimidation"), now)
	target.Top.Set("PersuasionCheck", char.GetProficiency("Persuasion"), now)
	target.Top.Set("ArcanaCheck", char.GetProficiency("Arcana"), now)
	target.Top.Set("InvestigationCheck", char.GetProficiency("Investigation"), now)
	target.Top.Set("ReligionCheck", char.GetProficiency("Religion"), now)
	target.Top.Set("AthleticsCheck", char.GetProficiency("Athletics"), now)
	target.Top.Set("MedicineCheck", char.GetProficiency("Medicine"), now)
	target.Top.Set("SleightOfHandCheck", char.GetProficiency("SleightOfHand"), now)
	target.Top.Set("DeceptionCheck", char.GetProficiency("Deception"), now)
	target.Top.Set("NatureCheck", char.GetProficiency("Nature"), now)
	target.Top.Set("StealthCheck", char.GetProficiency("Stealth"), now)
	target.Top.Set("HistoryCheck", char.GetProficiency("History"), now)
	target.Top.Set("PerceptionCheck", char.GetProficiency("Perception"), now)
	target.Top.Set("SurvivalCheck", char.GetProficiency("Survival"), now)

	// bio
	target.Bio.Set("Personality", char.Data.Traits.PersonalityTraits, now)
	target.Bio.Set("Ideals", char.Data.Traits.Ideals, now)
	target.Bio.Set("Bonds", char.Data.Traits.Bonds, now)
	target.Bio.Set("Flaws", char.Data.Traits.Flaws, now)
	target.Bio.Set("Features", char.Data.Traits.Appearance, now)
	target.Bio.Set("Notes", char.GetNotes(), now)

	// items
	target.Items.Set("Copper", strconv.Itoa(char.Data.Currencies.Cp), now)
	target.Items.Set("Silver", strconv.Itoa(char.Data.Currencies.Sp), now)
	target.Items.Set("Electrum", strconv.Itoa(char.Data.Currencies.Ep), now)
	target.Items.Set("Gold", strconv.Itoa(char.Data.Currencies.Gp), now)
	target.Items.Set("Platinum", strconv.Itoa(char.Data.Currencies.Pp), now)

	items := char.GetInventory(now)
	for index, item := range items {
		id := "Items.Name." + strconv.Itoa(index)
		target.Items.Set(id, item.Name.Val, item.Name.TS)
		id = "Items.Properties." + strconv.Itoa(index)
		target.Items.Set(id, item.Properties.Val, item.Properties.TS)
		id = "Items.Weight." + strconv.Itoa(index)
		target.Items.Set(id, item.Weight.Val, item.Weight.TS)
	}

	// HP
	modifier := (char.GetAbilityInt("Constitution") - 10) / 2
	maxHP := modifier*char.Data.Classes[0].Level + char.Data.BaseHitPoints
	target.Combat.Set("CurrentHP", strconv.Itoa(maxHP-char.Data.RemovedHitPoints), now)
	target.Combat.Set("MaxHP", strconv.Itoa(maxHP), now)
	target.Combat.Set("TemporaryHP", strconv.Itoa(char.Data.TemporaryHitPoints), now)

	// weapons
	weapons := char.GetWeapons(now)
	for index, weapon := range weapons {
		id := "Weapons.Name." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.Name.Val, weapon.Name.TS)
		id = "Weapons.Properties." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.Properties.Val, weapon.Properties.TS)
		id = "Weapons.DamageRoll." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.DamageRoll.Val, weapon.DamageRoll.TS)
		id = "Weapons.DamageDice." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.DamageDice.Val, weapon.DamageDice.TS)
		id = "Weapons.Weight." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.Weight.Val, weapon.Weight.TS)
		id = "Weapons.Proficient." + strconv.Itoa(index)
		target.Combat.Set(id, weapon.Proficient.Val, weapon.Proficient.TS)
	}

	// armor
	armors := char.GetArmor(now)
	for index, armor := range armors {
		id := "Armors.Name." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Name.Val, armor.Name.TS)
		id = "Armors.Class." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Class.Val, armor.Class.TS)
		id = "Armors.Strength." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Strength.Val, armor.Strength.TS)
		id = "Armors.Stealth." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Stealth.Val, armor.Stealth.TS)
		id = "Armors.Properties." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Properties.Val, armor.Properties.TS)
		id = "Armors.Type." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Type.Val, armor.Type.TS)
		id = "Armors.InUse." + strconv.Itoa(index)
		target.Combat.Set(id, armor.InUse.Val, armor.InUse.TS)
		id = "Armors.Weight." + strconv.Itoa(index)
		target.Combat.Set(id, armor.Weight.Val, armor.Weight.TS)
	}

	// spells
	// 2do: used spell counts: is set in export data
	//      levelSpellSlots: max spells per level
	spells := char.GetSpells(now)
	for index, spell := range spells {
		id := "Spells.Name." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Name.Val, spell.Name.TS)
		id = "Spells.Prepared." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Prepared.Val, spell.Prepared.TS)
		id = "Spells.Level." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Level.Val, spell.Level.TS)
		id = "Spells.DamageRoll." + strconv.Itoa(index)
		target.Spells.Set(id, spell.DamageRoll.Val, spell.DamageRoll.TS)
		id = "Spells.DamageDice." + strconv.Itoa(index)
		target.Spells.Set(id, spell.DamageDice.Val, spell.DamageDice.TS)
		id = "Spells.AdditionalDamage." + strconv.Itoa(index)
		target.Spells.Set(id, spell.AdditionalDamage.Val, spell.AdditionalDamage.TS)
		id = "Spells.Type." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Type.Val, spell.Type.TS)
		id = "Spells.School." + strconv.Itoa(index)
		target.Spells.Set(id, spell.School.Val, spell.School.TS)
		id = "Spells.CastingTime." + strconv.Itoa(index)
		target.Spells.Set(id, spell.CastingTime.Val, spell.CastingTime.TS)
		id = "Spells.Range." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Range.Val, spell.Range.TS)
		id = "Spells.Duration." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Duration.Val, spell.Duration.TS)
		id = "Spells.ComponentVerbal." + strconv.Itoa(index)
		target.Spells.Set(id, spell.ComponentVerbal.Val, spell.ComponentVerbal.TS)
		id = "Spells.ComponentSomatic." + strconv.Itoa(index)
		target.Spells.Set(id, spell.ComponentSomatic.Val, spell.ComponentSomatic.TS)
		id = "Spells.ComponentMaterial." + strconv.Itoa(index)
		target.Spells.Set(id, spell.ComponentMaterial.Val, spell.ComponentMaterial.TS)
		id = "Spells.Components." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Components.Val, spell.Components.TS)
		id = "Spells.Description." + strconv.Itoa(index)
		target.Spells.Set(id, spell.Description.Val, spell.Description.TS)
	}

	return target
}
