# Releasing

This SDK is versioned independently of the otari gateway, with its own semver.
Releases are automated with [release-please](https://github.com/googleapis/release-please).

## How a release happens

1. Merge changes to `main` using [Conventional Commits](https://www.conventionalcommits.org/)
   (`feat:`, `fix:`, etc.). This includes the gateway codegen's regeneration PRs
   and ordinary shell PRs.
2. release-please opens or updates a single **release PR** that writes
   `CHANGELOG.md` for the next version.
3. Review and merge the release PR. That tags the release (`vX.Y.Z`) and creates a
   GitHub Release.

There is no publish step. A Go module is released by its tag, so consumers pick up
the new version with `go get github.com/mozilla-ai/otari-sdk-go@vX.Y.Z`.

## Configuration

- **Registry:** none. The git tag is the release.
- **Auth:** none beyond the default `GITHUB_TOKEN`.
- **Version:** the git tag itself. A Go module has no version file to bump, so
  release-please only manages the changelog and the tag.

## Prerequisites (one time, repo settings)

- Enable **Settings to Actions: "Allow GitHub Actions to create and approve pull
  requests"** so release-please can open its release PR.

## If something goes wrong

The release is the tag plus the GitHub Release that release-please creates on the
merge of the release PR. Re-run the workflow from the Actions tab if it did not
complete; there is no separate publish to retry.

See the gateway's [SDK release coordination and compatibility](https://github.com/mozilla-ai/otari/blob/main/docs/sdk-compatibility.md)
for the cross-repo policy, the spec-version model, and the end-to-end flow.
