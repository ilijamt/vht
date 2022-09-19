package config

import (
	"strings"
	"testing"
)

func TestLoadConfig_noExist(t *testing.T) {
	config, err := LoadConfig("nope/not-once/.never")
	if err != nil {
		t.Fatal(err)
	}

	if config.TokenHelper != "" {
		t.Errorf("expected %q to be %q", config.TokenHelper, "")
	}
}

func TestParseConfig_badKeys(t *testing.T) {
	_, err := ParseConfig(`
token_helper = "/token"
nope = "true"
`)
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), `invalid key "nope" on line 3`) {
		t.Errorf("bad error: %s", err.Error())
	}
}
