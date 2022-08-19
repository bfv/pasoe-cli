package model

type GitRelease struct {
	Name              string `json:"name"`
	ReleasedAt        string `json:"published_at"`
	LinuxArchiveURL   string `json:"tarball_url"`
	WindowsArchiveURL string `json:"zipball_url"`
}

type GitReleases []GitRelease
