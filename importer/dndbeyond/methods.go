package beyond

import (
	"strconv"
)

func (char *Character) GetProficiency(name string) string {
	if char.GetRaceProficiency(name) == "true" {
		return "true"
	}
	if char.GetClassProficiency(name) == "true" {
		return "true"
	}
	if char.GetBackgroundProficiency(name) == "true" {
		return "true"
	}
	return "false"
}

func (char *Character) GetRaceProficiency(name string) string {
	for _, mod := range char.Character.Modifiers.Race {
		if mod.Type == "proficiency" && mod.FriendlySubtypeName == name {
			return "true"
		}
	}
	return "false"
}

// "suggestedProficiencies": [ "History", "Persuasion" ]
func (char *Character) GetBackgroundProficiency(name string) string {
	for _, mod := range char.Character.Modifiers.Background {
		if mod.Type == "proficiency" && mod.FriendlySubtypeName == name {
			return "true"
		}
	}
	return "false"
}

// EG "friendlySubtypeName": "Persuasion",
//    "friendlySubtypeName": "Deception",
func (char *Character) GetClassProficiency(name string) string {
	for _, mod := range char.Character.Modifiers.Class {
		if mod.Type == "proficiency" && mod.FriendlySubtypeName == name {
			return "true"
		}
	}
	return "false"
}

// EG "friendlySubtypeName": "Constitution Saving Throws",
//    "friendlySubtypeName": "Charisma Saving Throws",
func (char *Character) GetClassSaveProficiency(name string) string {
	for _, mod := range char.Character.Modifiers.Class {
		if mod.Type == "proficiency" && mod.FriendlySubtypeName == name+" Saving Throws" {
			return "true"
		}
	}
	return "false"
}

// EG "friendlySubtypeName": "Wisdom Score",
//    "friendlySubtypeName": "Constitution Score",
func (char *Character) GetRacialAbilityModifier(id int, name string) int {
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
	bonus := char.GetRacialAbilityModifier(id, name)
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
