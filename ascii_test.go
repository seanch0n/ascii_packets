package main

import "testing"

func TestGenerateBorder(t *testing.T) {
	assertBox := func(t testing.TB, input string, want string) {
		t.Helper()
		got := generateBorder(input)
		if got != want {
			t.Errorf("got\n%s\n\nwanted\n%s", got, want)
		}
	}

	t.Run("generate 1 liner border", func(t *testing.T) {
		input := "hello"
		expected := `*-------*
| hello |
*-------*`
		assertBox(t, input, expected)
	})

	t.Run("generate 2 liner border", func(t *testing.T) {
		input := `hello
world`
		expected := `*-------*
| hello |
| world |
*-------*`

		assertBox(t, input, expected)
	})

	t.Run("generate 2 liner border with varying line lenghts", func(t *testing.T) {
		input := `hello
world is so cool`
		expected := `*------------------*
| hello            |
| world is so cool |
*------------------*`
		assertBox(t, input, expected)
	})

}

func TestGenArrow(t *testing.T) {
	assertArrow := func(t testing.TB, arrowType string, len int, text string, want string) {
		t.Helper()
		got := genArrow(arrowType, len, text)
		if got != want {
			t.Errorf("got\n%q\nwant\n%q", got, want)
		}
	}
	t.Run("right arrow", func(t *testing.T) {
		arrowType := "right"
		arrowLen := 10
		arrowText := "SYN"
		want := "---SYN--->"
		assertArrow(t, arrowType, arrowLen, arrowText, want)
	})
	t.Run("left arrow", func(t *testing.T) {
		arrowType := "left"
		arrowLen := 10
		arrowText := "SYN"
		want := "<---SYN---"
		assertArrow(t, arrowType, arrowLen, arrowText, want)
	})
	t.Run("bi arrow", func(t *testing.T) {
		arrowType := "bi"
		arrowLen := 10
		arrowText := "SYN"
		want := "<--SYN--->"
		assertArrow(t, arrowType, arrowLen, arrowText, want)
	})
	t.Run("right arrow odd length", func(t *testing.T) {
		arrowType := "right"
		arrowLen := 11
		arrowText := "SYN"
		want := "---SYN---->"
		assertArrow(t, arrowType, arrowLen, arrowText, want)
	})
}
func TestBuildWall(t *testing.T) {
	assertWall := func(t testing.TB, input string, insert bool, want string) {
		t.Helper()
		got := buildWall(input, insert)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("no insert", func(t *testing.T) {
		input := "hello there"
		insert := false
		want := "|             |"
		assertWall(t, input, insert, want)
	})
	t.Run("insert", func(t *testing.T) {
		input := "hello there"
		insert := true
		want := "| hello there |"
		assertWall(t, input, insert, want)
	})
}
