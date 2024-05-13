# `sample-go-bzl-ch`
This is a sample Golang repository with Bazel.

## Checklist
- [x] CLI:
  - [x] Cobra CLI
- [x] Http server:
  - [x] `gorilla/mux`
- [ ] Database:
  - [ ] Clickhouse with `database/sql` & `Clickhouse/clickhouse-go`
- [ ] Bazel:
  - [x] Build binary with `rules_go`
  - [x] Generate rules with `Gazelle`
  - [ ] Build OCI image with `rule_oci`
  - [x] Convert to Bzlmod
- [x] Development environment:
  - [x] Local docker compose
  - [x] Makefile
  - [x] Go workspace
  - [ ] Lint

## Make command

| Name                    | Desc                                                                               |
| ----------------------- | ---------------------------------------------------------------------------------- |
| test-all                | Run all tests in all modules                                                       |
| build-binary            | Build `sample-go-bzl-ch` binary                                                    |
| run-binary              | Run `sample-go-bzl-ch` binary. Arg for binary can be passed by set `app-arg` param |
| run-fullstack-in-docker | Run both binary & Clickhouse DB                                                    |
| generate-rule           | Generate Bazel rule using Gazelle                                                  |

## Notes
- To allow Bazel pull private dependency, add to .gitconfig an `insteadOf` block: 
```bash
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```
