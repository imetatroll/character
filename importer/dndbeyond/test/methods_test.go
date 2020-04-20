package beyondtest

import (
	"io/ioutil"
	"testing"

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
