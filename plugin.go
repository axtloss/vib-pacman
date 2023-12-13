package main

import (
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/vanilla-os/vib/api"
)

type PacmanModule struct {
	Name string `json:"name"`
	Type string `json:"type"`

	ExtraFlags []string
	Packages   []string
}

func BuildModule(moduleInterface interface{}, _ *api.Recipe) (string, error) {
	var module PacmanModule
	err := mapstructure.Decode(moduleInterface, &module)
	if err != nil {
		return "", err
	}

	cmd := fmt.Sprintf("pacman -S --noconfirm %s %s", strings.Join(module.ExtraFlags, " "), strings.Join(module.Packages, " "))

	return cmd, nil
}
