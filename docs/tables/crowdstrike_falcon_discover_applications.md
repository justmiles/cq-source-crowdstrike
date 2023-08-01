# Table: crowdstrike_falcon_discover_applications

This table shows data for Crowdstrike Falcon Discover Applications.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|architectures|`list<item: utf8, nullable>`|
|category|`utf8`|
|cid|`utf8`|
|first_seen_timestamp|`utf8`|
|groups|`list<item: utf8, nullable>`|
|host|`json`|
|id (PK)|`utf8`|
|installation_paths|`list<item: utf8, nullable>`|
|installation_timestamp|`utf8`|
|is_normalized|`bool`|
|is_suspicious|`bool`|
|last_updated_timestamp|`utf8`|
|last_used_file_hash|`utf8`|
|last_used_file_name|`utf8`|
|last_used_timestamp|`utf8`|
|last_used_user_name|`utf8`|
|last_used_user_sid|`utf8`|
|name|`utf8`|
|name_vendor|`utf8`|
|name_vendor_version|`utf8`|
|vendor|`utf8`|
|version|`utf8`|
|versioning_scheme|`utf8`|