package scores

import (
	"slices"
	"testing"
)

func TestCalcScale(t *testing.T) {
	wants := []Scale{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048}
	for i, note := range allNotes {
		t.Run("scale with single note: "+note, func(t *testing.T) {
			got := CalcScale([]Note{note})
			want := wants[i]

			assertCorrectScale(t, got, want)
		})
	}

	wants = []Scale{1, 3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095}
	for i, note := range allNotes {
		t.Run("scale with single notes from C to "+note, func(t *testing.T) {
			got := CalcScale(allNotes[0 : i+1])
			want := wants[i]

			assertCorrectScale(t, got, want)
		})
	}
}

func TestGetNotes(t *testing.T) {
	t.Run("scale with single note: C", func(t *testing.T) {
		got := GetNotes(1)
		want := []Note{"C"}

		assertCorrectNotes(t, got, want)
	})
	t.Run("scale with multiple notes: C, E, F#", func(t *testing.T) {
		got := GetNotes(81)
		want := []Note{"C", "E", "F#"}

		assertCorrectNotes(t, got, want)
	})
}

func assertCorrectScale(t testing.TB, got, want Scale) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertCorrectNotes(t testing.TB, got, want []Note) {
	t.Helper()
	for _, note := range want {
		if !slices.Contains(got, note) {
			t.Errorf("got %q want %q", got, want)
		}
	}

}
