resource "MyTest_zone" "my_zone" {
    account_id = 1
            auto_recover_power_state = false
            code = "mycloud"
            config = {
        api_url = "...my_api_url..."
        datacenter = "...my_datacenter..."
        password = "...my_password..."
        username = "Micheal_Sporer"
    }
            credential = {
        type = "...my_type..."
    }
            description = "...my_description..."
            enabled = false
            group_id = 3
            linked_account_id = 4
            location = "US East"
            name = "My Cloud"
            scale_priority = 9
            security_mode = "...my_security_mode..."
            visibility = "private"
            zone_type = {
        code = "vmware"
    }
        }