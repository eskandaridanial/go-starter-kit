package logger

type BuildInfo struct {
	BuildVersion string
	BuildCommit  string
	BuildTime    string
}

func NewBuildInfo(versionKey, commitKey, timeKey string) BuildInfo {
	return BuildInfo{
		BuildVersion: getEnvOrKey(versionKey),
		BuildCommit:  getEnvOrKey(commitKey),
		BuildTime:    getEnvOrKey(timeKey),
	}
}

func (s BuildInfo) Fields() []Field {
	return []Field{
		{"build_version", s.BuildVersion},
		{"build_commit", s.BuildCommit},
		{"build_time", s.BuildTime},
	}
}
