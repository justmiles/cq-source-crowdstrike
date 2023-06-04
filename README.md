# CloudQuery CrowdStrike Falcon Source Plugin

[![release](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/release.yaml/badge.svg)](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/release.yaml) [![test](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/test.yaml/badge.svg)](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/test.yaml)

A crowdstrike source plugin for CloudQuery Falcon that loads data from crowdstrike to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Features

- Detections – Extracts details for Falcon detections, including behavior, severity, host, timestamps, and more.
- Hosts – Extracts host details including OS, version, sensor specific data, and more.
- Incidents - Extracts incidents and their details
- Vulnerabilities - Extracts vulnerabilities and their details

## Configuration

### Example

```yaml
# crowdstrike.yml
kind: source
spec:
  name: "crowdstrike"
  registry: "github"
  path: "justmiles/crowdstrike"
  version: "v0.0.0"
  destinations: ["sqlite"]
  tables: ["*"]
  spec:
    auth:
      strategy: "ondemand"
      creds:
        siteUrl: ${SP_SITE_URL}
        # align creds with the used strategy
```

## Tables

- [crowdstrike_falcon_detections](./docs/tables/crowdstrike_falcon_detections.md)
- [crowdstrike_falcon_hosts](./docs/tables/crowdstrike_falcon_hosts.md)
- [crowdstrike_falcon_incidents](./docs/tables/crowdstrike_falcon_incidents.md)
- [crowdstrike_falcon_vulnerabilities](./docs/tables/crowdstrike_falcon_vulnerabilities.md)

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```
