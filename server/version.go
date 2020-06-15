package server

// Protocol versions
const (
	versionLatest = 1002001 // MAJOR * 1000000 + MINOR * 1000 + PATCH
	versionLegacy = 1001001
)

const (
	versionCallResourceResponse              = 1002000
	versionSoftResourceReferenceAndDataValue = versionLatest
)
