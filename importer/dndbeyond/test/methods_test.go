package beyondtest

import (
	"io/ioutil"
	"strconv"
	"testing"

	"imetatroll.com/character.git/importer/dndbeyond"
	"imetatroll.com/character.git/lib/base"
)

func ReadCharacter(t *testing.T) *base.Character {
	data, err := ioutil.ReadFile("./2020-10-09.json")
	if err != nil {
		t.Fatal(err)
	}
	beyond, err := NewBeyondCharacter(string(data))
	if err != nil {
		t.Fatal(err)
	}
	return beyond.Transfer("1")
}

func TestCharacterName(t *testing.T) {
	character := ReadCharacter(t)

	expect := "Harkam Wynrim"
	name, ts := character.Top.Get("CharacterName")
	if name != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, name)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestRace(t *testing.T) {
	character := ReadCharacter(t)

	expect := "Scourge Aasimar"
	race, ts := character.Top.Get("Race")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestBackground(t *testing.T) {
	character := ReadCharacter(t)

	expect := "Adolescent"
	race, ts := character.Top.Get("Background")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestAlignment(t *testing.T) {
	character := ReadCharacter(t)

	expect := "Chaotic Good"
	race, ts := character.Top.Get("Alignment")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestClass(t *testing.T) {
	character := ReadCharacter(t)

	expect := "Sorcerer"
	race, ts := character.Top.Get("Class")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestXP(t *testing.T) {
	character := ReadCharacter(t)

	expect := "0"
	race, ts := character.Top.Get("XP")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestLevel(t *testing.T) {
	character := ReadCharacter(t)

	expect := "1"
	race, ts := character.Top.Get("Level")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestInspiration(t *testing.T) {
	character := ReadCharacter(t)

	expect := "false"
	race, ts := character.Top.Get("Inspiration")
	if race != expect {
		t.Fatalf("expecting '%s' but got '%s'", expect, race)
	}
	if ts == 0 {
		t.Fatal("expecting non-zero timestamp")
	}
}

func TestAbilities(t *testing.T) {
	character := ReadCharacter(t)

	abilities := [][]string{
		{"Strength", "10"},
		{"Dexterity", "13"},
		{"Constitution", "13"},
		{"Intelligence", "12"},
		{"Wisdom", "11"},
		{"Charisma", "17"},
	}
	for _, ability := range abilities {
		val, ts := character.Top.Get(ability[0])
		if val != ability[1] {
			t.Fatalf("%s expecting '%s' but got '%s'", ability[0], ability[1], val)
		}
		if ts == 0 {
			t.Fatal("expecting non-zero timestamp")
		}
	}
}

func TestSavingThrows(t *testing.T) {
	character := ReadCharacter(t)

	saves := [][]string{
		{"StrengthSaveCheck", "false"},
		{"DexteritySaveCheck", "false"},
		{"ConstitutionSaveCheck", "true"},
		{"IntelligenceSaveCheck", "false"},
		{"WisdomSaveCheck", "false"},
		{"CharismaSaveCheck", "true"},
	}
	for _, save := range saves {
		val, ts := character.Top.Get(save[0])
		if val != save[1] {
			t.Fatalf("%s expecting '%s' but got '%s'", save[0], save[1], val)
		}
		if ts == 0 {
			t.Fatal("expecting non-zero timestamp")
		}
	}
}

func TestSkills(t *testing.T) {
	character := ReadCharacter(t)

	skills := [][]string{
		{"AcrobaticsCheck", "false"},
		{"InsightCheck", "true"},
		{"PerformanceCheck", "false"},
		{"AnimalHandlingCheck", "false"},
		{"IntimidationCheck", "false"},
		{"PersuasionCheck", "false"},
		{"ArcanaCheck", "true"},
		{"InvestigationCheck", "false"},
		{"ReligionCheck", "false"},
		{"AthleticsCheck", "false"},
		{"MedicineCheck", "false"},
		{"SleightOfHandCheck", "false"},
		{"DeceptionCheck", "false"},
		{"NatureCheck", "false"},
		{"StealthCheck", "false"},
		{"HistoryCheck", "false"},
		{"PerceptionCheck", "false"},
		{"SurvivalCheck", "false"},
	}
	for _, skill := range skills {
		val, ts := character.Top.Get(skill[0])
		if val != skill[1] {
			t.Fatalf("%s expecting '%s' but got '%s'", skill[0], skill[1], val)
		}
		if ts == 0 {
			t.Fatal("expecting non-zero timestamp")
		}
	}
}

func TestBio(t *testing.T) {
	character := ReadCharacter(t)

	notes := `Organizations

Allies

Enemies

Backstory Here

Other Information

`

	entries := [][]string{
		{"Personality", "Despite my noble birth, I do not place myself above other folk. We all have the same blood."},
		{"Ideals", "Power. If I can attain more power, no one will tell me what to do. (Evil)"},
		{"Bonds", "I am in love with the heir of a family that my family despises."},
		{"Flaws", "I hide a truly scandalous secret that could ruin my family forever."},
		{"Features", "The way I look."},
		{"Notes", notes},
	}
	for _, entry := range entries {
		val, ts := character.Bio.Get(entry[0])
		if val != entry[1] {
			t.Fatalf("%s expecting '%s' but got '%s'", entry[0], entry[1], val)
		}
		if ts == 0 {
			t.Fatal("expecting non-zero timestamp")
		}
	}
}

func TestCurrencies(t *testing.T) {
	character := ReadCharacter(t)

	currencies := [][]string{
		{"Copper", "5"},
		{"Silver", "4"},
		{"Electrum", "3"},
		{"Gold", "2"},
		{"Platinum", "1"},
	}
	for _, currency := range currencies {
		val, ts := character.Items.Get(currency[0])
		if val != currency[1] {
			t.Fatalf("%s expecting '%s' but got '%s'", currency[0], currency[1], val)
		}
		if ts == 0 {
			t.Fatal("expecting non-zero timestamp")
		}
	}
}

func TestItems(t *testing.T) {
	character := ReadCharacter(t)

	length := character.Items.Length("Items")
	if length != 11 {
		t.Fatalf("expecting 11 items but got %d items", length)
	}

	// check first item
	id := "Items.Name.0"
	val, _ := character.Items.Get(id)
	if val != "Crossbow Bolts" {
		t.Fatalf("expecting 'Crossbow Bolts' but got '%s'", val)
	}
	id = "Items.Properties.0"
	val, _ = character.Items.Get(id)
	if val[0:len("Crossbow bolts are")] != "Crossbow bolts are" {
		t.Fatalf("expecting 'Crossbow bolts are'... but got '%s'", val)
	}
	id = "Items.Weight.0"
	val, _ = character.Items.Get(id)
	if val != "1.5" {
		t.Fatalf("expecting '1.5' but got '%s'", val)
	}
}

func TestWeapons(t *testing.T) {
	character := ReadCharacter(t)

	length := character.Combat.Length("Weapons")
	if length != 3 {
		t.Fatalf("expecting 3 weapons but got %d items", length)
	}

	// check first weapon
	id := "Weapons.Name.0"
	val, _ := character.Combat.Get(id)
	if val != "Dagger" {
		t.Fatalf("expecting 'Battleaxe' but got '%s'", val)
	}
	id = "Weapons.Properties.0"
	val, _ = character.Combat.Get(id)
	if val[0:len("Finesse")] != "Finesse" {
		t.Fatalf("expecting 'Finesse'... but got '%s'", val)
	}
	id = "Weapons.Weight.0"
	val, _ = character.Combat.Get(id)
	if val != "1.0" {
		t.Fatalf("expecting '1.0' but got '%s'", val)
	}
	id = "Weapons.Proficient.0"
	val, _ = character.Combat.Get(id)
	if val != "true" {
		t.Fatalf("expecting 'true' but got '%s'", val)
	}
}

func TestArmor(t *testing.T) {
	character := ReadCharacter(t)

	length := character.Combat.Length("Armors")
	if length != 1 {
		t.Fatalf("expecting 1 armor but got %d items", length)
	}
	for index := 0; index < length; index++ {
		id := "Armors.Name." + strconv.Itoa(index)
		val, _ := character.Combat.Get(id)
		if val != "Chain Mail" {
			t.Fatalf("expecting 'Chain Mail' but got '%s'", val)
		}
		id = "Armors.Class." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "16" {
			t.Fatalf("expecting '16' but got '%s'", val)
		}
		id = "Armors.Strength." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "Str 13" {
			t.Fatalf("expecting 'Str 13' but got '%s'", val)
		}
		id = "Armors.Stealth." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "Disadvantage" {
			t.Fatalf("expecting 'Disadvantage' but got '%s'", val)
		}
		id = "Armors.Properties." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "Made of interlocking metal rings, chain mail includes a layer of quilted fabric worn underneath the mail to prevent chafing and to cushion the impact of blows. The suit includes gauntlets." {
			t.Fatalf("expecting 'Made of ...' but got '%s'", val)
		}
		id = "Armors.Type." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "Heavy" {
			t.Fatalf("expecting 'Heavy' but got '%s'", val)
		}
		id = "Armors.InUse." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "true" {
			t.Fatalf("expecting 'true' but got '%s'", val)
		}
		id = "Armors.Weight." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "55.0" {
			t.Fatalf("expecting '55.0' but got '%s'", val)
		}
	}
}

func TestSpells(t *testing.T) {
	character := ReadCharacter(t)

	length := character.Spells.Length("Spells")
	if length != 1 {
		t.Fatalf("expecting 1 spells but got %d items", length)
	}
	index := 0

	id := "Spells.Name." + strconv.Itoa(index)
	val, _ := character.Spells.Get(id)
	if val != "Fire Bolt" {
		t.Fatalf("expecting 'Fire Bolt' but got '%s'", val)
	}
	id = "Spells.Prepared." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "false" {
		t.Fatalf("expecting 'false' but got '%s'", val)
	}
	id = "Spells.Level." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "0" {
		t.Fatalf("expecting '0' but got '%s'", val)
	}
	id = "Spells.DamageRoll." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "1" {
		t.Fatalf("expecting '1' but got '%s'", val)
	}
	id = "Spells.DamageDice." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "10" {
		t.Fatalf("expecting '10' but got '%s'", val)
	}
	id = "Spells.AdditionalDamage." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "0" {
		t.Fatalf("expecting '0' but got '%s'", val)
	}
	id = "Spells.Type." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "DamageAttack" {
		t.Fatalf("expecting 'DamageAttack' but got '%s'", val)
	}
	id = "Spells.School." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "Evocation" {
		t.Fatalf("expecting 'Evocation' but got '%s'", val)
	}
	id = "Spells.CastingTime." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "0" {
		t.Fatalf("expecting '0' but got '%s'", val)
	}
	id = "Spells.Range." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "Ranged 120" {
		t.Fatalf("expecting 'Ranged 120' but got '%s'", val)
	}
	id = "Spells.Duration." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "1 action" {
		t.Fatalf("expecting '1 action' but got '%s'", val)
	}
	id = "Spells.ComponentVerbal." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "true" {
		t.Fatalf("expecting 'true' but got '%s'", val)
	}
	id = "Spells.ComponentSomatic." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "true" {
		t.Fatalf("expecting 'true' but got '%s'", val)
	}
	id = "Spells.ComponentMaterial." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "false" {
		t.Fatalf("expecting 'false' but got '%s'", val)
	}
	id = "Spells.Components." + strconv.Itoa(index)
	val, _ = character.Spells.Get(id)
	if val != "" {
		t.Fatalf("expecting '' but got '%s'", val)
	}
}

func TestFilterP(t *testing.T) {
	val := "<p>Rope,&nbsp;has 2 hit points</p>\n<p>17 Strength check</p>"
	val = beyond.FilterP(val)
	if val != "Rope,&nbsp;has 2 hit points\n17 Strength check" {
		t.Fatalf("expected 'Rope,&nbsp;has 2 hit points\n17 Strength check' but got '%s'", val)
	}
}
