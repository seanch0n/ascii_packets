package ascii_packets

import (
	"bytes"
	"reflect"
	"testing"
)

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
		t.Helper()
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

func TestReadDataFile(t *testing.T) {
	assertRead := func(t testing.TB, input string, want []string) {
		t.Helper()
		var buffer bytes.Buffer
		buffer.WriteString(input)
		content, err := readFile(&buffer)
		if err != nil {
			t.Errorf("failed to read file")
		}
		if reflect.DeepEqual(content, want) != true {
			t.Errorf("got %v want %v", content, want)
		}
	}
	t.Run("one liner with multiple words", func(t *testing.T) {
		input := "fake data neato-burrito"
		want := []string{"fake data neato-burrito"}
		assertRead(t, input, want)
	})

	t.Run("one word one line", func(t *testing.T) {
		input := "word"
		want := []string{"word"}
		assertRead(t, input, want)
	})
	t.Run("two lines of data", func(t *testing.T) {
		input := `one word
two lines`
		want := []string{"one word", "two lines"}
		assertRead(t, input, want)
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

func TestParseMsg(t *testing.T) {
	assertParse := func(t testing.TB, input string, want string, wantErr bool) {
		t.Helper()
		got, err := parseMsg(input)
		if (err != nil) != wantErr {
			t.Errorf("got an error: %s", err)
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("SYN", func(t *testing.T) {
		input := "client-conn-server: SYN"
		want := "SYN"
		assertParse(t, input, want, false)
	})
	t.Run("error, no ':'", func(t *testing.T) {
		input := "client-conn-server SYN"
		want := ""
		assertParse(t, input, want, true)
	})
	t.Run("too many ':'", func(t *testing.T) {
		input := "client:conn:server: SYN"
		want := ""
		assertParse(t, input, want, true)
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

func TestBuildSequence(t *testing.T) {
	assertSequence := func(t testing.TB, input []string, leftNode string, rightNode string, want string) {
		t.Helper()
		got := buildSequence(input, leftNode, rightNode)
		if got != want {
			t.Errorf("got\n%q\nwant\n%q", got, want)
		}
	}
	t.Run("basic", func(t *testing.T) {
		input := []string{
			"client-server: SYN",
			"server-client: SYN/ACK",
			"client-server: ACK",
			"server=client: data",
		}
		leftNode := "client"
		rightNode := "server"
		want := `*----------------------------------*
|        |-----SYN------>|        |
| client |<---SYN/ACK----| server |
|        |-----ACK------>|        |
|        |<----data----->|        |
*----------------------------------*`
		assertSequence(t, input, leftNode, rightNode, want)
	})
}

func TestFindDirection(t *testing.T) {
	assertDirection := func(t testing.TB, input string, leftNode string, rightNode string, want string, wantErr bool) {
		t.Helper()
		got, err := findDirection(input, leftNode, rightNode)
		if (err != nil) != wantErr {
			t.Errorf("got an error, didn't want one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("left to right arrow", func(t *testing.T) {
		input := "client-server: SYN"
		leftNode := "client"
		rightNode := "server"
		want := "right"
		assertDirection(t, input, leftNode, rightNode, want, false)
	})
	t.Run("right to left arrow", func(t *testing.T) {
		input := "server-client: SYN"
		leftNode := "client"
		rightNode := "server"
		want := "left"
		assertDirection(t, input, leftNode, rightNode, want, false)
	})
	t.Run("bidirectional", func(t *testing.T) {
		input := "server=client: SYN"
		leftNode := "client"
		rightNode := "server"
		want := "bi"
		assertDirection(t, input, leftNode, rightNode, want, false)
	})
	t.Run("node not found", func(t *testing.T) {
		input := "server-client: SYN"
		leftNode := "imaleftnode"
		rightNode := "imarightnode"
		want := ""
		assertDirection(t, input, leftNode, rightNode, want, true)
	})
	t.Run("no nodes found", func(t *testing.T) {
		input := "server client: SYN"
		leftNode := "client"
		rightNode := "server"
		want := ""
		assertDirection(t, input, leftNode, rightNode, want, true)
	})
}
