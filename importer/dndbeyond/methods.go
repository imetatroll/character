package beyond

import (
	"strconv"
	"time"

	"imetatroll.com/character.git/lib/base"
	"imetatroll.com/character.git/lib/dnd"
)

func GetAlignment(id int) string {
	switch id {
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

func (char *Character) Transfer(userID string) *base.Character {
	now := time.Now().Unix()

	target := dnd.NewCharacter(userID)

	if char.Character.Inspiration {
		target.Top.Set("Inspiration", "true", now)
	} else {
		target.Top.Set("Inspiration", "false", now)
	}

	target.Top.Set("CharacterName", char.Character.Name, now)
	target.Top.Set("Race", char.Character.Race.FullName, now) // or BaseName
	target.Top.Set("Background", char.Character.Background.Definition.Name, now)
	target.Top.Set("Class", char.Character.Classes[0].Definition.Name, now)
	target.Top.Set("Alignment", GetAlignment(char.Character.AlignmentID), now)
	target.Top.Set("XP", strconv.Itoa(char.Character.CurrentXp), now)

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

func (char *Character) GetAbility(name string) int64 {
	id := -1
	switch name {
	case "str":
		id = 1
	case "dex":
		id = 2
	case "con":
		id = 3
	case "int":
		id = 4
	case "wis":
		id = 5
	case "cha":
		id = 6
	}
	for _, stat := range char.Character.Stats {
		if stat.ID == id {
			return int64(stat.Value)
		}
	}
	return 0
}
