package main

import (
	"testing"
)

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
		want := `*-------------------------------*
|        |----SYN----->|        |
| client |<--SYN/ACK---| server |
|        |----ACK----->|        |
|        |<---data---->|        |
*-------------------------------*`
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
