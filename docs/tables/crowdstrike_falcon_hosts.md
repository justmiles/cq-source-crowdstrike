# Table: crowdstrike_falcon_hosts

This table shows data for Crowdstrike Falcon Hosts.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|agent_load_flags|`utf8`|
|agent_local_time|`utf8`|
|agent_version|`utf8`|
|bios_manufacturer|`utf8`|
|bios_version|`utf8`|
|build_number|`utf8`|
|cid|`utf8`|
|config_id_base|`utf8`|
|config_id_build|`utf8`|
|config_id_platform|`utf8`|
|cpu_signature|`utf8`|
|detection_suppression_status|`utf8`|
|device_id|`utf8`|
|device_policies|`json`|
|email|`utf8`|
|external_ip|`utf8`|
|first_login_timestamp|`utf8`|
|first_seen|`utf8`|
|group_hash|`utf8`|
|groups|`list<item: utf8, nullable>`|
|host_hidden_status|`utf8`|
|hostname|`utf8`|
|instance_id|`utf8`|
|internet_exposure|`utf8`|
|kernel_version|`utf8`|
|last_login_timestamp|`utf8`|
|last_seen|`utf8`|
|local_ip|`utf8`|
|mac_address|`utf8`|
|machine_domain|`utf8`|
|major_version|`utf8`|
|managed_apps|`json`|
|meta|`json`|
|minor_version|`utf8`|
|modified_timestamp|`utf8`|
|notes|`list<item: utf8, nullable>`|
|os_build|`utf8`|
|os_version|`utf8`|
|ou|`list<item: utf8, nullable>`|
|platform_id|`utf8`|
|platform_name|`utf8`|
|pod_annotations|`list<item: utf8, nullable>`|
|pod_host_ip4|`utf8`|
|pod_host_ip6|`utf8`|
|pod_hostname|`utf8`|
|pod_id|`utf8`|
|pod_ip4|`utf8`|
|pod_ip6|`utf8`|
|pod_labels|`list<item: utf8, nullable>`|
|pod_name|`utf8`|
|pod_namespace|`utf8`|
|pod_service_account_name|`utf8`|
|pointer_size|`utf8`|
|policies|`json`|
|product_type|`utf8`|
|product_type_desc|`utf8`|
|provision_status|`utf8`|
|reduced_functionality_mode|`utf8`|
|release_group|`utf8`|
|serial_number|`utf8`|
|service_pack_major|`utf8`|
|service_pack_minor|`utf8`|
|service_provider|`utf8`|
|service_provider_account_id|`utf8`|
|site_name|`utf8`|
|status|`utf8`|
|system_manufacturer|`utf8`|
|system_product_name|`utf8`|
|tags|`list<item: utf8, nullable>`|
|zone_group|`utf8`|