package webserver

import (
	"fmt"
	"github.com/solidworx/proj/templates"
	"github.com/spf13/afero"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"strings"
)

func AddConfig(host []string, ip string, port int, projectDir string) {
	var config_path interface{} = viper.Get("webservers.nginx.config_path")
	var appFs afero.Fs = afero.NewOsFs()

	var fs afero.File
	fs, _ = appFs.Create(fmt.Sprintf("%s/%s.conf", cast.ToString(config_path), projectDir))

	c, err := fs.WriteString(fmt.Sprintf(templates.PhpFpmDefault, strings.Join(addPortToHost(host, port), " "), projectDir))

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(c)
}

func addPortToHost(hosts []string, port int) []string {
	y := hosts[:0]
	for _, n := range hosts {
		y = append(y, n+":"+cast.ToString(port))
	}

	return y
}
