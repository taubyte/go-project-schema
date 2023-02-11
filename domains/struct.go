package domains

import (
	"fmt"

	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
)

func (d *domain) SetWithStruct(sync bool, domain *structureSpec.Domain) error {
	ops := make([]basic.Op, 0)
	var opMapper = common.Mapper{
		{"Id", false, func() error {
			ops = append(ops, Id(domain.Id))
			return nil
		}},
		{"Description", false, func() error {
			ops = append(ops, Description(domain.Description))
			return nil
		}},
		{"Tags", false, func() error {
			ops = append(ops, Tags(domain.Tags))
			return nil
		}},
		{"FQDN", true, func() error {
			ops = append(ops, FQDN(domain.Fqdn))
			return nil
		}},
		{"Type", true, func() error {
			ops = append(ops, Type(domain.CertType))
			return nil
		}},

		// TODO handle auto cert ?
		{"CertFile", true, func() error {
			switch domain.CertType {
			case "inline":
				ops = append(ops, Cert(domain.CertFile))
			default:
				ops = append(ops, Cert(domain.CertFile))
			}
			return nil
		}},
		{"KeyFile", true, func() error {
			switch domain.CertType {
			case "inline":
				ops = append(ops, Key(domain.KeyFile))
			default:
				ops = append(ops, Key(domain.KeyFile))
			}
			return nil
		}},
		{"SmartOps", true, func() error {
			ops = append(ops, SmartOps(domain.SmartOps))
			return nil
		}},
	}

	err := opMapper.Run(domain)
	if err != nil {
		return fmt.Errorf("mapping values failed with: %s", err)
	}

	return d.Set(sync, ops...)
}
