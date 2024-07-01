package envutils

import (
	"os"
	"testing"
)

func TestMustGetString(t *testing.T) {
	key := "TEST_ENV_VAR"
	expectedValue := "test_value"

	// Set environment variable
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key) // Clean up after the test

	// Test when environment variable is set
	result := MustGetString(key)
	if result != expectedValue {
		t.Errorf("Expected %s, but got %s", expectedValue, result)
	}

	// Test when environment variable is not set
	// (Panic expected)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when environment variable is not set")
		}
	}()

	MustGetString("NON_EXISTENT_ENV_VAR")
}

func TestMustGetInt64(t *testing.T) {
	// Test case with environment variable set
	key := "TEST_VAR"
	expectedValue := int64(12345)
	os.Setenv(key, "12345")
	defer os.Unsetenv(key) // Clean up environment variable after test

	result := MustGetInt64(key)
	if result != expectedValue {
		t.Errorf("Expected %d, but got %d", expectedValue, result)
	}

	// Test case with environment variable not set
	key = "NOT_SET_VAR"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to environment variable not set, but did not get one")
		}
	}()

	MustGetInt64(key) // This should cause a panic
}

func TestMustGetInt64Invalid(t *testing.T) {
	// Test case with environment variable not set
	key := "TEST_VAR"
	os.Setenv(key, "invalid_value")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to environment variable not set, but did not get one")
		}
	}()

	MustGetInt64(key) // This should cause a panic
}

func TestMustGetInt32(t *testing.T) {
	key := "TEST_VAR"
	expectedValue := "123"

	// Set environment variable
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key) // Clean up after the test

	// Test case 1: Environment variable is set and has a valid integer value
	actualValue := MustGetInt32(key)
	if actualValue != 123 {
		t.Errorf("Expected %d, but got %d", 123, actualValue)
	}

	// Test case 2: Environment variable is not set
	// (Panic expected, so we'll use recover())
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when environment variable is not set")
		}
	}()

	MustGetInt32("UNSET_VAR")
}

func TestMustGetInt32Invalid(t *testing.T) {
	// Test case with environment variable not set
	key := "TEST_VAR"
	os.Setenv(key, "invalid_value")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to environment variable not set, but did not get one")
		}
	}()

	MustGetInt32(key) // This should cause a panic
}

func TestMustGetInt(t *testing.T) {
	key := "TEST_ENV_VAR"
	expectedValue := "123"

	// Set environment variable
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key) // Clean up after the test

	// Test case: Environment variable is set and is a valid integer
	actualValue := MustGetInt(key)
	if actualValue != 123 {
		t.Errorf("Expected value %d, but got %d", 123, actualValue)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when environment variable is not set")
		}
	}()

	MustGetInt("UNSET_VAR")
}

func TestMustGetIntInvalid(t *testing.T) {
	// Test case with environment variable not set
	key := "TEST_VAR"
	os.Setenv(key, "invalid_value")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to environment variable not set, but did not get one")
		}
	}()

	MustGetInt(key) // This should cause a panic
}

func TestMustGetBool(t *testing.T) {
	key := "TEST_VAR"

	// Test when environment variable is not set
	os.Unsetenv(key)
	defer os.Setenv(key, "true") // restore environment variable after test

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when %s is not set", key)
		}
	}()

	MustGetBool(key)
}

func TestMustGetBoolInvalidValue(t *testing.T) {
	key := "TEST_VAR"

	// Set environment variable with an invalid boolean value
	os.Setenv(key, "invalid")
	defer os.Unsetenv(key) // unset environment variable after test

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when %s has invalid value", key)
		}
	}()

	MustGetBool(key)
}

func TestMustGetBoolValidValue(t *testing.T) {
	key := "TEST_VAR"
	expectedValue := true

	// Set environment variable with a valid boolean value
	os.Setenv(key, "true")
	defer os.Unsetenv(key) // unset environment variable after test

	result := MustGetBool(key)

	if result != expectedValue {
		t.Errorf("Expected %v, but got %v", expectedValue, result)
	}
}

func TestMustGetf64(t *testing.T) {
	// Test case with environment variable set
	os.Setenv("TEST_VAR", "123.456")
	defer os.Unsetenv("TEST_VAR") // Clean up after the test

	expected := 123.456
	result := MustGetf64("TEST_VAR")

	if result != expected {
		t.Errorf("Expected %f, but got %f", expected, result)
	}

	// Test case with environment variable not set
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic due to environment variable not being set, but did not get one")
			}
		}()

		MustGetf64("UNSET_VAR")
	}()

	// Test case with environment variable set to invalid float value
	os.Setenv("INVALID_VAR", "abc")
	defer os.Unsetenv("INVALID_VAR") // Clean up after the test

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic due to invalid float value, but did not get one")
			}
		}()

		MustGetf64("INVALID_VAR")
	}()
}

func TestMustGetf32(t *testing.T) {
	// Test case with environment variable set
	os.Setenv("TEST_VAR", "123.456")
	defer os.Unsetenv("TEST_VAR") // Clean up after the test

	want := float32(123.456)
	got := MustGetf32("TEST_VAR")
	if got != want {
		t.Errorf("MustGetf32() = %v, want %v", got, want)
	}

	// Test case with environment variable not set
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("MustGetf32() did not panic when expected")
			}
		}()

		MustGetf32("UNSET_VAR")
	}()
}

func TestMustGetf32Invalid(t *testing.T) {
	// Test case with environment variable not set
	key := "TEST_VAR"
	os.Setenv(key, "invalid_value")
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic due to environment variable not set, but did not get one")
		}
	}()

	MustGetf32(key) // This should cause a panic
}
