package beyondtest

import (
	"encoding/json"

	"imetatroll.com/character.git/importer/dndbeyond"
)

func NewBeyondCharacter(data string) (*beyond.Character, error) {
	char := &beyond.Character{}
	if err := json.Unmarshal([]byte(data), char); err != nil {
		return char, err
	}
	return char, nil
}
