// Copyright 2021 Linka Cloud  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fields

import (
	"strings"
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
		"package": p.ctx.PackageName,
		"name":    p.ctx.Name,
		"comment": func(s string) string {
			var out string
			parts := strings.Split(s, "\n")
			for i, v := range parts {
				if i == len(parts)-1 && v == "" {
					return out
				}
				out += "//" + v + "\n"
			}
			return out
		},
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
	if len(f.Messages()) == 0 && len(f.Services()) == 0 {
		return
	}
	name := p.ctx.OutputPath(f).SetExt(".fields.go")
	p.AddGeneratorTemplateFile(name.String(), p.tpl, f)
}

const fieldsTpl = `{{ comment .SyntaxSourceCodeInfo.LeadingComments }}
{{ range .SyntaxSourceCodeInfo.LeadingDetachedComments }}
{{ comment . }}
{{ end }}
// Code generated by protoc-gen-defaults. DO NOT EDIT.

package {{ package . }}

{{ $file := . }}
{{ range .Services }}
{{- $svc := . }} 
var {{ .Name }}Methods = struct {
	{{- range .Methods }}
	{{ name . }} string
	{{- end }}
}{
	{{- range .Methods }}
	{{ name . }}: "/{{ $file.Package.ProtoName }}.{{ $svc.Name }}/{{ .Name }}",
	{{- end }}
}
{{ end }}

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
