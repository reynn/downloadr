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
	plugins map[string]downloadr.ScraperPlugin
	logger  downloadr.Logger
}

func (m *manager) GetScraperPlugins() []downloadr.ScraperPlugin {
	var ps []downloadr.ScraperPlugin
	for _, v := range m.plugins {
		ps = append(ps, v)
	}
	return ps
}

func New(p string, l downloadr.Logger) *manager {
	return &manager{
		path:    p,
		logger:  l,
		plugins: map[string]downloadr.ScraperPlugin{},
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

		plug, ok := s.(downloadr.ScraperPlugin)
		if !ok {
			m.logger.Error("plugin, %s, does not adhere do the DownloadPlugin interface", f.Name())
		} else {
			plug.Register(rootCmd, m.logger)
			m.plugins[plug.Name()] = plug
		}
	}
	m.logger.Debug("Loaded %d plugins", len(m.plugins))
	return nil
}
