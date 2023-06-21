# CloudQuery CrowdStrike Falcon Source Plugin

[![release](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/release.yaml/badge.svg)](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/release.yaml) [![test](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/test.yaml/badge.svg)](https://github.com/justmiles/cq-source-crowdstrike/actions/workflows/test.yaml)

A crowdstrike source plugin for CloudQuery Falcon that loads data from crowdstrike to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Features

- Detections – Extracts details for Falcon detections, including behavior, severity, host, timestamps, and more.
- Hosts – Extracts host details including OS, version, sensor specific data, and more.
- Incidents - Extracts incidents and their details
- Vulnerabilities - Extracts vulnerabilities and their details

## Configuration

Export the following environment variables to configure CrowdStrike:

- FALCON_CLIENT_ID - Falcon ClientID
- FALCON_SECRET - Falcon Secret

To obtain an OAuth2 API Client navigate to the [CrowdStrike Falcon / API clients and keys](https://falcon.crowdstrike.com/api-clients-and-keys/clients) portal.

CloudQuery should only have _read_ access to CrowdStrike resources. The intent is to index all resources in CrowdStrike, so all scopes should be accessible, but read-only.

- Alerts
- Custom IOA rules
- Detections
- Hosts
- Falcon Discover
- Falcon Complete Dashboard
- Actors (Falcon Intelligence)
- Reports (Falcon Intelligence)
- Host groups
- Incidents
- Installation Tokens
- IOC Management
- IOCs (Indicators of Compromise)
- Message Center for Overwatch
- Message Center
- Machine Learning Exclusions
- On-demand scans (ODS)
- OverWatch Dashboard
- Prevention policies
- Quarantined Files
- Real time response (admin)
- Real time responseResponse policies
- Scheduled ReportsIOA Exclusions
- Sensor DownloadSensor update policies
- Sensor Visibility Exclusions
- Spotlight vulnerabilities
- Event streams
- User management
- Zero Trust Assessment

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
