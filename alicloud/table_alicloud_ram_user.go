package alicloud

import (
	"context"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"

	"github.com/turbot/steampipe-plugin-sdk/v2/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v2/plugin/transform"
)

type userInfo = struct {
	UserName      string
	UserId        string
	DisplayName   string
	Email         string
	MobilePhone   string
	Comments      string
	CreateDate    string
	UpdateDate    string
	LastLoginDate string
}

//// TABLE DEFINITION

func tableAlicloudRAMUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "alicloud_ram_user",
		Description: "Resource Access Management users who can login via the console or access keys.",
		List: &plugin.ListConfig{
			Hydrate: listRAMUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns:        plugin.SingleColumn("name"),
			ShouldIgnoreError: isNotFoundError([]string{"EntityNotExist.User", "MissingParameter"}),
			Hydrate:           getRAMUser,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "name",
				Description: "The username of the RAM user.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UserName"),
			},
			{
				Name:        "user_id",
				Description: "The unique ID of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "display_name",
				Description: "The display name of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "email",
				Description: "The email address of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "last_login_date",
				Description: "The time when the RAM user last logged on to the console by using the password.",
				Type:        proto.ColumnType_TIMESTAMP,
				Hydrate:     getRAMUser,
			},
			{
				Name:        "mobile_phone",
				Description: "The mobile phone number of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "comments",
				Description: "The description of the RAM user.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "create_date",
				Description: "The time when the RAM user was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "update_date",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time when the RAM user was modified.",
			},
			{
				Name:        "mfa_enabled",
				Description: "The MFA status of the user",
				Type:        proto.ColumnType_BOOL,
				Hydrate:     getRAMUserMfaDevices,
				Transform:   transform.From(userMfaStatus),
			},
			{
				Name:        "mfa_device_serial_number",
				Description: "The serial number of the MFA device.",
				Type:        proto.ColumnType_STRING,
				Hydrate:     getRAMUserMfaDevices,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "attached_policy",
				Description: "A list of policies attached to a RAM user.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMUserPolicies,
				Transform:   transform.FromField("Policies.Policy"),
			},
			{
				Name:        "cs_user_permissions",
				Description: "User permissions for Container Service Kubernetes clusters.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getCsUserPermissions,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "groups",
				Description: "A list of groups attached to the user.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMUserGroups,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "virtual_mfa_devices",
				Description: "The list of MFA devices.",
				Type:        proto.ColumnType_JSON,
				Hydrate:     getRAMUserMfaDevices,
				Transform:   transform.FromValue(),
			},

			// Steampipe standard columns
			{
				Name:        "akas",
				Description: ColumnDescriptionAkas,
				Type:        proto.ColumnType_JSON,
				Hydrate:     getUserAkas,
				Transform:   transform.FromValue(),
			},
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("UserName"),
			},

			// alicloud standard columns
			{
				Name:        "region",
				Description: ColumnDescriptionRegion,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromConstant("global"),
			},
			{
				Name:        "account_id",
				Description: ColumnDescriptionAccount,
				Type:        proto.ColumnType_STRING,
				Hydrate:     getCommonColumns,
				Transform:   transform.FromField("AccountID"),
			},
		},
	}
}

//// LIST FUNCTION

func listRAMUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_user.listRAMUser", "connection_error", err)
		return nil, err
	}
	request := ram.CreateListUsersRequest()
	request.Scheme = "https"
	for {
		response, err := client.ListUsers(request)
		if err != nil {
			plugin.Logger(ctx).Error("alicloud_ram_user.listRAMUser", "query_error", err, "request", request)
			return nil, err
		}
		for _, i := range response.Users.User {
			plugin.Logger(ctx).Warn("listRAMUser", "item", i)
			d.StreamListItem(ctx, userInfo{i.UserName, i.UserId, i.DisplayName, i.Email, i.MobilePhone, i.Comments, i.CreateDate, i.UpdateDate, ""})
		}
		if !response.IsTruncated {
			break
		}
		request.Marker = response.Marker
	}
	return nil, nil
}

//// HYDRATE FUNCTIONS

func getRAMUser(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMUser")

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_user.getRAMUser", "connection_error", err)
		return nil, err
	}

	var name string
	if h.Item != nil {
		i := h.Item.(userInfo)
		name = i.UserName
	} else {
		name = d.KeyColumnQuals["name"].GetStringValue()
	}

	request := ram.CreateGetUserRequest()
	request.Scheme = "https"
	request.UserName = name

	response, err := client.GetUser(request)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_user.getRAMUser", "query_error", err, "request", request)
		return nil, err
	}

	data := response.User
	return userInfo{data.UserName, data.UserId, data.DisplayName, data.Email, data.MobilePhone, data.Comments, data.CreateDate, data.UpdateDate, data.LastLoginDate}, nil
}

func getRAMUserGroups(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMUserGroups")
	data := h.Item.(userInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserGroups", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListGroupsForUserRequest()
	request.Scheme = "https"
	request.UserName = data.UserName

	response, err := client.ListGroupsForUser(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserGroups", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response.Groups.Group, nil
}

func getRAMUserPolicies(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMUserPolicies")
	data := h.Item.(userInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserPolicies", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListPoliciesForUserRequest()
	request.Scheme = "https"
	request.UserName = data.UserName

	response, err := client.ListPoliciesForUser(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserPolicies", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	return response, nil
}

func getRAMUserMfaDevices(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getRAMUserMfaDevices")
	data := h.Item.(userInfo)

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserMfaDevices", "connection_error", err)
		return nil, err
	}

	request := ram.CreateListVirtualMFADevicesRequest()
	request.Scheme = "https"

	response, err := client.ListVirtualMFADevices(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("alicloud_ram_group.getRAMUserMfaDevices", "query_error", serverErr, "request", request)
		return nil, serverErr
	}

	var items []ram.VirtualMFADeviceInListVirtualMFADevices
	for _, i := range response.VirtualMFADevices.VirtualMFADevice {
		if i.User.UserName == data.UserName {
			items = append(items, i)
		}
	}

	return items, nil
}

func getUserAkas(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getUserAkas")
	data := h.Item.(userInfo)

	// Get project details
	getCommonColumnsCached := plugin.HydrateFunc(getCommonColumns).WithCache()
	commonData, err := getCommonColumnsCached(ctx, d, h)
	if err != nil {
		return nil, err
	}
	commonColumnData := commonData.(*alicloudCommonColumnData)
	accountID := commonColumnData.AccountID

	return []string{"acs:ram::" + accountID + ":user/" + data.UserName}, nil
}

func getCsUserPermissions(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	plugin.Logger(ctx).Trace("getCsUserPermissions")

	// Create service connection
	client, err := RAMService(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getCsUserPermissions", "connection_error", err)
		return nil, err
	}

	data := h.Item.(userInfo)

	request := requests.NewCommonRequest()
	request.Method = "GET"
	request.Scheme = "https"
	request.Domain = "cs.aliyuncs.com"
	request.Version = "2015-12-15"
	request.PathPattern = "/permissions/users/" + data.UserId
	request.Headers["Content-Type"] = "application/json"

	response, err := client.ProcessCommonRequest(request)
	if serverErr, ok := err.(*errors.ServerError); ok {
		plugin.Logger(ctx).Error("getCsUserPermissions", "query_error", serverErr, "request", request)
		return nil, err
	}

	return response.GetHttpContentString(), nil
}

//// TRANSFORM FUNCTION

func userMfaStatus(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.([]ram.VirtualMFADeviceInListVirtualMFADevices)

	if len(data) > 0 {
		return true, nil
	}

	return false, nil
}
