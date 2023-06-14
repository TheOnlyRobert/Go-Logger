package log

// LogCat holds data regarding a logging category.
// Logging levels (INFO, WARN, ERROR)
type LogCat struct {
	Code string
	Type string
}

var (
	// LogCatStartUp usage: service startup message
	LogCatStartUp = LogCat{Code: "STT001", Type: "service_startup"}

	// LogCatStartUp usage: service startup message
	LogCatShutDown = LogCat{Code: "STT002", Type: "service_shutdown"}

	// LogCatHealth usage: health check message
	LogCatHealth = LogCat{Code: "HTH001", Type: "health_check"}

	// LogCatDatabase usage: datastore connection message
	LogCatDatabaseConnection = LogCat{Code: "DTA001", Type: "datastore_connection"}

	// LogCatDatastoreClose usage: close datastore connection message
	LogCatDatastoreClose = LogCat{Code: "DTA002", Type: "datastore_close"}

	// LogCatDatabase usage: datastore interaction
	// Logs, e.g. query exec or record read errors
	LogCatDatabase = LogCat{Code: "DTA003", Type: "datastore_interaction"}
)
