package beyond

import (
	"strconv"
	"strings"

	"imetatroll.com/character.git/lib/base"
	"imetatroll.com/character.git/lib/dnd"
)

var NonItems = []string{"Armor", "Weapon", "Wand"}

func IsNonItem(filterType string) bool {
	for _, non := range NonItems {
		if filterType == non {
			return true
		}
	}
	return false
}

func FilterP(val string) string {
	filtered := make([]byte, 0, len(val))
	for i := 0; i < len(val); i++ {
		if val[i] == '<' && i+2 < len(val) && val[i+1] == 'p' && val[i+2] == '>' {
			i += 2
			continue
		}
		if val[i] == '<' && i+3 < len(val) && val[i+1] == '/' && val[i+2] == 'p' && val[i+3] == '>' {
			i += 3
			continue
		}
		filtered = append(filtered, val[i])
	}
	return string(filtered)
}

func (char *Character) GetInventory(now int64) []dnd.CharacterItem {
	items := []dnd.CharacterItem{}
	for index, item := range char.Character.Inventory {
		if !IsNonItem(item.Definition.FilterType) {
			item := dnd.CharacterItem{
				UUID: base.CharacterField{
					Val: strconv.Itoa(index),
					TS:  now,
				},
				Name: base.CharacterField{
					Val: item.Definition.Name,
					TS:  now,
				},
				Properties: base.CharacterField{
					Val: FilterP(item.Definition.Description),
					TS:  now,
				},
				Weight: base.CharacterField{
					Val: strconv.FormatFloat(item.Definition.Weight, 'f', 1, 64),
					TS:  now,
				},
			}
			items = append(items, item)
		}
	}
	return items
}

func (char *Character) GetWeapons(now int64) []dnd.CharacterWeapon {
	weapons := []dnd.CharacterWeapon{}
	for index, item := range char.Character.Inventory {
		if item.Definition.FilterType == "Weapon" {
			/*
				Critical         base.CharacterField
				AdditionalDamage base.CharacterField
				Ability          base.CharacterField
			*/
			properties := ""
			for _, prop := range item.Definition.Properties {
				properties += strings.TrimSpace(prop.Name) + "\n"
				properties += strings.TrimSpace(prop.Description) + "\n"
				properties += strings.TrimSpace(prop.Notes) + "\n"
			}
			weapon := dnd.CharacterWeapon{
				UUID: base.CharacterField{
					Val: strconv.Itoa(index),
					TS:  now,
				},
				Name: base.CharacterField{
					Val: item.Definition.Name,
					TS:  now,
				},
				Properties: base.CharacterField{
					Val: properties,
					TS:  now,
				},
				DamageRoll: base.CharacterField{
					Val: strconv.Itoa(item.Definition.Damage.DiceCount),
					TS:  now,
				},
				DamageDice: base.CharacterField{
					Val: strconv.Itoa(item.Definition.Damage.DiceValue),
					TS:  now,
				},
				Weight: base.CharacterField{
					Val: strconv.FormatFloat(item.Definition.Weight, 'f', 1, 64),
					TS:  now,
				},
			}
			if char.GetProficiency(item.Definition.Name) == "true" {
				weapon.Proficient = base.CharacterField{
					Val: "true",
					TS:  now,
				}
			}
			weapons = append(weapons, weapon)
		}
	}
	return weapons
}

func (char *Character) GetArmor(now int64) []dnd.CharacterArmor {
	armors := []dnd.CharacterArmor{}
	for index, item := range char.Character.Inventory {
		if item.Definition.FilterType == "Armor" {
			/*
				Stealth
			*/
			equipped := "false"
			if item.Equipped {
				equipped = "true"
			}
			stealth := ""
			if item.Definition.StealthCheck > 1 {
				stealth = "Disadvantage"
			}
			armor := dnd.CharacterArmor{
				UUID: base.CharacterField{
					Val: strconv.Itoa(index),
					TS:  now,
				},
				Name: base.CharacterField{
					Val: item.Definition.Name,
					TS:  now,
				},
				Class: base.CharacterField{
					Val: strconv.Itoa(item.Definition.ArmorClass),
					TS:  now,
				},
				Strength: base.CharacterField{
					Val: "Str " + strconv.Itoa(item.Definition.StrengthRequirement),
					TS:  now,
				},
				Stealth: base.CharacterField{
					Val: stealth,
					TS:  now,
				},
				Properties: base.CharacterField{
					Val: FilterP(item.Definition.Description),
					TS:  now,
				},
				Type: base.CharacterField{
					Val: strings.Split(item.Definition.Type, " ")[0],
					TS:  now,
				},
				InUse: base.CharacterField{
					Val: equipped,
					TS:  now,
				},
				Weight: base.CharacterField{
					Val: strconv.FormatFloat(item.Definition.Weight, 'f', 1, 64),
					TS:  now,
				},
			}
			armors = append(armors, armor)
		}
	}
	return armors
}

func (char *Character) GetNotes() string {
	notes := strings.TrimSpace(char.Character.Notes.Organizations)
	if len(notes) > 0 {
		notes += "\n\n"
	}
	val := strings.TrimSpace(char.Character.Notes.Allies)
	if len(val) > 0 {
		notes += val + "\n\n"
	}
	val = strings.TrimSpace(char.Character.Notes.Enemies)
	if len(val) > 0 {
		notes += val + "\n\n"
	}
	val = strings.TrimSpace(char.Character.Notes.Backstory)
	if len(val) > 0 {
		notes += val + "\n\n"
	}
	val = strings.TrimSpace(char.Character.Notes.OtherNotes)
	if len(val) > 0 {
		notes += val + "\n\n"
	}
	return notes
}

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
