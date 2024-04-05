package oci

import (
	"errors"
	"io/ioutil"

	"path/filepath"

	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/oras"
)

// Push pushes a WASM module to an OCI registry
func Push(ref, module string, insecure, useHTTP bool, profile string, configMediaType string, contentLayerMediaType string) error {
	ctx, resolver, store := newORASContext(insecure, useHTTP)

	contents, err := ioutil.ReadFile(module)
	if err != nil {
		return err
	}

	if configMediaType != "" && contentLayerMediaType != "" {
		log.Infof("Using Config Media Type: %v", configMediaType)
		log.Infof("Using Content Layer Media Type: %v", contentLayerMediaType)
	} else {
		if profile != "" {
			value, ok := GetProfileMap()[profile]
			if ok {
				log.Infof("Using profile: %v - %v", profile, value.Description)
				configMediaType = value.ConfigMediaType
				contentLayerMediaType = value.ContentLayerMediaType
			} else {
				extension := filepath.Ext(module)
				value, ok = GetExtensionMap()[extension]
				if ok {
					log.Infof("Using extension: %v - %v", extension, value.Description)
					configMediaType = value.ConfigMediaType
					contentLayerMediaType = value.ContentLayerMediaType
				} else {
					return errors.New("cannot determine config media type and content layer media type to use")
				}
			}
		}
	}

	desc := store.Add(module, contentLayerMediaType, contents)
	layers := []ocispec.Descriptor{desc}

	pushOpts := []oras.PushOpt{
		oras.WithConfigMediaType(configMediaType),
		oras.WithNameValidation(nil),
	}

	manifest, err := oras.Push(ctx, resolver, ref, store, layers, pushOpts...)
	if err != nil {
		return err
	}

	log.Infof("Pushed: %v", ref)
	log.Infof("Size: %v", desc.Size)
	log.Infof("Digest: %v", manifest.Digest)

	return nil
}
