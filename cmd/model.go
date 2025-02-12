package cmd

import "github.com/juanMaAV92/user-auth-api/services/health"

type AppServices struct {
	HealthService health.HealthService
}
