package main

import (
	"testing"

	"github.com/vanilla-os/vib/api"
)

type testModule struct {
	Name       string
	Type       string
	ExtraFlags []string
	Packages   []string
}

type testCases struct {
	module   interface{}
	expected string
}

var test = []testCases{
	{testModule{"Single Package, Single Flag", "pacman", []string{"--overwrite=\"*\""}, []string{"bash"}}, "pacman -S --noconfirm --overwrite=\"*\" bash"},
	{testModule{"Single Package, No Flag", "pacman", []string{""}, []string{"bash"}}, "pacman -S --noconfirm  bash"},
	{testModule{"Multiple Packages, No Flag", "pacman", []string{""}, []string{"bash", "fish"}}, "pacman -S --noconfirm  bash fish"},
	{testModule{"Multiple Packages, Multiple Flags", "pacman", []string{"--overwrite=\"*\"", "--verbose"}, []string{"bash", "fish"}}, "pacman -S --noconfirm --overwrite=\"*\" --verbose bash fish"},
}

func TestBuildModule(t *testing.T) {
	for _, testCase := range test {
		if output, _ := BuildModule(testCase.module, &api.Recipe{}); output != testCase.expected {
			t.Errorf("Output %q not equivalent to expected %q", output, testCase.expected)
		}
	}

}
