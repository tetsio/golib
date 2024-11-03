package types9

type SecretString string

const (
	secretStringRedacted = "**HIDDEN**"
)

func (s SecretString) String() string {
	return secretStringRedacted
}

func (s SecretString) Value() string {
	return string(s)
}

func (s SecretString) MarshalText() ([]byte, error) {
	return []byte(secretStringRedacted), nil
}
