package enviroment

import (
	"os"
	"testing"
)

func Test_getEnvAsInWithDefault(t *testing.T) {
	test := []struct {
		name         string
		envKey       string
		envValue     string
		defaultValue int
		expected     int
		shouldPanic  bool
	}{
		{
			name:         "return default value when env is no set",
			envKey:       "NON_EXISTENT_ENV",
			envValue:     "",
			defaultValue: 10,
			expected:     10,
			shouldPanic:  false,
		},
		{
			name:         "return env value when set",
			envKey:       "EXISTENT_ENV",
			envValue:     "100",
			defaultValue: 42,
			expected:     100,
			shouldPanic:  false,
		},
		{
			name:         "panics when env values is not int",
			envKey:       "INVALID_INT_ENV",
			envValue:     "invalid",
			defaultValue: 42,
			expected:     0,
			shouldPanic:  true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.envValue != "" {
				os.Setenv(tt.envKey, tt.envValue)
				defer os.Unsetenv(tt.envKey)
			}

			if tt.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			result := GetEnvAsIntWithDefault(tt.envKey, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
