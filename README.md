# santosh-pipeline

The real pipeline running live on [santoshparajuli.sr.dev](https://santoshparajuli.sr.dev) — a small Go HTTP service with unit tests, a multi-stage Docker image, and a GitHub Pages deploy, all wired through one GitHub Actions workflow.

## Pipeline

![workflow status](https://github.com/spantosh/santosh-pipeline/actions/workflows/santosh-pipeline.yml/badge.svg)

| Stage        | What it does                          |
| ------------ | ------------------------------------- |
| Quality      | go vet + go test ./...                |
| Build        | multi-stage Docker build → GHCR       |
| Deploy Demo  | publishes a demo page to gh-pages     |

## Run it

```sh
go test ./...
docker build -t santosh-pipeline .
docker run --rm -p 8080:8080 santosh-pipeline
curl localhost:8080
```

The demo page is deployed to https://spantosh.github.io/santosh-pipeline/
| Pipeline |
| --- |
| Quality | `go vet`, unit tests and gofmt formatting gate |
| Hadolint | Dockerfile lint, fails on any warning |
| Vuln check | govulncheck audits Go module dependencies |
| Build + Trivy | multi-stage image pushed to GHCR, Trivy scan, fails on CRITICAL |
| Deploy Demo | demo page published to GitHub Pages from gh-pages |
