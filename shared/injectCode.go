package shared

const (
	// InjectCodeAuthorizationInfoByUserId 根据给定的用户ID获取授权信息【用户的角色ID列表】  {"userId":uid}
	InjectCodeAuthorizationInfoByUserId = "authorizationInfoByUserId"
	// InjectCodeAuthorizationInfoByRoleId 根据给定的角色ID获取授权信息【授权的资源ID列表】 {"roleId":rid}
	InjectCodeAuthorizationInfoByRoleId = "authorizationInfoByRoleId"
	// InjectCodeAddResourceInfo 添加授权资源信息
	InjectCodeAddResourceInfo = "addResourceInfo"
)

// 确保用户中心登录请求返回给用户的是标准jwt，使用统一的加密密钥，有名为 uid 的claims键，值为用户ID

const (
	JwtSecret         = "yyyooccckkiiiiiiii"
	JwtClaimUserId    = "uid"
	JwtClaimTenantId  = "tid"
	JwtClaimSessionId = "sid"
	RedisSessionIdKey = "sessionId:"

	SuperAdmin = "superAdmin"
)

type AuthorizationInfo struct {
	ResourceCodes []string
	RoleIds       []string
}

const (
	RedisKeyUserRoles        = "userRole:"
	RedisKeyRoleResourceCode = "roleResourceCode:"
)
