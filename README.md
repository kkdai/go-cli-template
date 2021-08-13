# Go CLI Template
> With Github Actions CI flow to build your package as a release

[![Go CI](https://github.com/kkdai/go-cli-template/workflows/Go/badge.svg)](https://github.com/kkdai/go-cli-template/actions?query=workflow:"Go")
[![kkdai - go-cli-template](https://img.shields.io/static/v1?label=kkdai&message=go-cli-template&color=blue&logo=github)](https://github.com/kkdai/go-cli-template)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)

[![Made with Go](https://img.shields.io/badge/Go-1-blue?logo=go&logoColor=white)](https://golang.org)


## Installation

```sh
go get -u -x github.com/kkdai/go-cli-template
```

## How To Use

```sh
go-cli-template --t="any test you want"  
```

Just clone it for your CLI program template.


## Set up CI

### Manual

_Warning: this token will have access to write to _all_ of your GitHub repos. Keep it secret!_

Add a GitHub token to your repo's secrets.

[Settings] -> [Secrets] -> [New Secret]

- `Name`: `GO_RELEASER_GITHUB_TOKEN`
- `Token`: Get GitHub token from https://github.com/settings/tokens with appropriate write permissions.

That will be used at the end of [release_build](/.github/workflows/release_build.yml).

```yaml
env:
  GITHUB_TOKEN: ${{ secrets.GO_RELEASER_GITHUB_TOKEN }}
```

### Automated

For a safer and low-effort alternative, you can use a token that GitHub Actions generates for you on each run. You don't even have to generate or view this value _and_ it is scoped to _only_ work for the current repo, so it is much safer.

Change the end of [release_build](/.github/workflows/release_build.yml) to use this instead:

```yaml
env:
  GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```


## License

Released under [MIT](/LICENSE) by [@kkdai](https://github.com/kkdai).
