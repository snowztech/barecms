package configs

import "testing"

func TestValidateRejectsInsecureProductionSecret(t *testing.T) {
	config := AppConfig{
		Env:                    "production",
		JWTSecret:              DefaultJWTSecret,
		MaxRequestBody:         "2M",
		AuthRateLimitPerMinute: 10,
	}

	if err := config.Validate(); err == nil {
		t.Fatal("expected the default production JWT secret to be rejected")
	}
}

func TestValidateAcceptsSecureProductionConfig(t *testing.T) {
	config := AppConfig{
		Env:                    "production",
		JWTSecret:              "a-secure-secret-with-at-least-32-characters",
		MaxRequestBody:         "2M",
		AuthRateLimitPerMinute: 10,
	}

	if err := config.Validate(); err != nil {
		t.Fatalf("expected valid production config: %v", err)
	}
}

func TestValidateRejectsInvalidLimits(t *testing.T) {
	tests := []AppConfig{
		{Env: "dev", MaxRequestBody: "", AuthRateLimitPerMinute: 10},
		{Env: "dev", MaxRequestBody: "2M", AuthRateLimitPerMinute: 0},
	}

	for _, config := range tests {
		if err := config.Validate(); err == nil {
			t.Fatalf("expected invalid limits to be rejected: %+v", config)
		}
	}
}
