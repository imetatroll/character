package base

import (
	"reflect"
	"strconv"
	"strings"
)

type CharacterField struct {
	Val string
	TS  int64
}

func (cf *CharacterField) Merge(val string, ts int64) {
	if ts >= cf.TS {
		cf.Val = val
		cf.TS = ts
	}
}

// ByLevel is a slice that will be ordered by the spell level.
type ByLevel []OrderedLevel

// At returns the given ordered object by index.
func (a ByLevel) At(i int) OrderedLevel { return a[i] }

func (a ByLevel) Len() int           { return len(a) }
func (a ByLevel) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLevel) Less(i, j int) bool { return a[i].Level < a[j].Level }

// OrderedLevel gives an array ordered by spell levels.
type OrderedLevel struct {
	Index string
	Name  string
	Level int
}

// Orderable is something that can be ordered.
type Orderable interface {
	At(i int) OrderedLevel
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

// CharacterSaver is a tab in the character sheet that can have values that are set, get, removed, and saved.
// This includes the top page, weapon, spell, and item pages (etc).
type CharacterSaver interface {
	// Takes an html id like "Weapons.Name.0" and returns the internally stored Name value of the first element of the Weapons array
	Get(id string) (string, int64)
	Set(id string, value string, timestamp int64)
	Remove(property, index string)
	Ordered(string) Orderable
	CopyTo(char *Character)  // Copy the object as is into Character
	MergeTo(char *Character) // Compare individual entry timestamps to determine what to copy
	Length(property string) int
}

// GetPath takes period separated terms and uses them to find a reflected property in an object
// Id values on the page will be period separated values that point to a property of an element in an array
// Html-ID: "Property1.Property2.Index" -> Class.Property1[index].Property2 where class is inferred from Property1.
// EG
// Items.Name.0 -> Character.Items[0].Name
// Spells.Name.0 -> Character.Spells[0].Name
func GetPath(page interface{}, htmlID, expectedProperty string) (string, int64) {
	path := strings.Split(htmlID, ".")
	if len(path) != 3 {
		println("Malformed", expectedProperty, "path", path)
		return "", 0
	}
	if path[0] == expectedProperty {
		index, err := strconv.Atoi(path[2])
		if err != nil {
			println("Malformed", expectedProperty, "path[2]", path, "need Index (int)")
			return "", 0
		}
		ps := reflect.ValueOf(page)
		s := ps.Elem()
		if s.Kind() == reflect.Struct {
			field := s.FieldByName(path[0])
			if field.IsValid() {
				// field should now be the array property
				// EG Items in CharacterItems.Items, Spells in CharacterSpells.Spells etc
				if index >= 0 && index < field.Len() {
					// the field is an array (EG Items) and its internal field can be found and the value returned.
					characterField := field.Index(index).FieldByName(path[1])
					if characterField.IsValid() && characterField.Kind() == reflect.Struct {
						value := characterField.FieldByName("Val")
						ts := characterField.FieldByName("TS")
						return value.String(), ts.Int()
					}
					println(htmlID, "invalid field error")
				} else {
					println("GetPath", index, "is invalid; ", htmlID, "in", expectedProperty, "does not exist")
				}
			} else {
				println(path[0], "does not appear to be a property", htmlID, expectedProperty)
			}
		}
	}
	return "", 0
}

// Get immediately returns a value from a reflected property
func Get(page interface{}, id, pageName string) (string, int64) {
	ps := reflect.ValueOf(page)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		characterField := s.FieldByName(id)
		if characterField.IsValid() && characterField.Kind() == reflect.Struct {
			value := characterField.FieldByName("Val")
			ts := characterField.FieldByName("TS")
			if value.IsValid() && ts.IsValid() {
				return value.String(), ts.Int()
			}
			panic("invalid field")
		}
		panic(id + " does not appear to be a valid " + pageName + " field")
	}
	return "", 0
}

// Set immediately sets the value in a reflected property
func Set(page interface{}, id, val string, timestamp int64, pageName string) {
	ps := reflect.ValueOf(page)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		characterField := s.FieldByName(id)
		if characterField.IsValid() && characterField.Kind() == reflect.Struct {
			// the character sheet field consists of a string + timestamp value
			value := characterField.FieldByName("Val")
			if value.CanSet() && value.Kind() == reflect.String {
				value.SetString(val)
			}
			ts := characterField.FieldByName("TS")
			if ts.CanSet() && ts.Kind() == reflect.Int64 {
				ts.SetInt(timestamp)
			}
		} else {
			println(id, "does not appear to be a valid", pageName, " field")
		}
	}
}

// HTMLIdToStructPath validates (to some extent) and returns a path for consumption by SetPath
func HTMLIdToStructPath(htmlID, expectedProperty string) ([]string, int, bool) {
	path := strings.Split(htmlID, ".")
	if len(path) != 3 {
		println("Malformed", expectedProperty, "path", path, "from", htmlID)
		return path, -1, false
	}
	if path[0] != expectedProperty {
		return path, -1, false
	}
	index, err := strconv.Atoi(path[2])
	if err != nil {
		println("Malformed", expectedProperty, "path[2]", path, "need Index (int)")
		return path, -1, false
	}
	return path, index, true
}

// SetPath takes in one of the Characters top level properties (CharacterTop, CharacterSpells, etc) and sets a value internally based on the htmlId
// Must make sure that the underlying array first gets a new value if that is required in order to handle the index request
// See GetPath description
func SetPath(inter interface{}, path []string, index int, val string, timestamp int64) {
	ps := reflect.ValueOf(inter)
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		field := s.FieldByName(path[0])
		if field.IsValid() {
			// field should now be the array property
			// EG Items in CharacterItems.Items, Spells in CharacterSpells.Spells etc
			if index >= 0 && index < field.Len() {
				// the field is an array (EG Items) and its internal field can be found and the value returned.
				internalField := field.Index(index).FieldByName(path[1])
				if internalField.IsValid() {
					if internalField.CanSet() && internalField.Kind() == reflect.Struct {
						value := internalField.FieldByName("Val")
						value.SetString(val)
						ts := internalField.FieldByName("TS")
						ts.SetInt(timestamp)
					} else {
						println(path, "cannot set field error")
					}
				} else {
					println(path, "invalid field error")
				}
			} else {
				println(index, "does not appear to be a valid index")
			}
		} else {
			println(path[0], "does not appear to be a property")
		}
	}
}

func FullMerge(from, to *Character, updatedAt int64) {
	to.UpdatedAt = updatedAt

	to.IsShared = from.IsShared
	to.IsSharedWithUserID = from.IsSharedWithUserID

	Merge(from.Top, to.Top)
	Merge(from.Bio, to.Bio)
	Merge(from.Combat, to.Combat)
	Merge(from.Spells, to.Spells)
	Merge(from.Items, to.Items)
}

func MergeCharacterField(fromField CharacterField, toField reflect.Value) {
	toTS := toField.FieldByName("TS")
	if toTS.CanSet() && toTS.Kind() == reflect.Int64 {
		if toTS.Int() < fromField.TS {
			toValue := toField.FieldByName("Val")
			if toValue.CanSet() && toValue.Kind() == reflect.String {
				toValue.SetString(fromField.Val)
				toTS.SetInt(fromField.TS)
			}
		}
	}
}

// Merge the from character tab into the to character tab
func Merge(from, to CharacterSaver) {
	fromStruct := reflect.ValueOf(from).Elem()
	for structIndex := 0; structIndex < fromStruct.NumField(); structIndex++ {
		switch fromStruct.Field(structIndex).Interface().(type) {
		case CharacterField:
			fromName := fromStruct.Type().Field(structIndex).Name

			// find the named field in the "to" struct and then set if the timestamp is old
			fromField := fromStruct.Field(structIndex).Interface().(CharacterField)
			toField := reflect.ValueOf(to).Elem().FieldByName(fromName)
			MergeCharacterField(fromField, toField)

		default:
			// armor, weapons, spells, items
			// iterate the from array object and search for the matching uuid in the to object
			// - if found, replace granular values based on timestamp
			// - if not found, add a new item
			// - all items in the to object that are not matched must also be removed
			fromUUIDS := map[string]bool{}
			fromArray := fromStruct.Field(structIndex)

			toStruct := reflect.ValueOf(to).Elem()
			toArray := toStruct.Field(structIndex)

			// handle array and return
			if toArray.Type().Kind() == reflect.Array {
				for fromIndex := 0; fromIndex < fromArray.Len(); fromIndex++ {
					fromEntry := fromArray.Index(fromIndex)

					for i := 0; i < fromEntry.NumField(); i++ {
						// fmt.Printf("HERE %+v %+v\n", fromEntry.Type().Field(i).Name, fromEntry.Field(i))
						fromField := fromEntry.Field(i).Interface().(CharacterField)
						toField := toArray.Index(fromIndex).Field(i)
						MergeCharacterField(fromField, toField)
					}
				}
				continue
			}

			// additions
			for fromIndex := 0; fromIndex < fromArray.Len(); fromIndex++ {
				// println(fromArray.Index(fromIndex).Type().Field(0).Name)

				uuid := fromArray.Index(fromIndex).FieldByName("UUID").FieldByName("Val").String()
				fromUUIDS[uuid] = true
				//fmt.Printf("%+v %+v\n", fromArrayEntry.Index(i).Kind(), fromArrayEntry.Index(i))

				found := false
				for toIndex := 0; toIndex < toArray.Len(); toIndex++ {
					toUUID := toArray.Index(toIndex).FieldByName("UUID").FieldByName("Val").String()

					if uuid == toUUID {
						// merge the matching structs
						found = true
						fromEntry := fromArray.Index(fromIndex)

						for i := 0; i < fromEntry.NumField(); i++ {
							// fmt.Printf("HERE %+v %+v\n", fromEntry.Type().Field(i).Name, fromEntry.Field(i))
							fromField := fromEntry.Field(i).Interface().(CharacterField)
							toField := toArray.Index(toIndex).Field(i)
							MergeCharacterField(fromField, toField)
						}
					}
				}
				// add this new object to the end of the array
				if !found {
					toArray.Set(reflect.Append(toArray, fromArray.Index(fromIndex)))
				}
			}

			// if empty, just return
			if toArray.Len() == 0 {
				continue
			}

			// remove deleted items
			toType := toArray.Index(0).Type()
			//fmt.Printf("TYPE %+v\n", toArray.Type())
			//fmt.Printf("KIND %+v\n", toArray.Type().Kind())

			// fixed arrays are *skipped* EG SpellCounts
			slice := reflect.MakeSlice(reflect.SliceOf(toType), 0, toArray.Len())
			for arrayIndex := 0; arrayIndex < toArray.Len(); arrayIndex++ {
				toUUID := toArray.Index(arrayIndex).FieldByName("UUID").FieldByName("Val").String()
				if _, ok := fromUUIDS[toUUID]; ok {
					slice = reflect.Append(slice, toArray.Index(arrayIndex))
				}
			}
			toArray.Set(slice)
		}
	}
}
