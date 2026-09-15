package swagger

import (
	"embed"
	_ "embed"
)

// FS embeds all swagger JSON files in this directory.
//
//go:embed content/*
var FS embed.FS

// ContentServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for ContentService (ROM & Image).
//
//go:embed content/service.swagger.json
var ContentServiceSwaggerJSON []byte

// ContentSwaggerJSON is an alias for ContentServiceSwaggerJSON.
var ContentSwaggerJSON = ContentServiceSwaggerJSON

// ContentRomServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for RomService.
//
//go:embed content/v1/rom/romService.swagger.json
var ContentRomServiceSwaggerJSON []byte

// ContentImageServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for ImageService.
//
//go:embed content/v1/image/imageService.swagger.json
var ContentImageServiceSwaggerJSON []byte

// IdentityServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for IdentityService.
//
//go:embed identity/v1/service.swagger.json
var IdentityServiceSwaggerJSON []byte

// ManagementServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for ManagementService (Console & Game).
//
//go:embed management/service.swagger.json
var ManagementServiceSwaggerJSON []byte

// ManagementSwaggerJSON is an alias for ManagementServiceSwaggerJSON.
var ManagementSwaggerJSON = ManagementServiceSwaggerJSON

// ManagementConsoleServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for ConsoleService.
//
//go:embed management/v1/console/consoleService.swagger.json
var ManagementConsoleServiceSwaggerJSON []byte

// ManagementGameServiceSwaggerJSON contains the raw OpenAPI 2.0 JSON for GameService.
//
//go:embed management/v1/game/gameService.swagger.json
var ManagementGameServiceSwaggerJSON []byte
