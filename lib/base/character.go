package base

// Values in this definition are not easily changed as it requires frontend migration code

var (
	// NotNPC means this is not an NPC value
	NotNPC NpcVal = 0

	// IsNPC means this is an NPC value
	IsNPC NpcVal = 1
)

// NpcVal is an integer that is either NotNPC or IsNPC
type NpcVal int

// Character contains character sheet data
type Character struct {
	UserID string // The game-wide, unique value for the character data

	IsNPC              NpcVal
	IsShared           bool
	IsSharedWithUserID string

	// Character sheet tabs
	Top    CharacterSaver
	Bio    CharacterSaver
	Combat CharacterSaver
	Spells CharacterSaver
	Items  CharacterSaver

	// For paid users this value will be stored in the db.
	// If they open a browser with old data, the browser will then request the newer data from the db.
	UpdatedAt int64 // unix time
}

// Getter panics for a nonexistant value
func (character *Character) Getter(tab string) CharacterSaver {
	switch tab {
	case "character_top":
		return character.Top
	case "character_bio":
		return character.Bio
	case "character_combat":
		return character.Combat
	case "character_items":
		return character.Items
	case "character_spells":
		return character.Spells
	default:
		panic("No such getter: " + tab)
	}
}
