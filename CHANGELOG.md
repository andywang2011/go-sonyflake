# Changelog

## [Unreleased]

### Changed

- Migrated to `github.com/sony/sonyflake/v2` (`sonyflake.New`, `int` machine IDs, `int64` IDs internally). Public `NextID() (uint64, error)` unchanged; `New` errors now wrap the underlying v2 error.

## [0.2.0] - 2026-11-07

### Breaking

- See repository diff for API changes in this release (secure defaults, renamed packages where applicable, context-aware constructors).

### Security

- Hardened defaults: TLS, timeouts, body limits, PII redaction where logging applies.

### Fixed

- Bug fixes from the v0.2.0 hardening pass.

### Changed

- CI uses go-version-file; added race, staticcheck, govulncheck, golangci-lint (gosec).


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2026-11-07

### Added

- Config-driven Sonyflake wrapper with validated machine IDs.
