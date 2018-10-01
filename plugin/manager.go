package plugin

import (
	"io/ioutil"
	"log"
	"path"
	"plugin"
	"strings"

	"github.com/reynn/downloadr"
	"github.com/spf13/cobra"
)

type manager struct {
	path    string
	plugins map[string]downloadr.DownloadPlugin
	logger  downloadr.Logger
}

func New(p string, l downloadr.Logger) *manager {
	return &manager{
		path:    p,
		logger:  l,
		plugins: map[string]downloadr.DownloadPlugin{},
	}
}

func (m *manager) GetCount() int {
	return len(m.plugins)
}

func (m *manager) Gather(rootCmd *cobra.Command) error {
	ff, e := ioutil.ReadDir(m.path)
	if e != nil {
		return e
	}

	for _, f := range ff {
		if !strings.HasSuffix(f.Name(), ".so") || f.IsDir() {
			continue
		}
		p, e := plugin.Open(path.Join(m.path, f.Name()))
		if e != nil {
			log.Printf("failed to load plugin from file: %s [%v]", f.Name(), e)
			continue
		}
		s, e := p.Lookup("Plugin")
		if e != nil {
			m.logger.Error("plugin, %s, loaded but is not properly configured", f.Name())
			continue
		}

		plug, ok := s.(downloadr.DownloadPlugin)
		if !ok {
			m.logger.Error("plugin, %s, does not adhere do the DownloadPlugin interface", f.Name())
		} else {
			plug.Register(rootCmd, m.logger)
			m.plugins[plug.GetName()] = plug
		}
	}
	m.logger.Debug("Loaded %d plugins", len(m.plugins))
	return nil
}

func (m *manager) GetPlugins() []downloadr.DownloadPlugin {
	var ps []downloadr.DownloadPlugin
	for _, v := range m.plugins {
		ps = append(ps, v)
	}
	return ps
}