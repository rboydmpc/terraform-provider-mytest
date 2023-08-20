resource "MyTest_zone" "my_zone" {
    account_id = 6
            auto_recover_power_state = false
            code = "...my_code..."
            config = {
        zone_aws_config =     {
                use_host_credentials = "...my_use_host_credentials..."
                access_key = "...my_access_key..."
                appliance_url = "...my_appliance_url..."
                backup_mode = "...my_backup_mode..."
                certificate_provider = "...my_certificate_provider..."
                config_cmdb_discovery = true
                config_management_id = "...my_config_management_id..."
                costing_access_key = "...my_costing_access_key..."
                costing_bucket = "...my_costing_bucket..."
                costing_bucket_name = "...my_costing_bucket_name..."
                costing_folder = "...my_costing_folder..."
                costing_region = "...my_costing_region..."
                costing_report = "...my_costing_report..."
                costing_report_name = "...my_costing_report_name..."
                costing_secret_key = "...my_costing_secret_key..."
                costing_secret_key_hash = "...my_costing_secret_key_hash..."
                datacenter_name = "...my_datacenter_name..."
                dns_integration_id = "...my_dns_integration_id..."
                ebs_encryption = "...my_ebs_encryption..."
                endpoint = "...my_endpoint..."
                image_store_id = "...my_image_store_id..."
                is_vpc = "...my_is_vpc..."
                network_server = {
                    id = "d9d8d69a-674e-40f4-a7cc-8796ed151a05"
                }
                network_server_id = "...my_network_server_id..."
                replication_mode = "...my_replication_mode..."
                secret_key = "...my_secret_key..."
                secret_key_hash = "...my_secret_key_hash..."
                security_server = "...my_security_server..."
                service_registry_id = "...my_service_registry_id..."
                sts_assume_role = "...my_sts_assume_role..."
                vpc = "...my_vpc..."
            }
    }
            credential = {
        id = 9
        name = "Cedric Connelly"
        type = "...my_type..."
    }
            description = "...my_description..."
            enabled = false
            group_id = 3
            linked_account_id = 8
            location = "...my_location..."
            name = "Erik Lebsack"
            scale_priority = 1
            security_mode = "...my_security_mode..."
            visibility = "...my_visibility..."
            zone_type = {
        code = "...my_code..."
        id = 7
        name = "Seth Conroy"
    }
        }