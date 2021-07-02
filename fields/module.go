package fields

import (
	"text/template"

	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

func Module() *goFields {
	return &goFields{
		ModuleBase: &pgs.ModuleBase{},
	}
}

type goFields struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func (p *goFields) Name() string {
	return "go-fields"
}

func (p *goFields) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	tpl := template.New("fields").Funcs(map[string]interface{}{
		"package":     p.ctx.PackageName,
		"name":        p.ctx.Name,
	})
	p.tpl = template.Must(tpl.Parse(fieldsTpl))
}

func (p *goFields) Execute(targets map[string]pgs.File, _ map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		p.generate(f)
	}
	return p.Artifacts()
}

func (p *goFields) generate(f pgs.File) {
	if len(f.Messages()) == 0 {
		return
	}
	name := p.ctx.OutputPath(f).SetExt(".fields.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

const fieldsTpl = `package {{ package . }}

{{ range .AllMessages }}

var {{ name . }}Fields = struct {
	{{- range .Fields }}
	{{ name . }} string
	{{- end }}
}{
	{{- range .Fields }}
	{{ name . }}: "{{ .Name }}",
	{{- end }}
}

{{ end }}
`
