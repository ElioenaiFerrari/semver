package semver

import (
	"errors"
	"strconv"
	"strings"
)

var InvalidVersionErr = errors.New("invalid version in list")

func Parse(version string) (int, error) {
	version = strings.ReplaceAll(version, ".", "")

	parsedVersion, err := strconv.Atoi(version)
	if err != nil {
		return 0, InvalidVersionErr
	}

	return parsedVersion, nil
}

func GetLessVersion(versions ...string) (string, error) {
	lessVersion := "100.100.100"
	for _, version := range versions {
		parsedLessVersion, _ := Parse(lessVersion)
		parsedVersion, err := Parse(version)
		if err != nil {
			return "", err
		}

		if parsedVersion < parsedLessVersion {
			lessVersion = version
		}

	}

	return lessVersion, nil
}

func GetLatestVersion(versions ...string) (string, error) {
	lessVersion := "0.0.0"
	for _, version := range versions {
		parsedLessVersion, _ := Parse(lessVersion)
		parsedVersion, err := Parse(version)
		if err != nil {
			return "", err
		}

		if parsedVersion > parsedLessVersion {
			lessVersion = version
		}

	}

	return lessVersion, nil
}
