## Generating a new character.go file

The character.go file is used when moving the dnd beyond character sheet into
the imetatroll.com 5ESRD system.

### Go to your character in DND Beyond

- Login to dnd beyond.
- Fill out everything that you possibly can.
- Go to the characters page.
- Select your character.

The url will look like this:
```
https://character-service.dndbeyond.com/character/v3/character/xxxxxxxx
```
Replace xxxxxxxx with your character id from
```
https://www.dndbeyond.com/profile/username/characters/id
```

###


### Save the page as a json document

- Right click in the browser and save as a json document.

### Pretty Print Your Character Sheet

- Open a terminal on your machine.
- Reformat the character sheet in order to more easily inspect the values.

```
python3 -m json.tool < character.json > pretty_printed.json
```

### Remove unused elements

Remove the following portions of the downloaded json file:

```
"characterConfiguration": {
  "startingEquipmentType": null,
  "abilityScoreType": 2,
  "showHelpText": false
},
"characterData": {
...
}
```

### character.go changes

- Use https://transform.tools/json-to-go to generated a new struct
- Replace character.go with the new structure.

### Preventing Regressions

- Any objects in the resulting sheet should become more specific rather than less.

EG an int -> interface{} change would be a regression.

- Rearrangements of internal structure must be dealt with by updating the methods.go file.

### Testing

Located under the folder '/test'.
