package security

type Authorization struct {
	Id         int64  `json:"id" xorm:"pk varchar(50)"`
	PolicyType int    `json:"policyType" xorm:"comment('权限类型 1-p, 2-p2')"`  // p  p2
	SubjectId  int64  `json:"subjectId" xorm:"comment('主体ID')"`             // 主体
	Resource   string `json:"resource" xorm:"comment('资源内容')"`              // 资源
	Action     string `json:"action" xorm:"comment('资源使用行为')"`              // 方法
	Effect     int    `json:"effect" xorm:"comment('策略行为 1-allow，2-deny')"` // 1 allow / 2 deny
	Priority   int    `json:"priority" xorm:"comment('优先级')"`               // 优先级
	TenantId   int64  `json:"tenantId" xorm:"comment('租户ID')"`
	ResourceId int64  `json:"resourceId" xorm:"varchar(50) comment('对应的资源ID')"` // 资源ID
}

func (_ Authorization) TableComment() string {
	return "权限配置表"
}

type Relationship struct {
	Id              int64 `json:"id,omitempty" xorm:"pk varchar(50)"`
	RelationType    int   `json:"relationType,omitempty" xorm:"comment('关系类型 1-g 2-g2')"`
	SubjectId       int64 `json:"subjectId,omitempty" xorm:"comment('主体ID')"`
	ParentSubjectId int64 `json:"parentSubjectId,omitempty" xorm:"comment('继承主体ID')"`
	TenantId        int64 `json:"tenantId" xorm:"comment('租户ID')"`
}

func (_ Authorization) Relationship() string {
	return "权限关系表"
}
