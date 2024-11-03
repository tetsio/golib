package types9_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tt.want, tt.s.String())
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
			assert.Equal(t, tt.want, tt.s.Value())
		})
	}
}

func TestSecretString_MarshalText(t *testing.T) {
	tests := []struct {
		name    string
		s       types9.SecretString
		want    []byte
		wantErr bool
	}{
		{
			name:    "simple marshal text call",
			s:       types9.SecretString(testValue),
			want:    []byte(redacted),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.MarshalText()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSecretString_JSON(t *testing.T) {
	type TestStructure struct {
		SecretString types9.SecretString `json:"secret_string"`
	}
	var wantedOutput = []byte(`{"secret_string":"**HIDDEN**"}`)
	testStructure := TestStructure{SecretString: types9.SecretString(testValue)}
	got, err := json.Marshal(testStructure)
	assert.NoError(t, err)
	assert.Equal(t, wantedOutput, got)
}
