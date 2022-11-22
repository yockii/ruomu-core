package security

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	logger "github.com/sirupsen/logrus"
	"xorm.io/xorm"

	"github.com/yockii/ruomu-core/util"
)

type adapter struct {
	engine *xorm.Engine
}

func finalizer(a *adapter) {
	if a.engine == nil {
		return
	}
	err := a.engine.Close()
	if err != nil {
		logger.Warnln("Failed to close casbin adapter engine [xorm-security], err: ", err)
	}
}

func NewAdapter(engine *xorm.Engine) (*adapter, error) {
	if err := engine.Sync2(Authorization{}, Relationship{}); err != nil {
		return nil, err
	}
	return &adapter{engine: engine}, nil
}

func (a *adapter) LoadPolicy(m model.Model) error {
	var policies []*Authorization
	if err := a.engine.Find(&policies); err != nil {
		return err
	}
	for _, policy := range policies {
		policyType := "p"
		if policy.PolicyType != 1 {
			policyType = fmt.Sprintf("p%d", policy.PolicyType)
		}
		effect := "allow"
		if policy.Effect == 2 {
			effect = "deny"
		}
		tokens := []string{
			policyType,
			strconv.FormatInt(policy.SubjectId, 10),
			policy.Resource,
			policy.Action,
			effect,
			strconv.Itoa(policy.Priority),
			strconv.FormatInt(policy.TenantId, 10),
			strconv.FormatInt(policy.ResourceId, 10),
		}

		persist.LoadPolicyArray(tokens, m)
		//mpp := m["p"][policyType]
		//mpp.Policy = append(mpp.Policy, tokens)
		//mpp.PolicyMap[strings.Join(tokens, model.DefaultSep)] = len(mpp.Policy) - 1
	}

	var relations []*Relationship
	if err := a.engine.Find(&relations); err != nil {
		return err
	}
	for _, relation := range relations {
		relationType := "g"
		if relation.RelationType != 1 {
			relationType = fmt.Sprintf("g%d", relation.RelationType)
		}
		tokens := []string{
			relationType,
			strconv.FormatInt(relation.SubjectId, 10),
			strconv.FormatInt(relation.ParentSubjectId, 10),
			strconv.FormatInt(relation.TenantId, 10),
		}
		persist.LoadPolicyArray(tokens, m)
		//mgg := m["g"][relationType]
		//mgg.Policy = append(mgg.Policy, tokens)
		//mgg.PolicyMap[strings.Join(tokens, model.DefaultSep)] = len(mgg.Policy) - 1
	}
	return nil
}

// SavePolicy saves all policy rules to the storage.
func (a *adapter) SavePolicy(m model.Model) error {
	sess := a.engine.NewSession()
	defer sess.Close()
	sess.Delete(&Relationship{})
	sess.Delete(&Authorization{})

	var policies []*Authorization
	var roles []*Relationship
	for policyType, ast := range m["p"] {
		for _, rule := range ast.Policy {
			if len(rule) == 7 {
				policy, err := a.parsePolicy(policyType, rule)
				if err != nil {
					logger.Error("策略规则必须是7个元素的数组", err)
					continue
				}
				policies = append(policies, policy)
			}
		}
	}
	for relationType, ast := range m["g"] {
		for _, rule := range ast.Policy {
			role, err := a.parseRelation(relationType, rule)
			if err != nil {
				logger.Error("关系规则必须是3个元素的数组", err)
				continue
			}
			roles = append(roles, role)
		}
	}

	_, err := sess.Insert(policies)
	if err != nil {
		return err
	}
	_, err = sess.Insert(roles)
	if err != nil {
		return err
	}

	return sess.Commit()
}

func (a *adapter) parseRelation(relationType string, rule []string) (*Relationship, error) {
	if len(rule) != 3 {
		return nil, errors.New("非法的父子关系规则数量，数量必须是3，实际" + strconv.Itoa(len(rule)))
	}
	gt := 1
	if s := relationType[1:]; s != "" {
		gt, _ = strconv.Atoi(relationType[1:])
	}
	subjectId, _ := strconv.ParseInt(rule[0], 10, 64)
	parentSubjectId, _ := strconv.ParseInt(rule[1], 10, 64)
	tenantId, _ := strconv.ParseInt(rule[2], 10, 64)
	role := &Relationship{
		Id:              util.SnowflakeId(),
		RelationType:    gt,
		SubjectId:       subjectId,
		ParentSubjectId: parentSubjectId,
		TenantId:        tenantId,
	}
	return role, nil
}

func (a *adapter) parsePolicy(policyType string, rule []string) (*Authorization, error) {
	if len(rule) != 7 {
		return nil, errors.New("invalid policy rule ")
	}
	pt := 1
	if s := policyType[1:]; s != "" {
		pt, _ = strconv.Atoi(policyType[1:])
	}
	effect := 1
	if rule[3] == "deny" {
		effect = 2
	}
	priority, _ := strconv.Atoi(rule[4])

	subjectId, _ := strconv.ParseInt(rule[0], 10, 64)
	tenantId, _ := strconv.ParseInt(rule[5], 10, 64)
	resourceId, _ := strconv.ParseInt(rule[6], 10, 64)
	policy := &Authorization{
		Id:         util.SnowflakeId(),
		PolicyType: pt,
		SubjectId:  subjectId,
		Resource:   rule[1],
		Action:     rule[2],
		Effect:     effect,
		Priority:   priority,
		TenantId:   tenantId,
		ResourceId: resourceId,
	}
	return policy, nil
}

// AddPolicy adds a policy rule to the storage.
func (a *adapter) AddPolicy(sec string, ptype string, rule []string) error {
	if sec == "p" {
		policy, err := a.parsePolicy(ptype, rule)
		if err != nil {
			return err
		}
		_, err = a.engine.InsertOne(policy)
		return err
	} else if sec == "g" {
		role, err := a.parseRelation(ptype, rule)
		if err != nil {
			return err
		}
		_, err = a.engine.InsertOne(role)
		return err
	} else {
		return errors.New("策略添加的sec非法! ")
	}
}

// RemovePolicy removes a policy rule from the storage.
func (a *adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	if sec == "p" {
		policy, err := a.parsePolicy(ptype, rule)
		if err != nil {
			return err
		}
		_, err = a.engine.Delete(&policy)
		return err
	} else if sec == "g" {
		role, err := a.parseRelation(ptype, rule)
		if err != nil {
			return err
		}
		_, err = a.engine.Delete(&role)
		return err
	} else {
		return errors.New("要删除的策略入参非法! ")
	}
}

// RemoveFilteredPolicy removes policy rules that match the filter from the storage.
func (a *adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	idx := fieldIndex + len(fieldValues)
	if sec == "p" {
		pt := 1
		if s := ptype[1:]; s != "" {
			pt, _ = strconv.Atoi(ptype[1:])
		}
		policy := Authorization{
			PolicyType: pt,
		}
		if fieldIndex <= 0 && idx > 0 {
			policy.SubjectId, _ = strconv.ParseInt(fieldValues[0-fieldIndex], 10, 64)
		}
		if fieldIndex <= 1 && idx > 1 {
			policy.Resource = fieldValues[1-fieldIndex]
		}
		if fieldIndex <= 2 && idx > 2 {
			policy.Action = fieldValues[2-fieldIndex]
		}
		if fieldIndex <= 3 && idx > 3 {
			es := fieldValues[3-fieldIndex]
			effect := 1
			if es == "deny" {
				effect = 2
			}
			policy.Effect = effect
		}
		if fieldIndex <= 3 && idx > 3 {
			policy.Priority, _ = strconv.Atoi(fieldValues[4-fieldIndex])
		}
		if fieldIndex <= 4 && idx > 4 {
			policy.TenantId, _ = strconv.ParseInt(fieldValues[5-fieldIndex], 10, 64)
		}
		if fieldIndex <= 5 && idx > 5 {
			policy.ResourceId, _ = strconv.ParseInt(fieldValues[6-fieldIndex], 10, 64)
		}
		_, err := a.engine.Delete(&policy)
		return err
	} else if sec == "g" {
		gt := 1
		if s := ptype[1:]; s != "" {
			gt, _ = strconv.Atoi(ptype[1:])
		}
		role := Relationship{
			RelationType: gt,
		}
		if fieldIndex <= 0 && idx > 0 {
			role.SubjectId, _ = strconv.ParseInt(fieldValues[0-fieldIndex], 10, 64)
		}
		if fieldIndex <= 1 && idx > 1 {
			role.ParentSubjectId, _ = strconv.ParseInt(fieldValues[1-fieldIndex], 10, 64)
		}
		if fieldIndex <= 2 && idx > 2 {
			role.TenantId, _ = strconv.ParseInt(fieldValues[2-fieldIndex], 10, 64)
		}
		_, err := a.engine.Delete(&role)
		return err
	} else {
		return errors.New("要删除的策略入参非法! ")
	}
}
