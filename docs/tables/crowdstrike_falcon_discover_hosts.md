# Table: crowdstrike_falcon_discover_hosts

This table shows data for Crowdstrike Falcon Discover Hosts.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_enabled|`utf8`|
|ad_user_account_control|`int64`|
|agent_version|`utf8`|
|aid|`utf8`|
|assigned_to|`utf8`|
|available_disk_space|`int64`|
|available_disk_space_pct|`int64`|
|average_memory_usage|`int64`|
|average_memory_usage_pct|`int64`|
|average_processor_usage|`int64`|
|bios_hashes_data|`json`|
|bios_id|`utf8`|
|bios_manufacturer|`utf8`|
|bios_version|`utf8`|
|cid|`utf8`|
|city|`utf8`|
|classification|`utf8`|
|confidence|`int64`|
|country|`utf8`|
|cpu_manufacturer|`utf8`|
|cpu_processor_name|`utf8`|
|creation_timestamp|`utf8`|
|current_local_ip|`utf8`|
|data_providers|`list<item: utf8, nullable>`|
|data_providers_count|`int64`|
|department|`utf8`|
|descriptions|`list<item: utf8, nullable>`|
|discoverer_aids|`list<item: utf8, nullable>`|
|discoverer_count|`int64`|
|discoverer_platform_names|`list<item: utf8, nullable>`|
|discoverer_product_type_descs|`list<item: utf8, nullable>`|
|discoverer_tags|`list<item: utf8, nullable>`|
|disk_sizes|`json`|
|email|`utf8`|
|encrypted_drives|`list<item: utf8, nullable>`|
|encrypted_drives_count|`int64`|
|encryption_status|`utf8`|
|entity_type|`utf8`|
|external_ip|`utf8`|
|field_metadata|`json`|
|first_discoverer_aid|`utf8`|
|first_seen_timestamp|`utf8`|
|fqdn|`utf8`|
|groups|`list<item: utf8, nullable>`|
|hostname|`utf8`|
|id (PK)|`utf8`|
|internet_exposure|`utf8`|
|kernel_version|`utf8`|
|last_discoverer_aid|`utf8`|
|last_seen_timestamp|`utf8`|
|local_ip_addresses|`list<item: utf8, nullable>`|
|local_ips_count|`int64`|
|location|`utf8`|
|logical_core_count|`int64`|
|mac_addresses|`list<item: utf8, nullable>`|
|machine_domain|`utf8`|
|managed_by|`utf8`|
|max_memory_usage|`int64`|
|max_memory_usage_pct|`int64`|
|max_processor_usage|`int64`|
|mount_storage_info|`json`|
|network_interfaces|`json`|
|number_of_disk_drives|`int64`|
|object_guid|`utf8`|
|object_sid|`utf8`|
|os_is_eol|`utf8`|
|os_security|`json`|
|os_service_pack|`utf8`|
|os_version|`utf8`|
|ou|`utf8`|
|owned_by|`utf8`|
|physical_core_count|`int64`|
|platform_name|`utf8`|
|processor_package_count|`int64`|
|product_type|`utf8`|
|product_type_desc|`utf8`|
|reduced_functionality_mode|`utf8`|
|servicenow_id|`utf8`|
|site_name|`utf8`|
|state|`utf8`|
|system_manufacturer|`utf8`|
|system_product_name|`utf8`|
|system_serial_number|`utf8`|
|tags|`list<item: utf8, nullable>`|
|total_bios_files|`int64`|
|total_disk_space|`int64`|
|total_memory|`int64`|
|unencrypted_drives|`list<item: utf8, nullable>`|
|unencrypted_drives_count|`int64`|
|used_disk_space|`int64`|
|used_disk_space_pct|`int64`|
|used_for|`utf8`|