package main

import (
	_ "embed"
	"flag"
	"fmt"
	"text/template"

	"google.golang.org/protobuf/compiler/protogen"
)

var version = "(unknown)"

//go:embed main.go.tpl
var mainTemplate string

type Config struct {
	SourcePath    string
	ProtocVersion string
	PluginVersion string
	GoPackageName string
	Services      []Service
}

type Service struct {
	Name    string
	Methods []Method
}

type Method struct {
	Name   string
	Input  string
	Output string
}

func main() {
	var flags flag.FlagSet

	options := protogen.Options{
		ParamFunc: flags.Set,
	}

	options.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if !file.Generate {
				continue
			}

			config := Config{
				SourcePath:    file.Desc.Path(),
				ProtocVersion: protocVersion(plugin),
				PluginVersion: version,
				GoPackageName: string(file.GoPackageName),
				Services:      []Service{},
			}

			for _, service := range file.Services {
				serviceTemplate := Service{
					Name:    service.GoName,
					Methods: []Method{},
				}

				for _, method := range service.Methods {
					methodTemplate := Method{
						Name:   method.GoName,
						Input:  method.Input.GoIdent.GoName,
						Output: method.Output.GoIdent.GoName,
					}

					serviceTemplate.Methods = append(serviceTemplate.Methods, methodTemplate)
				}

				config.Services = append(config.Services, serviceTemplate)
			}

			fileTemplate, err := template.New("main").Parse(mainTemplate)
			if err != nil {
				panic(err)
			}

			genFile := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+"_random.pb.go", file.GoImportPath)

			if err := fileTemplate.ExecuteTemplate(genFile, "main", config); err != nil {
				panic(err)
			}
		}

		return nil
	})
}

func protocVersion(gen *protogen.Plugin) string {
	version := gen.Request.GetCompilerVersion()
	if version == nil {
		return "(unknown)"
	}

	var suffix string
	if s := version.GetSuffix(); s != "" {
		suffix = "-" + s
	}

	return fmt.Sprintf("v%d.%d.%d%s", version.GetMajor(), version.GetMinor(), version.GetPatch(), suffix)
}
