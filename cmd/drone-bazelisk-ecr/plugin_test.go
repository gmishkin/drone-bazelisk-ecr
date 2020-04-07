package main

import (
	"os"
	"testing"
)

func TestPlugin(t *testing.T) {
	tests := []struct {
		env  map[string]string
		want plugin
		fail bool
	}{
		// test setting struct fields
		{
			env: map[string]string{
				"PLUGIN_TAG":        "tag",
				"PLUGIN_TARGET":     "target",
				"PLUGIN_REGISTRY":   "registry",
				"PLUGIN_REPOSITORY": "repository",
				"PLUGIN_ACCESS_KEY": "access",
				"PLUGIN_SECRET_KEY": "secret",
			},
			want: plugin{
				Tag:        "tag",
				Target:     "target",
				Registry:   "registry",
				Repository: "repository",
				AccessKey:  "access",
				SecretKey:  "secret",
			},
			fail: false,
		},
		// test empty environment
		{
			env:  map[string]string{},
			want: plugin{},
			fail: true,
		},
	}

	for _, test := range tests {
		setEnvMap(test.env)

		// check desired output
		got, err := newPlugin()
		if err != nil && !test.fail {
			t.Errorf(err.Error())
		}

		if test.want != got {
			t.Errorf("%v is not equal to %v", test.want, got)
		}

		unsetEnvMap(test.env)
	}
}

// set environment
func setEnvMap(env map[string]string) {
	for key, val := range env {
		os.Setenv(key, val)
	}
}

// unset environment
func unsetEnvMap(env map[string]string) {
	for key := range env {
		os.Unsetenv(key)
	}
}
