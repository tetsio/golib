package types9_test

import (
	"encoding/json"
	"fmt"
	"github.com/tetsio/golib/types9"
	"testing"
)

const (
	testValue = "mySecretValue"
	redacted  = "**HIDDEN**"
)

func TestSecretString_String(t *testing.T) {
	tests := []struct {
		name string
		s    types9.SecretString
		want string
	}{
		{
			name: "simple stringer call",
			s:    types9.SecretString(testValue),
			want: redacted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.s.String() {
				t.Errorf("Results are not equal!\nExpected:\n%s\nActual:\n%s\n", tt.want, tt.s.String())
			}
		})
	}
}

func TestSecretString_Value(t *testing.T) {
	tests := []struct {
		name string
		s    types9.SecretString
		want string
	}{
		{
			name: "simple value call ",
			s:    types9.SecretString(testValue),
			want: testValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want != tt.s.Value() {
				t.Errorf("Results are not equal!\nExpected:\n%s\nActual:\n%s\n", tt.want, tt.s.Value())
			}
		})
	}
}

func TestSecretString_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		s       types9.SecretString
		want    string
		wantErr bool
	}{
		{
			name:    "simple marshal text call",
			s:       types9.SecretString(testValue),
			want:    redacted,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalText()
			if err != nil {
				t.Errorf("Unexpected error: %s", err.Error())
			}
			if tt.want != string(got) {
				t.Errorf("Results are not equal!\nExpected:\n%s\nActual:\n%s\n", tt.want, got)
			}
		})
	}
}

func TestSecretString_JSON(t *testing.T) {
	type TestStructure struct {
		SecretString types9.SecretString `json:"secret_string"`
	}
	var wantedOutput = fmt.Sprintf(`{"secret_string":"%s"}`, redacted)
	testStructure := TestStructure{SecretString: types9.SecretString(testValue)}
	got, err := json.Marshal(testStructure)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if wantedOutput != string(got) {
		t.Errorf("Results are not equal!\nExpected:\n%s\nActual:\n%s\n", wantedOutput, got)
	}
}
