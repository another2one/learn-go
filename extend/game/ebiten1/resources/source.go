package resources

import "embed"

////go:embed assets/config.json
//var ConfigByte []byte

//go:embed assets
var EmbedPath embed.FS
