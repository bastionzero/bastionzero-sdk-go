# Contributing

This document details how to make contributions to this repo.

## Changelog

The `CHANGELOG` file follows a format similar to HashiCorp's changelog
specifications they recommend to use for Terraform providers found
[here](https://www.terraform.io/plugin/sdkv2/best-practices/versioning#changelog-specification).

Information in the changelog should be broken down into the following categories:
- **BREAKING CHANGES:** This section documents in brief any incompatible changes and how to handle them. This should only be present in major version upgrades.
- **FEATURES:** These are _major_ new improvements that deserve a special highlight, such as a new API endpoint or service.
- **ENHANCEMENTS:** Smaller features added to the project such as a new field for an already existing API service or endpoint.
- **BUG FIXES:** Any bugs that were fixed.
- **NOTES:** Catch-all category for changes that do not fit the categories above, such as potentially unexpected upgrade behavior, upcoming deprecations, or 3rd party library upgrades.

### Changie Automation Tool

`bastionzero-sdk-go` uses the [Changie](https://changie.dev/) automation tool
for changelog automation.

To add a new entry to the `CHANGELOG`, install Changie using the following [instructions](https://changie.dev/guide/installation/)
and run:

```bash
changie new
```

then choose a `kind` of change corresponding to the categories specified above.
Make sure to fill out the `body` following the entry format described below
(note: you can always update the auto-generated file later if you prefer to edit
the body in your text editor of choice).

Changie will then prompt for a Github issue or pull request number (note: If
your change spans across multiple issues or PRs, you can include all of them as
a comma separated list of numbers). _Repeat_ this process for any additional
changes. The `.yaml` files created in the `.changes/unreleased` folder should be
pushed to the repository along with any code changes.

#### Entry (`body`) format

Please try to use the following format when specifying the `body` field:

```markdown
<package>/<optional-extra-identifier>: <required-short-description>
```

For example,

```markdown
policies/jit: Add support for JIT policy API
```

If the change is not specific to any one package or does not map 1-to-1 with an
API service package, use the outermost package folder as the prefix for the
entry.

For example,

```markdown
bastionzero/service: Use new `types.Timestamp` instead of `time.Time` when an API request/response has a timestamp part of its spec.
This change is not a breaking change as `types.Timestamp` embeds a `time.Time` and is therefore backward compatible with prior usage
```

Do not include a trailing period as the generated file includes one for you.

#### Pull Request Types to `CHANGELOG`

The `CHANGELOG` is intended to show consumer-impacting changes to the codebase
for a particular version. If every change or commit to the code resulted in an
entry, the `CHANGELOG` would become less useful for consumers. The lists below
are general _guidelines_ to decide whether a change should have an entry.

##### Changes that should not have a `CHANGELOG` entry

* Documentation updates
* New tests or changes to existing tests
* Code refactoring that has no visible impact to the consumer (does not affect how they use exported package members of this library)

##### Changes that should have a `CHANGELOG` entry

* Major features
* Bug fixes
* Enhancements
* Deprecations
* Breaking changes and removals
* Code refactoring that has visible impact to the consumer
* Dependency updates

## Releasing

Releasing a new version of `bastionzero-sdk-go` is a semi-automated process.

Use the "Generate release pull request" workflow to auto-generate a release PR
that collates all unreleased changes in `.changes/unreleased` and updates the
`CHANGELOG.md` accordingly.

Please double check the auto-generated PR and complete the remaining TODOs as
outlined in the PR description. Namely, please update
`bastionzero/bastionzero.go#libraryVersion` and commit this change before
merging.

When the release PR is merged to `master`, a draft release is automatically
generated. Please double check the release and publish the release when ready.
