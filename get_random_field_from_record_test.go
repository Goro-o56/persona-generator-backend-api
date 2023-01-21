package main

import "testing"

func TestGetRecord(t *testing.T) {
	record := []string{"ゲーム", "釣り", "勉強"}

	t.Run("returns field", func(t *testing.T) {
		var index = 0
		got := record[index]
		want := "ゲーム"

		assertRandomField(t, got, want, record)
	})

}
func assertRandomField(t testing.TB, got string, want string, record []string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s given, %v", got, want, record)
	}
}
