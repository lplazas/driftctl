package resource

type ResourceType string

type ServiceType string

const (
	aws_s3  ServiceType = "s3"
	aws_ec2             = "ec2"
)

var supportedTypes = map[string]ResourceTypeMeta{
	"aws_ami":                     {services: []ServiceType{aws_ec2}},
	"aws_cloudfront_distribution": {},
	"aws_db_instance":             {},
	"aws_db_subnet_group":         {},
	"aws_default_network_acl": {children: []ResourceType{
		"aws_network_acl_rule",
	}},
	"aws_default_route_table": {children: []ResourceType{
		"aws_route",
	}},
	"aws_default_security_group": {children: []ResourceType{
		"aws_security_group_rule",
	}},
	"aws_default_subnet": {},
	"aws_default_vpc": {children: []ResourceType{
		// VPC are used by aws_internet_gateway to determine if internet gateway is the default one in middleware
		"aws_internet_gateway",
	}},
	"aws_dynamodb_table": {},
	"aws_ebs_snapshot":   {services: []ServiceType{aws_ec2}},
	"aws_ebs_volume":     {services: []ServiceType{aws_ec2}},
	"aws_ecr_repository": {},
	"aws_eip": {children: []ResourceType{
		"aws_eip_association",
	}, services: []ServiceType{aws_ec2}},
	"aws_eip_association":       {services: []ServiceType{aws_ec2}},
	"aws_iam_access_key":        {},
	"aws_iam_policy":            {},
	"aws_iam_policy_attachment": {},
	"aws_iam_role": {children: []ResourceType{
		"aws_iam_role_policy",
		"aws_iam_policy_attachment",
	}},
	"aws_iam_role_policy": {children: []ResourceType{
		"aws_iam_role_policy_attachment",
	}},
	"aws_iam_role_policy_attachment": {children: []ResourceType{
		"aws_iam_policy_attachment",
	}},
	"aws_iam_user": {children: []ResourceType{
		"aws_iam_user_policy",
	}},
	"aws_iam_user_policy": {children: []ResourceType{
		"aws_iam_user_policy_attachment",
	}},
	"aws_iam_user_policy_attachment": {children: []ResourceType{
		"aws_iam_policy_attachment",
	}},
	"aws_instance": {children: []ResourceType{
		"aws_ebs_volume",
	}, services: []ServiceType{aws_ec2}},
	"aws_internet_gateway": {children: []ResourceType{
		// This is used to determine internet gateway default rule
		"aws_route",
	}},
	"aws_key_pair":                    {services: []ServiceType{aws_ec2}},
	"aws_kms_alias":                   {},
	"aws_kms_key":                     {},
	"aws_lambda_event_source_mapping": {},
	"aws_lambda_function":             {},
	"aws_nat_gateway":                 {},
	"aws_network_acl": {children: []ResourceType{
		"aws_network_acl_rule",
	}},
	"aws_network_acl_rule":     {},
	"aws_route":                {},
	"aws_route53_health_check": {},
	"aws_route53_record":       {},
	"aws_route53_zone":         {},
	"aws_route_table": {children: []ResourceType{
		"aws_route",
	}},
	"aws_route_table_association": {},
	"aws_s3_bucket": {children: []ResourceType{
		"aws_s3_bucket_policy",
	}, services: []ServiceType{aws_s3}},
	"aws_s3_bucket_analytics_configuration": {services: []ServiceType{aws_s3}},
	"aws_s3_bucket_inventory":               {services: []ServiceType{aws_s3}},
	"aws_s3_bucket_metric":                  {services: []ServiceType{aws_s3}},
	"aws_s3_bucket_notification":            {services: []ServiceType{aws_s3}},
	"aws_s3_bucket_policy":                  {services: []ServiceType{aws_s3}},
	"aws_security_group": {children: []ResourceType{
		"aws_security_group_rule",
	}},
	"aws_security_group_rule": {},
	"aws_sns_topic": {children: []ResourceType{
		"aws_sns_topic_policy",
	}},
	"aws_sns_topic_policy":       {},
	"aws_sns_topic_subscription": {},
	"aws_sqs_queue": {children: []ResourceType{
		"aws_sqs_queue_policy",
	}},
	"aws_sqs_queue_policy":     {},
	"aws_subnet":               {},
	"aws_vpc":                  {},
	"aws_rds_cluster":          {},
	"aws_cloudformation_stack": {},
	"aws_api_gateway_rest_api": {children: []ResourceType{
		"aws_api_gateway_resource",
		"aws_api_gateway_rest_api_policy",
		"aws_api_gateway_gateway_response",
	}},
	"aws_api_gateway_account":    {},
	"aws_api_gateway_api_key":    {},
	"aws_api_gateway_authorizer": {},
	"aws_api_gateway_deployment": {children: []ResourceType{
		"aws_api_gateway_stage",
	}},
	"aws_api_gateway_stage": {},
	"aws_api_gateway_resource": {children: []ResourceType{
		"aws_api_gateway_method",
		"aws_api_gateway_integration",
	}},
	"aws_api_gateway_domain_name":       {},
	"aws_api_gateway_vpc_link":          {},
	"aws_api_gateway_request_validator": {},
	"aws_api_gateway_rest_api_policy":   {},
	"aws_api_gateway_base_path_mapping": {},
	"aws_api_gateway_model":             {},
	"aws_api_gateway_method": {children: []ResourceType{
		"aws_api_gateway_method_response",
	}},
	"aws_api_gateway_method_response":  {},
	"aws_api_gateway_gateway_response": {},
	"aws_api_gateway_method_settings":  {},
	"aws_api_gateway_integration": {children: []ResourceType{
		"aws_api_gateway_integration_response",
	}},
	"aws_api_gateway_integration_response": {},
	"aws_appautoscaling_target":            {},
	"aws_rds_cluster_instance": {children: []ResourceType{
		"aws_db_instance",
	}},
	"aws_appautoscaling_policy":           {},
	"aws_appautoscaling_scheduled_action": {},
	"aws_apigatewayv2_api": {children: []ResourceType{
		"aws_apigatewayv2_route",
		"aws_apigatewayv2_integration",
	}},
	"aws_apigatewayv2_model":                {},
	"aws_apigatewayv2_stage":                {},
	"aws_apigatewayv2_route_response":       {},
	"aws_apigatewayv2_deployment":           {},
	"aws_apigatewayv2_domain_name":          {},
	"aws_apigatewayv2_api_mapping":          {},
	"aws_apigatewayv2_route":                {},
	"aws_apigatewayv2_vpc_link":             {},
	"aws_apigatewayv2_authorizer":           {},
	"aws_apigatewayv2_integration":          {},
	"aws_apigatewayv2_integration_response": {},
	"aws_launch_template":                   {},
	"aws_launch_configuration":              {},

	"github_branch_protection": {},
	"github_membership":        {},
	"github_repository":        {},
	"github_team":              {},
	"github_team_membership":   {},

	"google_storage_bucket":   {},
	"google_compute_firewall": {},
	"google_compute_router":   {},
	"google_compute_instance": {},
	"google_compute_network":  {},
	"google_storage_bucket_iam_binding": {children: []ResourceType{
		"google_storage_bucket_iam_member",
	}},
	"google_storage_bucket_iam_member": {},
	"google_storage_bucket_iam_policy": {children: []ResourceType{
		"google_storage_bucket_iam_member",
	}},
	"google_dns_managed_zone":       {},
	"google_compute_instance_group": {},
	"google_bigquery_dataset":       {},
	"google_bigquery_table":         {},
	"google_project_iam_member":     {},
	"google_project_iam_binding": {children: []ResourceType{
		"google_project_iam_member",
	}},
	"google_project_iam_policy": {children: []ResourceType{
		"google_project_iam_member",
	}},
	"google_compute_address":         {},
	"google_compute_subnetwork":      {},
	"google_cloudfunctions_function": {},
	"google_compute_disk":            {},
	"google_bigtable_instance":       {},
	"google_bigtable_table":          {},
	"google_sql_database_instance":   {},
	"google_compute_image":           {},
	"google_compute_health_check":    {},
	"google_compute_global_address":  {},
	"google_compute_node_group":      {},
	"google_cloud_run_service":       {},

	"azurerm_storage_account":   {},
	"azurerm_storage_container": {},
	"azurerm_virtual_network": {children: []ResourceType{
		"azurerm_subnet",
	}},
	"azurerm_route_table": {children: []ResourceType{
		"azurerm_route",
	}},
	"azurerm_route":                    {},
	"azurerm_resource_group":           {},
	"azurerm_subnet":                   {},
	"azurerm_container_registry":       {},
	"azurerm_firewall":                 {},
	"azurerm_postgresql_server":        {},
	"azurerm_postgresql_database":      {},
	"azurerm_public_ip":                {},
	"azurerm_network_security_group":   {},
	"azurerm_lb":                       {},
	"azurerm_lb_rule":                  {},
	"azurerm_private_dns_zone":         {},
	"azurerm_private_dns_a_record":     {},
	"azurerm_private_dns_aaaa_record":  {},
	"azurerm_private_dns_cname_record": {},
	"azurerm_private_dns_ptr_record":   {},
	"azurerm_private_dns_srv_record":   {},
	"azurerm_private_dns_mx_record":    {},
	"azurerm_private_dns_txt_record":   {},
	"azurerm_image":                    {},
	"azurerm_ssh_public_key":           {},
}

func IsResourceTypeSupported(ty string) bool {
	_, exist := supportedTypes[ty]
	return exist
}

func (ty ResourceType) String() string {
	return string(ty)
}

func GetMeta(ty ResourceType) ResourceTypeMeta {
	return supportedTypes[ty.String()]
}

type ResourceTypeMeta struct {
	children []ResourceType
	services []ServiceType
}

func (ty ResourceTypeMeta) GetChildrenTypes() []ResourceType {
	return ty.children
}

func FindResourcesForServiceType(ty string) []string {
	var res []string
	for r, meta := range supportedTypes {
		for _, s := range meta.services {
			if string(s) == ty {
				res = append(res, r)
			}
		}
	}
	return res
}

func GetServiceTypes() []string {
	var res []string
	for _, meta := range supportedTypes {
		for _, s := range meta.services {
			if !findInSlice(res, string(s)) {
				res = append(res, string(s))
			}
		}
	}
	return res
}

func IsServiceTypeSupported(ty string) bool {
	return findInSlice(GetServiceTypes(), ty)
}

func findInSlice(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
