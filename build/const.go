package build

var (
	version  string
	datetime string
)

// Version returns application version string
func Version() string {
	return version
}

// DateTime returns build date and time as string
func DateTime() string {
	return datetime
}
