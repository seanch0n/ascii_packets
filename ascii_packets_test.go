package ascii_packets

import "testing"

func TestGenerateBorder(t *testing.T) {
	t.Run("generate 1 liner border", func(t *testing.T) {
		input := "hello"
		expected := `*-------*
| hello |
*-------*`
		got := generateBorder(input)
		if got != expected {
			t.Errorf("did not get what we expected. Got\n%s\n\nwanted\n\n%s", got, expected)
		}
	})
	t.Run("generate 2 liner border", func(t *testing.T) {
		input := `hello
world`
		expected := `*-------*
| hello |
| world |
*-------*`
		got := generateBorder(input)
		if got != expected {
			t.Errorf("did not get what we expected. Got\n%s\n\nwanted\n\n%s", got, expected)
		}
	})
	t.Run("generate 2 liner border with varying line lenghts", func(t *testing.T) {
		input := `hello
world is so cool`
		expected := `*------------------*
| hello            |
| world is so cool |
*------------------*`
		got := generateBorder(input)
		if got != expected {
			t.Errorf("did not get what we expected. Got\n%s\n\nwanted\n\n%s", got, expected)
		}
	})

}

func TestGetMaxLengthString(t *testing.T) {
	t.Run("length with 2 strings", func(t *testing.T) {
		input := []string{"a", "aa"}
		expected := 2
		got := getMaxLengthString(input)
		if got != expected {
			t.Errorf("expected %d got %d", expected, got)
		}
	})
	t.Run("length with 1 string", func(t *testing.T) {
		input := []string{"aaaa"}
		expected := 4
		got := getMaxLengthString(input)
		if got != expected {
			t.Errorf("expected %d got %d", expected, got)
		}
	})
	t.Run("length with 4 strings", func(t *testing.T) {
		input := []string{"a", "aa", "b", "ccc"}
		expected := 3
		got := getMaxLengthString(input)
		if got != expected {
			t.Errorf("expected %d got %d", expected, got)
		}
	})
	t.Run("length with duplicate strings", func(t *testing.T) {
		input := []string{"a", "aa", "a", "bbb"}
		expected := 3
		got := getMaxLengthString(input)
		if got != expected {
			t.Errorf("expected %d got %d", expected, got)
		}
	})
}

func TestGenStringOfLen(t *testing.T) {
	t.Run("5 of character -", func(t *testing.T) {
		c := "-"
		len := 5
		expected := "-----"
		got := genStringOfLen(c, len)
		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	})
	t.Run("1 of character -", func(t *testing.T) {
		c := "-"
		len := 1
		expected := "-"
		got := genStringOfLen(c, len)
		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	})
	t.Run("0 of character -", func(t *testing.T) {
		c := "-"
		len := 0
		expected := ""
		got := genStringOfLen(c, len)
		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	})
	t.Run("5 of spaces", func(t *testing.T) {
		c := " "
		len := 5
		expected := "     "
		got := genStringOfLen(c, len)
		if got != expected {
			t.Errorf("got %s expected %s", got, expected)
		}
	})

}
