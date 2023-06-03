# Table: crowdstrike_falcon_detections

This table shows data for Crowdstrike Falcon Detections.

The primary key for this table is **detection_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|adversary_ids|`list<item: int64, nullable>`|
|assigned_to_name|`utf8`|
|assigned_to_uid|`utf8`|
|behaviors|`json`|
|behaviors_processed|`list<item: utf8, nullable>`|
|cid|`utf8`|
|created_timestamp|`json`|
|detection_id (PK)|`utf8`|
|device|`json`|
|email_sent|`bool`|
|first_behavior|`json`|
|hostinfo|`json`|
|last_behavior|`json`|
|max_confidence|`int64`|
|max_severity|`int64`|
|max_severity_displayname|`utf8`|
|overwatch_notes|`utf8`|
|quarantined_files|`json`|
|seconds_to_resolved|`int64`|
|seconds_to_triaged|`int64`|
|show_in_ui|`bool`|
|status|`utf8`|