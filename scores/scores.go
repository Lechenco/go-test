package scores

import (
	"slices"
)

type Note = string

var allNotes = []Note{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

type Scale = uint16

func CalcScale(scaleNotes []Note) Scale {
	scale := 0

	for i, note := range allNotes {
		if slices.Contains(scaleNotes, note) {
			scale = scale | 1<<i
		}
	}

	return Scale(scale)
}

func GetNotes(scale Scale) []Note {
	scaleNotes := []Note{}

	for i, note := range allNotes {
		if (scale & 1 << i) > 0 {
			scaleNotes = append(scaleNotes, note)
		}
	}
	return scaleNotes
}
