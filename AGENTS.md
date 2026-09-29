# Terraform provider

See `../AGENTS.md` for the workspace overview.

- `api/v1` is the **generated** API client: `make generate-client` (Docker, reads `../api`).
- Resources in `internal/provider/*_resource.go`: a `…Model` struct with `tfsdk` tags,
  schema, `fromAPI` / `toAPIRequest`. Use the `omittable…Value` helpers from
  `helpers.go` for optional attributes so unset values stay null.
- Attributes the server derives when omitted (e.g. allocation key `type`, `identifier`)
  are `Optional` + `Computed` with `UseStateForUnknown`. Use `stringSlice(api.Allowed…EnumValues)`
  with `stringvalidator.OneOf` for enums.
- After schema changes: `make generate-docs` (tfplugindocs) and commit `docs/`.
- Acceptance tests need a running backend and `UDOMA_ACCOUNT_REF` etc.; without them
  `go build ./... && go vet ./...` is the available check.

New findings about this repo go into this file; see "Keeping these notes up to date" in `../AGENTS.md`.
