# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `helm.sh/resource-policy: keep` annotation to default credential

## [2.3.0] - 2023-10-03

### Changed

- Update deployment to be PSS compliant and PSP toggle.

## [2.2.0] - 2023-07-03

### Added

- Add Service Monitor.

### Removed

- Stop pushing to `openstack-app-collection`.

## [2.1.1] - 2021-06-09

## [2.1.0] - 2021-06-09

### Changed

- Prepare helm values to configuration management.
- Update architect-orb to v3.0.0.

## [2.0.1] - 2020-10-26

### Fixed

- Change imports so that they work with v2.

## [2.0.0] - 2020-10-26

### Changed

- Update k8s dependencies to 1.18.9.

## [0.3.2] - 2020-07-13

### Changed

- Quotation around secret values.

## [0.3.1] 2020-06-19

### Changed

- No notable changes.



## [0.3.0] 2020-05-21

### Changed

- Deploy as a unique app in app collection



## [0.2.0] 2020-04-10

### Changed

- Migrate from dep to Go modules



## [0.1.0] 2020-04-10

### Added

- First release



[Unreleased]: https://github.com/giantswarm/giantswarm/compare/v2.3.0...HEAD
[2.3.0]: https://github.com/giantswarm/giantswarm/compare/v2.2.0...v2.3.0
[2.2.0]: https://github.com/giantswarm/giantswarm/compare/v2.1.1...v2.2.0
[2.1.1]: https://github.com/giantswarm/credentiald/compare/v2.1.0...v2.1.1
[2.1.0]: https://github.com/giantswarm/credentiald/compare/v2.17.0...v2.1.0
[2.17.0]: https://github.com/giantswarm/credentiald/compare/v2.0.1...v2.17.0
[2.0.1]: https://github.com/giantswarm/credentiald/compare/v2.0.0...v2.0.1
[2.0.0]: https://github.com/giantswarm/credentiald/compare/v0.3.2...v2.0.0
[0.3.2]: https://github.com/giantswarm/credentiald/compare/v0.3.1...v0.3.2
[0.3.0]: https://github.com/giantswarm/credentiald/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/credentiald/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/credentiald/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/credentiald/releases/tag/v0.1.0
