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

	if char.Character.Inspiration {
		target.Top.Set("Inspiration", "true", now)
	} else {
		target.Top.Set("Inspiration", "false", now)
	}

	// top
	// 2do: class dice/count
	// 2do: save modifiers - if i remember correctly these are actually calculated values
	target.Top.Set("CharacterName", char.Character.Name, now)
	target.Top.Set("Race", char.Character.Race.FullName, now) // or BaseName
	target.Top.Set("Background", char.Character.Background.Definition.Name, now)
	target.Top.Set("Class", char.Character.Classes[0].Definition.Name, now)
	target.Top.Set("Alignment", char.GetAlignment(), now)
	target.Top.Set("XP", strconv.Itoa(char.Character.CurrentXp), now)
	target.Top.Set("Level", strconv.Itoa(char.Character.Classes[0].Level), now)

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

	return target
}
