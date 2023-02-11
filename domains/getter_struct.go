package domains

import (
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
)

func (g getter) Struct() (dom *structureSpec.Domain, err error) {
	dom = &structureSpec.Domain{
		Id:          g.Id(),
		Name:        g.Name(),
		Description: g.Description(),
		Tags:        g.Tags(),
		Fqdn:        g.FQDN(),
		CertType:    g.Type(),
	}

	if dom.CertType == "inline" {
		dom.CertFile = g.Cert()
		dom.KeyFile = g.Key()
	}

	return
}
