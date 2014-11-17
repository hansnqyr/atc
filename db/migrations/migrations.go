package migrations

import "github.com/BurntSushi/migration"

var Migrations = []migration.Migrator{
	InitialSchema,
	MoveSourceAndMetadataToVersionedResources,
	AddTypeToVersionedResources,
	RemoveTransitionalCurrentVersions,
	NonNullableVersionInfo,
	AddOneOffNameSequence,
	AddHijackURLToBuilds,
	AddTimestampsToBuilds,
	CreateLocks,
	AddBuildEvents,
	ReplaceBuildsAbortHijackURLsWithGuidAndEndpoint,
	ReplaceBuildEventsIDWithEventID,
	AddLocks,
	DropOldLocks,
	AddConfig,
	AddNameToBuildInputs,
}
