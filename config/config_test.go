package config

import (
	"testing"
	"time"
)

func Test_LoadConfig(t *testing.T) {
	test := []struct {
		name        string
		env         string
		expectedErr bool
		expectedCfg Config
	}{
		{
			name:        "return local config",
			env:         "local",
			expectedErr: false,
			expectedCfg: Config{
				MicroserviceName: "user-auth-api",
				Env:              "local",
				HTTP: HTTPConfig{
					Port:          "8080",
					GracefullTime: 5 * time.Second,
				},
			},
		},
		{
			name:        "return stg config",
			env:         "stg",
			expectedErr: false,
			expectedCfg: Config{
				MicroserviceName: "user-auth-api",
				Env:              "stg",
				HTTP: HTTPConfig{
					Port:          "",
					GracefullTime: 60 * time.Second,
				},
			},
		},
		{
			name:        "return error when env not found",
			env:         "unknown",
			expectedErr: true,
			expectedCfg: Config{},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load(tt.env)
			if tt.expectedErr && err == nil {
				t.Errorf("expected error, got nil")
			}
			if !tt.expectedErr && err != nil {
				t.Errorf("expected nil, got error")
			}
			if cfg != tt.expectedCfg {
				t.Errorf("expected %v, got %v", tt.expectedCfg, cfg)
			}
		})
	}

}
