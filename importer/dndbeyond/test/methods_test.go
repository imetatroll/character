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
