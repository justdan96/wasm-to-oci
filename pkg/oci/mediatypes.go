package oci

/*
const PROVIDER_ARCHIVE_MEDIA_TYPE: &str = "application/vnd.wasmcloud.provider.archive.layer.v1+par";
const PROVIDER_ARCHIVE_CONFIG_MEDIA_TYPE: &str =
    "application/vnd.wasmcloud.provider.archive.config";
const PROVIDER_ARCHIVE_FILE_EXTENSION: &str = ".par.gz";
const WASM_MEDIA_TYPE: &str = "application/vnd.module.wasm.content.layer.v1+wasm";
const WASM_CONFIG_MEDIA_TYPE: &str = "application/vnd.wasmcloud.actor.archive.config";
const OCI_MEDIA_TYPE: &str = "application/vnd.oci.image.layer.v1.tar";
const WASM_FILE_EXTENSION: &str = ".wasm";
*/

type MediaTypes struct {
	Description           string
	ConfigMediaType       string
	ContentLayerMediaType string
}

func GetProfileMap() map[string]MediaTypes {
	return map[string]MediaTypes{
		"wc-actor": {
			Description:           "Actor configuration used by WasmCloud",
			ConfigMediaType:       "application/vnd.wasmcloud.actor.archive.config",
			ContentLayerMediaType: "application/vnd.module.wasm.content.layer.v1+wasm",
		},
		"wc-provider": {
			Description:           "Provider configuration used by WasmCloud",
			ConfigMediaType:       "application/vnd.wasmcloud.provider.archive.config",
			ContentLayerMediaType: "application/vnd.wasmcloud.provider.archive.layer.v1+par",
		},
	}
}

func GetExtensionMap() map[string]MediaTypes {
	return map[string]MediaTypes{
		"wasm": {
			Description:           "Generic WebAssembly module configuration",
			ConfigMediaType:       "application/vnd.wasm.config.v1+json",
			ContentLayerMediaType: "application/vnd.wasm.content.layer.v1+wasm",
		},
	}
}
