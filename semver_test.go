package semver_test

import (
	"errors"
	"testing"

	"github.com/ElioenaiFerrari/semver"
)

func TestParse(t *testing.T) {
	// When
	parsedVersion1, err := semver.Parse("2.2.0")
	if err != nil {
		t.Error("unexpected error")
	}
	parsedVersion2, err := semver.Parse("100")
	if err != nil {
		t.Error("unexpected error")
	}

	// Then
	if parsedVersion1 != 220 {
		t.Errorf("expected version=%d, got=%d", 220, parsedVersion1)
	}

	if parsedVersion2 != 100 {
		t.Errorf("expected version=%d, got=%d", 100, parsedVersion1)
	}

	if _, err := semver.Parse("x"); !errors.Is(err, semver.InvalidVersionErr) {
		t.Error("unexpected error")
	}
}

type Test struct {
	expectedVersion string
	expectedErr     error
	handler         func(versions ...string) (string, error)
	title           string
	versions        []string
}

var tests = []Test{
	{
		expectedVersion: "0.0.1",
		expectedErr:     nil,
		handler:         semver.GetLessVersion,
		title:           "when given valid versions, should returns less version",
		versions:        []string{"1.1.0", "0.1.1", "0.0.1", "0.0.2"},
	},
	{
		expectedVersion: "1.1.0",
		expectedErr:     nil,
		handler:         semver.GetLatestVersion,
		title:           "when given valid versions, should returns latest version",
		versions:        []string{"1.1.0", "0.1.1", "0.0.1", "0.0.2"},
	},
	{
		expectedVersion: "",
		expectedErr:     semver.InvalidVersionErr,
		handler:         semver.GetLessVersion,
		title:           "when given invalid version in less fn, should returns err",
		versions:        []string{"x", "0.1.1", "0.0.1", "0.0.2"},
	},
	{
		expectedVersion: "",
		expectedErr:     semver.InvalidVersionErr,
		handler:         semver.GetLatestVersion,
		title:           "when given invalid version in latest fn, should returns err",
		versions:        []string{"x", "0.1.1", "0.0.1", "0.0.2"},
	},
}

func TestGetLessAndLatestVersions(t *testing.T) {
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			// When
			version, err := test.handler(test.versions...)

			// Then
			if err != nil {
				if !errors.Is(err, semver.InvalidVersionErr) {
					t.Error("unexpected error")
				}
			} else {
				if test.expectedVersion != version {
					t.Errorf("expected version=%s, got=%s", test.expectedVersion, version)
				}
			}
		})
	}
}
