package logger

// type 'BuildInfo' represents build information of the application
type BuildInfo struct {
	BuildVersion string
	BuildCommit  string
	BuildTime    string
}

// function 'NewBuildInfo' creates a new build info instance,
// it takes version key, commit key, and time environment variable key as arguments,
// if the environment variable is not set, it returns the key name itself,
// it returns a new instance of 'BuildInfo' with the current build information
func NewBuildInfo(versionKey, commitKey, timeKey string) BuildInfo {
	return BuildInfo{
		BuildVersion: getEnvOrKey(versionKey),
		BuildCommit:  getEnvOrKey(commitKey),
		BuildTime:    getEnvOrKey(timeKey),
	}
}

// function 'Fields' returns the build info as a slice of 'Field' for logging
func (s BuildInfo) Fields() []Field {
	return []Field{
		{"build_version", s.BuildVersion},
		{"build_commit", s.BuildCommit},
		{"build_time", s.BuildTime},
	}
}
