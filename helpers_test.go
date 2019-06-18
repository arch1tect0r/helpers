package helpers

import (
	"testing"

	"os"
	"strconv"
)

const (
	EnvNotDefinedVar = "ENVNOTDEFINEDVAR"
)

func TestGetEnvStringWithDefaultValueReturnDefault(t *testing.T) {
	variable := EnvNotDefinedVar
	defaultValue := "default value"

	result := GetEnvStringWithDefaultValue(variable, defaultValue)

	if result != defaultValue {
		t.Errorf("environ variable %s has value: %s, want: %s", variable, result, defaultValue)
	}
}

func TestGetEnvStringWithDefaultValueReturnRealValue(t *testing.T) {
	variable := "ENVDEFINEDVAR"
	defaultValue := "default value"
	expectedValue := "not default value"

	err := os.Setenv(variable, expectedValue)
	defer os.Unsetenv(variable)

	if err != nil {
		t.Error(err)
		return
	}

	result := GetEnvStringWithDefaultValue(variable, defaultValue)

	if result != expectedValue {
		t.Errorf("environ variable %s has value: %s, want: %s", variable, result, expectedValue)
	}
}

func TestGetRequiredEnvString(t *testing.T) {
	variable := EnvNotDefinedVar

	oldFailOnError := FailOnError
	defer func() { FailOnError = oldFailOnError }()

	FailOnError = func(err error, msg string) {
		if err == nil {
			t.Errorf("No error when should be")
		}
	}

	result := GetRequiredEnvString(variable)

	if result != "" {
		t.Errorf("Got: %s, Want: ''", result)
	}
}

func TestGetEnvIntWithDefaultValueReturnDefault(t *testing.T) {
	variable := EnvNotDefinedVar
	defaultValue := 100

	result := GetEnvIntWithDefaultValue(variable, defaultValue)

	if result != defaultValue {
		t.Errorf("environ variable %s has value: %d, want: %d", variable, result, defaultValue)
	}
}

func TestGetEnvIntWithDefaultValueReturnRealValue(t *testing.T) {
	variable := EnvNotDefinedVar
	defaultValue := 100
	expectedValue := 200

	err := os.Setenv(variable, strconv.Itoa(expectedValue))
	defer os.Unsetenv(variable)

	if err != nil {
		t.Error(err)
		return
	}

	result := GetEnvIntWithDefaultValue(variable, defaultValue)

	if result != expectedValue {
		t.Errorf("environ variable %s has value: %d, want: %d", variable, result, expectedValue)
	}
}
