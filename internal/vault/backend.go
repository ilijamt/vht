package vault

import (
	"errors"
	"github.com/hashicorp/vault/api"
	"strings"
)

func MountPointFromPath(path string) (mountPoint string, err error) {
	if path == "" {
		return mountPoint, errors.New(ErrMissingPath)
	}
	parts := strings.Split(path, "/")
	mountPoint = parts[0]
	return
}

func IsKV(mountPoint string, client *api.Client) (yes bool, err error) {

	if client == nil {
		return false, errors.New(ErrMissingVaultClient)
	}

	if mountPoint == "" {
		return false, errors.New(ErrMissingPath)
	}

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		return false, err
	}

	for path, mount := range mounts {
		if path == mountPoint {
			yes = mount.Type == "kv"
			return
		}
	}

	return
}

func IsV2(mountPoint string, client *api.Client) (yes bool, err error) {
	if client == nil {
		return false, errors.New(ErrMissingVaultClient)
	}

	if mountPoint == "" {
		return false, errors.New(ErrMissingPath)
	}

	mounts, err := client.Sys().ListMounts()
	if err != nil {
		return false, err
	}

	for path, mount := range mounts {
		if path == mountPoint {
			yes = mount.Options["version"] == "2"
			return
		}
	}

	return
}
