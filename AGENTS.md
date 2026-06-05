# ogc-osgeo

Go structs for OGC web service requests, responses, capabilities, and exceptions. Not a server — pure data modelling with XML + YAML marshalling and KVP query-parameter parsing.

## Commands

```bash
go build ./...
go test ./...
go test -bench=. ./pkg/wms130/...   # benchmarks in wms130, wfs200
```

## Package map

| Package | Standard | Completeness |
|---|---|---|
| `pkg/wms130` | WMS 1.3.0 | Full (Time/Elevation dimensions as TODOs) |
| `pkg/wfs200` | WFS 2.0.0 | Read-only ops (GetFeature, GetCapabilities, DescribeFeatureType). Transaction, LockFeature, StoredQueries absent. |
| `pkg/wmts100` | WMTS 1.0.0 | Full request/response with KVP + RESTful path parsing |
| `pkg/wcs201` | WCS 2.0.1 | GetCoverage, DescribeCoverage, GetCapabilities implemented. No tests. |
| `pkg/tms100` | TMS 1.0.0 (OSGeo) | Complete |
| `pkg/wsc110` | OWS Common 1.1.0 | Foundation — interfaces used by wfs200, wmts100 |
| `pkg/wsc200` | OWS Common 2.0 | Exception-only subset of wsc110. Used by wcs201. |
| `pkg/utils` | Shared helpers | `KeysToUpper/Lower`, `XMLAttribute`, `StripDuplicateAttr` |

## Architecture

- **wsc110** defines `Capabilities`, `Exception`, `OperationRequest` interfaces; wfs200 and wmts100 implement them. wms130 has its own parallel `OperationRequest`. wcs201 uses `wsc200` exceptions (not wsc110) and does not implement the wsc110 interfaces.
- **Two exception systems**: OWS Common (`wsc110`/`wsc200` — `ows:ExceptionReport`) and WMS (`wms130` — `ServiceExceptionReport`). Exception constructors return types with `Code()`, `Locator()`, `Error()`, collected in `Exceptions` slices → `ToReport().ToBytes()`.
- **`_pv.go` suffix** = KVP ("Parameter Value") encoding/decoding helper structs.
- **Dual-struct pattern** for XML namespaces: `*Unmarshal` structs parse (no prefixes), regular structs marshal (with prefixes). See `BoundingBox` vs `BoundingBoxUnmarshal`.
- All structs carry both `xml` and `yaml` tags.
- Each request type has `Validate(Capabilities)` returning exception slices.
- INSPIRE `ExtendedCapabilities` in wms130, wfs200, wmts100, wcs201.

## Testing

- 37 test files across all packages, all table-driven unit tests with inline data (no fixtures). wcs201 went from 0 to 3 test files (22 tests), wmts100 from 1 to 4 test files (30 tests).
- No integration, e2e, or mock-based tests.

## Quirks

- Go 1.16 (no `any`, `errors.Is`, generics). Single dependency: `gopkg.in/yaml.v3`.
- wcs201 has 3 test files (22 tests). Its KVP parsing method is named `QueryParameters` (not `ParseQueryParameters` like other packages).
- No CI, no linter config, no Makefile, no generated code. `standards/` contains spec PDFs for reference only.
