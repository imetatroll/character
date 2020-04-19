package dnd

import (
	"testing"
	"time"
)

func TestCharacterTopMerge(t *testing.T) {
	c1 := NewCharacter("1")
	c2 := NewCharacter("1")

	testName := "My New Name"
	c2.Top.Set("CharacterName", testName, 1)

	testLevel := "15"
	c2.Top.Set("Level", testLevel, 1)

	c2.Top.MergeTo(c1)

	if name, _ := c1.Top.Get("CharacterName"); name != testName {
		t.Logf("%+v", c2.Top)
		t.Fatalf("missing character name '%s' from %+v", testName, c1.Top)
	}
	if name, _ := c1.Top.Get("Level"); name != testLevel {
		t.Logf("%+v", c2.Top)
		t.Fatalf("missing level '%s' from %+v", testLevel, c1.Top)
	}
}

func TestCharacterWeaponsMerge(t *testing.T) {
	c1 := NewCharacter("1")
	c2 := NewCharacter("1")

	uuid := "uuid-val"
	testName := "My Weapon"

	c1.Combat.Set("Weapons.UUID.0", uuid, 0)
	c1.Combat.Set("Weapons.Name.0", testName, 1)
	if name, _ := c1.Combat.Get("Weapons.Name.0"); name != testName {
		t.Fatalf("missing weapon name '%s' from %+v", testName, c1.Combat)
	}
	time.Sleep(time.Millisecond * 100)

	testName = "My New Weapon"
	c2.Combat.Set("Weapons.UUID.0", uuid, 0)
	c2.Combat.Set("Weapons.Name.0", testName, 2)
	c2.Combat.MergeTo(c1)

	if name, _ := c1.Combat.Get("Weapons.Name.0"); name != testName {
		t.Logf("%+v", c2.Combat)
		t.Fatalf("missing weapon name '%s' from %+v", testName, c1.Combat)
	}
}

func TestCharacterWeaponsAddMerge(t *testing.T) {
	c1 := NewCharacter("1")
	c2 := NewCharacter("1")

	c1.Combat.Set("Weapons.UUID.0", "no-collision", 0)
	c1.Combat.Set("Weapons.Name.0", "other weapon", 2)

	uuid := "uuid-val"
	testName := "My New Weapon"
	c2.Combat.Set("Weapons.UUID.0", "no-collision", 0)
	c2.Combat.Set("Weapons.Name.0", "other weapon", 2)
	c2.Combat.Set("Weapons.UUID.1", uuid, 0)
	c2.Combat.Set("Weapons.Name.1", testName, 2)
	c2.Combat.MergeTo(c1)

	if name, _ := c1.Combat.Get("Weapons.Name.1"); name != testName {
		t.Logf("%+v", c2.Combat)
		t.Fatalf("missing weapon name '%s' from %+v", testName, c1.Combat)
	}
}

func TestCharacterWeaponsDeleteMerge(t *testing.T) {
	c1 := NewCharacter("1")
	c2 := NewCharacter("1")

	c1.Combat.Set("Weapons.UUID.0", "must-be-deleted", 0)
	c1.Combat.Set("Weapons.Name.0", "delete this weapon", 1)

	c2.Combat.MergeTo(c1)

	if c1.Combat.Length("Weapons") != 0 {
		t.Logf("%+v", c2.Combat)
		t.Fatalf("expecting deleted weapon name 'delete this weapon' to be gone %+v", c1.Combat)
	}
}

func TestCharacterSpellCountMerge(t *testing.T) {
	c1 := NewCharacter("1")
	c2 := NewCharacter("1")

	c1.Spells.Set("SpellsRemaining0", "1", 0)
	c1.Spells.Set("SpellCountMax0", "2", 1)

	c1.Spells.Set("SpellsRemaining1", "3", 0)
	c1.Spells.Set("SpellCountMax1", "4", 1)

	c2.Spells.MergeTo(c1)

	if val, _ := c1.Spells.Get("SpellsRemaining0"); val != "1" {
		t.Fatalf("expecting SpellsRemaining1 to be '1' %+v", c1.Spells)
	}
	if val, _ := c1.Spells.Get("SpellCountMax0"); val != "2" {
		t.Fatalf("expecting SpellCountMax1 to be '2' %+v", c1.Spells)
	}
	if val, _ := c1.Spells.Get("SpellsRemaining1"); val != "3" {
		t.Fatalf("expecting SpellsRemaining1 to be '3' %+v", c1.Spells)
	}
	if val, _ := c1.Spells.Get("SpellCountMax1"); val != "4" {
		t.Fatalf("expecting SpellCountMax1 to be '4' %+v", c1.Spells)
	}

	// empty "zero" values - not yet set
	if val, _ := c1.Spells.Get("SpellsRemaining2"); val != "" {
		t.Fatalf("expecting SpellsRemaining1 to be '0' %+v", c1.Spells)
	}
	if val, _ := c1.Spells.Get("SpellCountMax2"); val != "" {
		t.Fatalf("expecting SpellCountMax1 to be '0' %+v", c1.Spells)
	}
}
