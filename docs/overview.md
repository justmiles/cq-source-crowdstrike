A crowdstrike source plugin for CloudQuery Falcon that loads data from crowdstrike to any database, data warehouse or data lake supported by CloudQuery, such as PostgreSQL, BigQuery, Athena, and many more.

## Features

- Detections – Extracts details for Falcon detections, including behavior, severity, host, timestamps, and more.
- Hosts – Extracts host details including OS, version, sensor specific data, and more.
- Incidents - Extracts incidents and their details
- Vulnerabilities - Extracts vulnerabilities and their details
- Disover Hosts - Extracts managed host details from Discover data
- Discover Applications - Extracts application information from Discover for all managed hosts

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
  registry: "cloudquery
  path: "justmiles/crowdstrike"
  version: "v2.0.0"
  # use this to enable incremental syncing - unimplemented
  # backend_options:
  #   table_name: "cq_state_crowdstrike"
  #   connection: "@@plugins.DESTINATION_NAME.connection"
  destinations: ["sqlite"]
  tables: ["*"]
  spec:
    # plugin spec section
```

### Plugin Spec

- `concurrency` (int, optional, default: `1000`):
  Best effort maximum number of Go routines to use. Lower this number to reduce memory usage.

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

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).

### Publish a new version to the Cloudquery Hub

After tagging a release, you can build and publish a new version to the [Cloudquery Hub](https://hub.cloudquery.io/) by running the following commands.
Replace `v1.0.0` with the new version number.

```bash
# "make dist" uses the README as main documentation and adds a generic release note. Output is created in dist/
VERSION=v1.0.0 make dist

# Login to cloudquery hub and publish the new version
cloudquery login
cloudquery plugin publish --finalize
```

After publishing the new version, it will show up in the [hub](https://hub.cloudquery.io/).

For more information please refer to the official [Publishing a Plugin to the Hub](https://cloudquery.io/docs/developers/publishing-a-plugin-to-the-hub) guide.
