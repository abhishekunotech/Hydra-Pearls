package utils

import (
"net/http"
"github.com/abhishekunotech/Hydra-Pearls/utils/logger"
)

// Wrapper Function to logger/logger.go
func Logger(handler http.Handler, name string) http.Handler {
	return (logger.Logger(handler, name))
}
