package ascii_packets

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

func TestGetMaxLengthString(t *testing.T) {
	assertLength := func(t testing.TB, input []string, want int) {
		t.Helper()
		got := getMaxLengthString(input)
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("length with 2 strings", func(t *testing.T) {
		input := []string{"a", "aa"}
		expected := 2
		assertLength(t, input, expected)
	})
	t.Run("length with 1 string", func(t *testing.T) {
		input := []string{"aaaa"}
		expected := 4
		assertLength(t, input, expected)
	})
	t.Run("length with 4 strings", func(t *testing.T) {
		input := []string{"a", "aa", "b", "ccc"}
		expected := 3
		assertLength(t, input, expected)
	})
	t.Run("length with duplicate strings", func(t *testing.T) {
		input := []string{"a", "aa", "a", "bbb"}
		expected := 3
		assertLength(t, input, expected)
	})
}

func TestGenStringOfLen(t *testing.T) {
	assertStringOfLen := func(t testing.TB, c string, len int, want string) {
		got := genStringOfLen(c, len)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("5 of character -", func(t *testing.T) {
		c := "-"
		len := 5
		expected := "-----"
		assertStringOfLen(t, c, len, expected)
	})
	t.Run("1 of character -", func(t *testing.T) {
		c := "-"
		len := 1
		expected := "-"
		assertStringOfLen(t, c, len, expected)
	})
	t.Run("0 of character -", func(t *testing.T) {
		c := "-"
		len := 0
		expected := ""
		assertStringOfLen(t, c, len, expected)
	})
	t.Run("5 of spaces", func(t *testing.T) {
		c := " "
		len := 5
		expected := "     "
		assertStringOfLen(t, c, len, expected)
	})
	t.Run("2 of A", func(t *testing.T) {
		c := "A"
		len := 2
		expected := "AA"
		assertStringOfLen(t, c, len, expected)
	})

}
