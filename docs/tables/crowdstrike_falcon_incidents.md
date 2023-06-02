# Table: crowdstrike_falcon_incidents

This table shows data for Crowdstrike Falcon Incidents.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|assigned_to|`utf8`|
|assigned_to_name|`utf8`|
|cid|`utf8`|
|created|`json`|
|description|`utf8`|
|end|`json`|
|events_histogram|`json`|
|fine_score|`int64`|
|host_ids|`list<item: utf8, nullable>`|
|hosts|`json`|
|incident_id|`utf8`|
|incident_type|`int64`|
|lm_host_ids|`list<item: utf8, nullable>`|
|lm_hosts_capped|`bool`|
|modified_timestamp|`json`|
|name|`utf8`|
|objectives|`list<item: utf8, nullable>`|
|start|`json`|
|state|`utf8`|
|status|`int64`|
|tactics|`list<item: utf8, nullable>`|
|tags|`list<item: utf8, nullable>`|
|techniques|`list<item: utf8, nullable>`|
|users|`list<item: utf8, nullable>`|
|visibility|`int64`|