package dnd

import (
	"imetatroll.com/character.git/lib/base"

	"sort"
	"strconv"
)

func NewCharacter(userId string) *base.Character {
	c := &base.Character{}
	c.Top = &CharacterTop{}
	c.Bio = &CharacterBio{}
	c.Combat = &CharacterCombat{}
	c.Items = &CharacterItems{}
	c.Spells = &CharacterSpells{}

	c.UserID = userId
	return c
}

// Top
type CharacterTop struct {
	CharacterName      base.CharacterField
	Race               base.CharacterField
	Background         base.CharacterField
	Class              base.CharacterField
	Alignment          base.CharacterField
	Inspiration        base.CharacterField
	XP                 base.CharacterField
	Level              base.CharacterField
	ClassDice          base.CharacterField
	ClassDiceCount     base.CharacterField
	OtherProficiencies base.CharacterField

	// abilities
	Strength     base.CharacterField
	Dexterity    base.CharacterField
	Constitution base.CharacterField
	Intelligence base.CharacterField
	Wisdom       base.CharacterField
	Charisma     base.CharacterField

	// saves
	StrengthSaveModifier     base.CharacterField
	StrengthSaveCheck        base.CharacterField
	DexteritySaveModifier    base.CharacterField
	DexteritySaveCheck       base.CharacterField
	ConstitutionSaveModifier base.CharacterField
	ConstitutionSaveCheck    base.CharacterField
	IntelligenceSaveModifier base.CharacterField
	IntelligenceSaveCheck    base.CharacterField
	WisdomSaveModifier       base.CharacterField
	WisdomSaveCheck          base.CharacterField
	CharismaSaveModifier     base.CharacterField
	CharismaSaveCheck        base.CharacterField

	// skills
	AcrobaticsModifier     base.CharacterField
	AcrobaticsCheck        base.CharacterField
	InsightCheck           base.CharacterField
	InsightModifier        base.CharacterField
	PerformanceCheck       base.CharacterField
	PerformanceModifier    base.CharacterField
	AnimalHandlingCheck    base.CharacterField
	AnimalHandlingModifier base.CharacterField
	IntimidationCheck      base.CharacterField
	IntimidationModifier   base.CharacterField
	PersuasionCheck        base.CharacterField
	PersuasionModifier     base.CharacterField
	ArcanaCheck            base.CharacterField
	ArcanaModifier         base.CharacterField
	InvestigationCheck     base.CharacterField
	InvestigationModifier  base.CharacterField
	ReligionCheck          base.CharacterField
	ReligionModifier       base.CharacterField
	AthleticsCheck         base.CharacterField
	AthleticsModifier      base.CharacterField
	MedicineCheck          base.CharacterField
	MedicineModifier       base.CharacterField
	SleightOfHandCheck     base.CharacterField
	SleightOfHandModifier  base.CharacterField
	DeceptionCheck         base.CharacterField
	DeceptionModifier      base.CharacterField
	NatureCheck            base.CharacterField
	NatureModifier         base.CharacterField
	StealthCheck           base.CharacterField
	StealthModifier        base.CharacterField
	HistoryCheck           base.CharacterField
	HistoryModifier        base.CharacterField
	PerceptionCheck        base.CharacterField
	PerceptionModifier     base.CharacterField
	SurvivalCheck          base.CharacterField
	SurvivalModifier       base.CharacterField
}

func (c *CharacterTop) Get(id string) (string, int64) {
	return base.Get(c, id, "CharacterTop")
}

func (c *CharacterTop) Set(id, val string, ts int64) {
	base.Set(c, id, val, ts, "CharacterTop")
}

func (c *CharacterTop) CopyTo(char *base.Character) {
	char.Top = c
}

func (c *CharacterTop) MergeTo(into *base.Character) {
	base.Merge(c, into.Top)
}

func (c *CharacterTop) Ordered(by string) base.Orderable {
	return base.ByLevel{}
}

func (c *CharacterTop) Remove(property, index string) {
	println("CharacterTop values are not removeable.")
}

func (c *CharacterTop) Length(property string) int {
	println("CharacterTop has no length")
	return 0
}

// Bio
type CharacterBio struct {
	Personality base.CharacterField
	Ideals      base.CharacterField
	Bonds       base.CharacterField
	Flaws       base.CharacterField
	Features    base.CharacterField
	Bio         base.CharacterField
	Notes       base.CharacterField
}

func (c *CharacterBio) Get(id string) (string, int64) {
	return base.Get(c, id, "CharacterBio")
}

func (c *CharacterBio) Set(id, val string, ts int64) {
	base.Set(c, id, val, ts, "CharacterBio")
}

func (c *CharacterBio) CopyTo(char *base.Character) {
	char.Bio = c
}

func (c *CharacterBio) MergeTo(into *base.Character) {
	base.Merge(c, into.Bio)
}

func (c *CharacterBio) Ordered(by string) base.Orderable {
	return base.ByLevel{}
}

func (c *CharacterBio) Remove(property, index string) {
	println("CharacterBio values are not removeable.")
}

func (c *CharacterBio) Length(property string) int {
	println("CharacterBio has no length")
	return 0
}

// Combat

// Weapon
type CharacterWeapon struct {
	UUID base.CharacterField

	Name             base.CharacterField
	Proficient       base.CharacterField
	Critical         base.CharacterField
	DamageRoll       base.CharacterField
	DamageDice       base.CharacterField
	AdditionalDamage base.CharacterField
	Ability          base.CharacterField
	Properties       base.CharacterField
	Weight           base.CharacterField
}

// Armor
type CharacterArmor struct {
	UUID base.CharacterField

	Name       base.CharacterField
	InUse      base.CharacterField
	Class      base.CharacterField
	Type       base.CharacterField
	Stealth    base.CharacterField
	Strength   base.CharacterField
	Weight     base.CharacterField
	Properties base.CharacterField
}

type CharacterCombat struct {
	CurrentHP          base.CharacterField
	TemporaryHP        base.CharacterField
	MaxHP              base.CharacterField
	HitDice            base.CharacterField
	HitDiceMaximum     base.CharacterField
	Speed              base.CharacterField
	SpeedModifier      base.CharacterField
	InitiativeModifier base.CharacterField
	ArmorClassModifier base.CharacterField
	DeathSaveSuccess0  base.CharacterField
	DeathSaveSuccess1  base.CharacterField
	DeathSaveSuccess2  base.CharacterField
	DeathSaveFailure0  base.CharacterField
	DeathSaveFailure1  base.CharacterField
	DeathSaveFailure2  base.CharacterField

	Weapons []CharacterWeapon
	Armors  []CharacterArmor
}

func (c *CharacterCombat) Get(id string) (string, int64) {
	if len(id) > 5 {
		if id[0:7] == "Weapons" {
			return base.GetPath(c, id, "Weapons")
		}
		if id[0:6] == "Armors" {
			return base.GetPath(c, id, "Armors")
		}
	}
	return base.Get(c, id, "CharacterCombat")
}

func (c *CharacterCombat) Set(id, val string, timestamp int64) {
	if len(id) > 5 {
		if id[0:7] == "Weapons" {
			if path, index, ok := base.HTMLIdToStructPath(id, "Weapons"); ok {
				if index >= len(c.Weapons) {
					c.Weapons = append(c.Weapons, CharacterWeapon{})
				}
				base.SetPath(c, path, index, val, timestamp)
			}
			return
		}
		if id[0:6] == "Armors" {
			if path, index, ok := base.HTMLIdToStructPath(id, "Armors"); ok {
				if index >= len(c.Armors) {
					c.Armors = append(c.Armors, CharacterArmor{})
				}
				base.SetPath(c, path, index, val, timestamp)
			}
			return
		}
	}
	base.Set(c, id, val, timestamp, "CharacterCombat")
}

func (c *CharacterCombat) CopyTo(char *base.Character) {
	char.Combat = c
}

func (c *CharacterCombat) MergeTo(into *base.Character) {
	base.Merge(c, into.Combat)
}

func (c *CharacterCombat) Ordered(by string) base.Orderable {
	switch by {
	case "Weapons":
		ordered := make(base.ByLevel, 0, len(c.Weapons))
		for index, weapon := range c.Weapons {
			ordered = append(ordered, base.OrderedLevel{Index: strconv.Itoa(index), Level: index, Name: weapon.Name.Val})
		}
		sort.Sort(ordered)
		return base.Orderable(ordered)
	case "Armors":
		ordered := make(base.ByLevel, 0, len(c.Armors))
		for index, armor := range c.Armors {
			ordered = append(ordered, base.OrderedLevel{Index: strconv.Itoa(index), Level: index, Name: armor.Name.Val})
		}
		sort.Sort(ordered)
		return base.Orderable(ordered)
	}
	return base.ByLevel{}
}

func (c *CharacterCombat) Remove(property, index string) {
	switch property {
	case "Armors":
		clear := make([]CharacterArmor, 0)
		for at, armor := range c.Armors {
			if strconv.Itoa(at) != index {
				clear = append(clear, armor)
			}
		}
		c.Armors = clear
	case "Weapons":
		clear := make([]CharacterWeapon, 0)
		for at, weapon := range c.Weapons {
			if strconv.Itoa(at) != index {
				clear = append(clear, weapon)
			}
		}
		c.Weapons = clear
	}
}

func (c *CharacterCombat) Length(property string) int {
	switch property {
	case "Armors":
		return len(c.Armors)
	case "Weapons":
		return len(c.Weapons)
	default:
		println("Unknown CharacterCombat length for", property)
	}
	return 0
}

// Items
type CharacterItem struct {
	UUID base.CharacterField

	Name       base.CharacterField
	Properties base.CharacterField
	Weight     base.CharacterField
}

type CharacterItems struct {
	Copper   base.CharacterField
	Silver   base.CharacterField
	Electrum base.CharacterField
	Gold     base.CharacterField
	Platinum base.CharacterField

	Items []CharacterItem
}

func (c *CharacterItems) Get(id string) (string, int64) {
	if len(id) > 4 && id[0:5] == "Items" {
		return base.GetPath(c, id, "Items")
	}
	return base.Get(c, id, "CharacterItems")
}

func (c *CharacterItems) Set(id, val string, timestamp int64) {
	if len(id) > 4 && id[0:5] == "Items" {
		if path, index, ok := base.HTMLIdToStructPath(id, "Items"); ok {
			if index >= len(c.Items) {
				c.Items = append(c.Items, CharacterItem{})
			}
			base.SetPath(c, path, index, val, timestamp)
		}
		return
	}
	base.Set(c, id, val, timestamp, "CharacterCombat")
}

func (c *CharacterItems) Remove(property, index string) {
	clear := make([]CharacterItem, 0)
	for at, item := range c.Items {
		if strconv.Itoa(at) != index {
			clear = append(clear, item)
		}
	}
	c.Items = clear
}

func (c *CharacterItems) CopyTo(char *base.Character) {
	char.Items = c
}

func (c *CharacterItems) MergeTo(into *base.Character) {
	base.Merge(c, into.Items)
}

func (c *CharacterItems) Ordered(by string) base.Orderable {
	return base.ByLevel{}
}

func (c *CharacterItems) Length(property string) int {
	if property != "Items" {
		println("Unknown CharacterItems length for", property)
	}
	return len(c.Items)
}

// Spells
type CharacterSpell struct {
	UUID base.CharacterField

	Name              base.CharacterField
	Prepared          base.CharacterField
	Level             base.CharacterField
	DamageRoll        base.CharacterField
	DamageDice        base.CharacterField
	AdditionalDamage  base.CharacterField
	Type              base.CharacterField
	School            base.CharacterField
	CastingTime       base.CharacterField
	Range             base.CharacterField
	Duration          base.CharacterField
	ComponentVerbal   base.CharacterField
	ComponentSomatic  base.CharacterField
	ComponentMaterial base.CharacterField
	Components        base.CharacterField
	Description       base.CharacterField
}

type SpellCount struct {
	Current base.CharacterField
	Max     base.CharacterField
}

type CharacterSpells struct {
	SpellAbility    base.CharacterField
	SpellDCModifier base.CharacterField
	SpellCounts     [10]SpellCount
	Spells          []CharacterSpell
}

func (c *CharacterSpells) Ordered(by string) base.Orderable {
	switch by {
	case "Spells":
		ordered := make(base.ByLevel, 0, len(c.Spells))
		for index, spell := range c.Spells {
			level, err := strconv.Atoi(spell.Level.Val)
			if err != nil {
				println(err)
			}
			ordered = append(ordered, base.OrderedLevel{Index: strconv.Itoa(index), Level: level, Name: spell.Name.Val})
		}
		sort.Sort(ordered)
		return base.Orderable(ordered)
	}
	return base.ByLevel{}
}

func (c *CharacterSpells) Get(id string) (string, int64) {
	if id == "SpellAbility" {
		return c.SpellAbility.Val, c.SpellAbility.TS
	}
	if id == "SpellDCModifier" {
		return c.SpellDCModifier.Val, c.SpellDCModifier.TS
	}
	for i := 0; i < 10; i++ {
		if id == "SpellsRemaining"+strconv.Itoa(i) {
			return c.SpellCounts[i].Current.Val, c.SpellCounts[i].Current.TS
		}
		if id == "SpellCountMax"+strconv.Itoa(i) {
			return c.SpellCounts[i].Max.Val, c.SpellCounts[i].Max.TS
		}
	}
	return base.GetPath(c, id, "Spells")
}

func (c *CharacterSpells) Set(id, val string, timestamp int64) {
	if id == "SpellAbility" {
		c.SpellAbility.Val = val
		c.SpellAbility.TS = timestamp
		return
	}
	if id == "SpellDCModifier" {
		c.SpellDCModifier.Val = val
		c.SpellDCModifier.TS = timestamp
		return
	}
	for i := 0; i < 10; i++ {
		if id == "SpellsRemaining"+strconv.Itoa(i) {
			c.SpellCounts[i].Current.Val = val
			c.SpellCounts[i].Current.TS = timestamp
			return
		}
		if id == "SpellCountMax"+strconv.Itoa(i) {
			c.SpellCounts[i].Max.Val = val
			c.SpellCounts[i].Max.TS = timestamp
			return
		}
	}
	if path, index, ok := base.HTMLIdToStructPath(id, "Spells"); ok {
		if index >= len(c.Spells) {
			c.Spells = append(c.Spells, CharacterSpell{})
		}
		base.SetPath(c, path, index, val, timestamp)
	}
}

func (c *CharacterSpells) CopyTo(char *base.Character) {
	char.Spells = c
}

func (c *CharacterSpells) MergeTo(into *base.Character) {
	base.Merge(c, into.Spells)
}

func (c *CharacterSpells) Remove(property, index string) {
	clear := make([]CharacterSpell, 0)
	for at, spell := range c.Spells {
		if strconv.Itoa(at) != index {
			clear = append(clear, spell)
		}
	}
	c.Spells = clear
}

func (c *CharacterSpells) Length(property string) int {
	if property != "Spells" {
		println("Unknown CharacterSpells length for", property)
	}
	return len(c.Spells)
}
