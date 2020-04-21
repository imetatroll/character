package beyondtest

import (
	"io/ioutil"
	"strconv"
	"testing"

	"imetatroll.com/character.git/importer/dndbeyond"
	"imetatroll.com/character.git/lib/base"
)

func ReadCharacter(t *testing.T) *base.Character {
	data, err := ioutil.ReadFile("./unit_test.json")
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

	expect := "Karmana Sirake"
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

	expect := "Hill Dwarf"
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

	expect := "Noble"
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

	expect := "Neutral Evil"
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

	expect := "900"
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

	expect := "3"
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

	expect := "true"
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
		[]string{"Strength", "14"},
		[]string{"Dexterity", "12"},
		[]string{"Constitution", "10"},
		[]string{"Intelligence", "3"},
		[]string{"Wisdom", "13"},
		[]string{"Charisma", "11"},
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
		[]string{"StrengthSaveCheck", "false"},
		[]string{"DexteritySaveCheck", "false"},
		[]string{"ConstitutionSaveCheck", "true"},
		[]string{"IntelligenceSaveCheck", "false"},
		[]string{"WisdomSaveCheck", "false"},
		[]string{"CharismaSaveCheck", "true"},
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
		[]string{"AcrobaticsCheck", "false"},
		[]string{"InsightCheck", "true"},
		[]string{"PerformanceCheck", "false"},
		[]string{"AnimalHandlingCheck", "false"},
		[]string{"IntimidationCheck", "false"},
		[]string{"PersuasionCheck", "true"},
		[]string{"ArcanaCheck", "false"},
		[]string{"InvestigationCheck", "false"},
		[]string{"ReligionCheck", "false"},
		[]string{"AthleticsCheck", "false"},
		[]string{"MedicineCheck", "false"},
		[]string{"SleightOfHandCheck", "false"},
		[]string{"DeceptionCheck", "true"},
		[]string{"NatureCheck", "false"},
		[]string{"StealthCheck", "false"},
		[]string{"HistoryCheck", "true"},
		[]string{"PerceptionCheck", "false"},
		[]string{"SurvivalCheck", "false"},
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
		[]string{"Personality", "Despite my noble birth, I do not place myself above other folk. We all have the same blood."},
		[]string{"Ideals", "Power. If I can attain more power, no one will tell me what to do. (Evil)"},
		[]string{"Bonds", "I am in love with the heir of a family that my family despises."},
		[]string{"Flaws", "I hide a truly scandalous secret that could ruin my family forever."},
		[]string{"Features", "The way I look."},
		[]string{"Notes", notes},
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
		[]string{"Copper", "5"},
		[]string{"Silver", "4"},
		[]string{"Electrum", "3"},
		[]string{"Gold", "2"},
		[]string{"Platinum", "1"},
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
	if length != 1 {
		t.Fatalf("expecting 1 item but got %d items", length)
	}
	for index := 0; index < length; index++ {
		id := "Items.Name." + strconv.Itoa(index)
		val, _ := character.Items.Get(id)
		if val != "Rope, Hempen (50 feet)" {
			t.Fatalf("expecting 'Rope, Hempen (50 feet)' but got '%s'", val)
		}
		id = "Items.Properties." + strconv.Itoa(index)
		val, _ = character.Items.Get(id)
		if val != "Rope,&nbsp;has 2 hit points and can be burst with a DC 17 Strength check." {
			t.Fatalf("expecting 'Rope,&nbsp;has 2 hit points and can be burst with a DC 17 Strength check.' but got '%s'", val)
		}
		id = "Items.Weight." + strconv.Itoa(index)
		val, _ = character.Items.Get(id)
		if val != "10.0" {
			t.Fatalf("expecting '10.0' but got '%s'", val)
		}
	}
}

func TestWeapons(t *testing.T) {
	character := ReadCharacter(t)

	props := `Versatile
This weapon can be used with one or two hands. A damage value in parentheses appears with the property--the damage when the weapon is used with two hands to make a melee attack.
1d10
`

	length := character.Combat.Length("Weapons")
	if length != 1 {
		t.Fatalf("expecting 1 item but got %d items", length)
	}
	for index := 0; index < length; index++ {
		id := "Weapons.Name." + strconv.Itoa(index)
		val, _ := character.Combat.Get(id)
		if val != "Battleaxe" {
			t.Fatalf("expecting 'Battleaxe' but got '%s'", val)
		}
		id = "Weapons.Properties." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != props {
			t.Fatalf("expecting '%s' but got '%s'", props, val)
		}
		id = "Weapons.Weight." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "4.0" {
			t.Fatalf("expecting '4.0' but got '%s'", val)
		}
		id = "Weapons.Proficient." + strconv.Itoa(index)
		val, _ = character.Combat.Get(id)
		if val != "true" {
			t.Fatalf("expecting 'true' but got '%s'", val)
		}
	}
}

func TestFilterP(t *testing.T) {
	val := "<p>Rope,&nbsp;has 2 hit points</p>\n<p>17 Strength check</p>"
	val = beyond.FilterP(val)
	if val != "Rope,&nbsp;has 2 hit points\n17 Strength check" {
		t.Fatalf("expected 'Rope,&nbsp;has 2 hit points\n17 Strength check' but got '%s'", val)
	}
}
