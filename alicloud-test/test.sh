#!/bin/bash
 RED="\e[31m"
 GREEN="\e[32m"
 BLACK="\e[30m"
 BOLDGREEN="\e[1;${GREEN}"
 ENDCOLOR="\e[0m"


# Define your function here
run_test () {
   echo -e "${GREEN}Running $1 ${ENDCOLOR}"
 if ! ./tint.js $1 >> output.txt
   then
    echo -e "${RED}Failed -> $1 ${ENDCOLOR}"
    echo $1 >> failed_tests.txt
  else
    echo -e "${BOLDGREEN}Passed -> $1 ${ENDCOLOR}"
    echo $1 >> passed_tests.txt
   fi
 }

 # output.txt - store output of each test
 # failed_tests.txt - names of failed test
 # passed_tests.txt names of passed test

 # removes files from previous test
# rm -rf output.txt failed_tests.txt passed_tests.txt
 date >> output.txt
 date >> failed_tests.txt
 date >> passed_tests.txt

run_test alicloud_account
run_test alicloud_action_trail
run_test alicloud_cas_certificate
run_test alicloud_cms_monitor_host
run_test alicloud_cs_kubernetes_cluster
run_test alicloud_cs_kubernetes_cluster_node
run_test alicloud_ecs_auto_provisioning_group
run_test alicloud_ecs_key_pair
run_test alicloud_ecs_launch_template
run_test alicloud_ecs_network_interface
run_test alicloud_ecs_region
run_test alicloud_ecs_zone
run_test alicloud_kms_key
run_test alicloud_kms_secret
run_test alicloud_oss_bucket
run_test alicloud_ram_policy
run_test alicloud_ram_user
run_test alicloud_rds_instance
run_test alicloud_vpc_dhcp_options_set
run_test alicloud_vpc_nat_gateway
run_test alicloud_vpc_network_acl
run_test alicloud_vpc_route_entry
run_test alicloud_vpc_route_table
run_test alicloud_vpc_vpn_customer_gateway
run_test alicloud_vpc_vpn_gateway

date >> output.txt
date >> failed_tests.txt
date >> passed_tests.txt