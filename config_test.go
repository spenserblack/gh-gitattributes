package main

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name   string
		fails  bool
		reader ghConfigReader
		want   *Config
	}{
		{
			name:   "propogate error",
			fails:  true,
			reader: invalidGhConfigReader,
			want:   nil,
		},
		{
			name:   "reads from config",
			fails:  false,
			reader: newGhConfigReader(`gh_gitattributes_source: foo/bar`),
			want: &Config{
				Source: "foo/bar",
			},
		},
		{
			name:   "uses default value",
			fails:  false,
			reader: newGhConfigReader(``),
			want: &Config{
				Source: defaultConfig.Source,
			},
		},
	}

	for _, tt := range tests {
		cfg, err := newConfig(tt.reader)
		if tt.fails {
			if err == nil {
				t.Errorf("Expected error")
			}
			continue
		}
		if err != nil {
			t.Fatalf(`err = %v, want nil`, err)
		}
		if !reflect.DeepEqual(cfg, tt.want) {
			t.Errorf(`cfg = %v, want %v`, cfg, tt.want)
		}
	}
}
