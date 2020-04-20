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
	target.Top.Set("CharacterName", char.Character.Name, now)
	target.Top.Set("Race", char.Character.Race.FullName, now) // or BaseName
	target.Top.Set("Background", char.Character.Background.Definition.Name, now)
	target.Top.Set("Class", char.Character.Classes[0].Definition.Name, now)
	target.Top.Set("Alignment", char.GetAlignment(), now)
	target.Top.Set("XP", strconv.Itoa(char.Character.CurrentXp), now)
	target.Top.Set("Level", strconv.Itoa(char.Character.Classes[0].Level), now)

	target.Top.Set("Strength", char.GetAbility("Strength"), now)
	target.Top.Set("Dexterity", char.GetAbility("Dexterity"), now)
	target.Top.Set("Constitution", char.GetAbility("Constitution"), now)
	target.Top.Set("Intelligence", char.GetAbility("Intelligence"), now)
	target.Top.Set("Wisdom", char.GetAbility("Wisdom"), now)
	target.Top.Set("Charisma", char.GetAbility("Charisma"), now)

	// bio

	return target
}

/*
func (char *Character) GetSkill(name string) int64 {
			      { name: "acr", label: "Acrobatics", ability: "dex" },
		        { name: "ani", label: "Animal Handling", ability: "wis" },
		        { name: "arc", label: "Arcana", ability: "int" },
		        { name: "ath", label: "Athletics", ability: "str" },
		        { name: "dec", label: "Deception", ability: "cha" },
		        { name: "his", label: "History", ability: "int" },
		        { name: "ins", label: "Insight", ability: "wis" },
		        { name: "itm", label: "Intimidation", ability: "cha" },
		        { name: "inv", label: "Investigation", ability: "int" },
		        { name: "med", label: "Medicine", ability: "wis" },
		        { name: "nat", label: "Nature", ability: "int" },
		        { name: "prc", label: "Perception", ability: "wis" },
		        { name: "prf", label: "Performance", ability: "cha" },
		        { name: "per", label: "Persuasion", ability: "cha" },
		        { name: "rel", label: "Religion", ability: "int" },
		        { name: "slt", label: "Sleight of Hand", ability: "dex" },
		        { name: "ste", label: "Stealth", ability: "dex" },
		        { name: "sur", label: "Survival", ability: "wis" },
	return 0
}
*/

func (char *Character) GetRacialStatModifier(id int, name string) int {
	// EG "friendlySubtypeName": "Wisdom Score",
	//    "friendlySubtypeName": "Constitution Score",
	for _, mod := range char.Character.Modifiers.Race {
		if mod.Type == "bonus" && mod.FriendlySubtypeName == name+" Score" {
			return mod.Value
		}
	}
	return 0
}

func (char *Character) GetAbility(name string) string {
	id := -1
	switch name {
	case "Strength":
		id = 1
	case "Dexterity":
		id = 2
	case "Constitution":
		id = 3
	case "Intelligence":
		id = 4
	case "Wisdom":
		id = 5
	case "Charisma":
		id = 6
	}
	bonus := char.GetRacialStatModifier(id, name)
	for _, stat := range char.Character.Stats {
		if stat.ID == id {
			return strconv.Itoa(stat.Value + bonus)
		}
	}
	return "0"
}

func (char *Character) GetAlignment() string {
	switch char.Character.AlignmentID {
	case 1:
		return "Lawful Good"
	case 2:
		return "Neutral Good"
	case 3:
		return "Chaotic Good"
	case 4:
		return "Lawful Neutral"
	case 5:
		return "True Neutral"
	case 6:
		return "Chaotic Neutral"
	case 7:
		return "Lawful Evil"
	case 8:
		return "Neutral Evil"
	case 9:
		return "Chaotic Evil"
	}
	return ""
}
