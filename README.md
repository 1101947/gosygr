# gosygr
is Go syntax graph library

# CHANGELOG file
This project uses changelog file(see CHANGELOG.md).
Changelog file contains log of changes for each project version(see version section).
It allows users and developers to see what changes have been made in new version, what features have been added and if any bugs or security vulnerabilities was introduced or fixed.
Format was inspired by [Keep a Changelog](https://keepachangelog.com/en/1.1.0/).
Unlike Keep a Changelog, current version is tagged "current" and don't have a proper version string like all previous versions, we don't use [semver](https://semver.org/spec/v2.0.0.html)(see versioning section), we have additional Bug-introduced and Bug-found types of changes.

# Versioning
This project adheres to so called(by us) commitver: commit versioning scheme.
In this versioning shceme, version is a string, a combination of a timestamp of source code release and its hash.
For example, if version was released at 2026-07-22 16:20:43 and its unique hash string is 100cf2e0284c25e0c9a7c6433ee516409514ad63, the version string will be:
v2026-07-22_16-20-43Z__100cf2e0284c25e0c9a7c6433ee516409514ad63.
Every commit in master branch should contain working code, but to be sure always address to CHANGELOG.md.
To see semantics, added features, introduced and fixed bugs of any version address to CHANGELOG.md.
