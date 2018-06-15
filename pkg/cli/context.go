package cli

import (
	"crypto/tls"
	"io"
	"os/user"
	"path/filepath"

	"github.com/dcos/dcos-cli/pkg/config"
	"github.com/dcos/dcos-cli/pkg/httpclient"
	"github.com/dcos/dcos-cli/pkg/login"
	"github.com/dcos/dcos-cli/pkg/open"
	"github.com/dcos/dcos-cli/pkg/plugin"
	"github.com/dcos/dcos-cli/pkg/prompt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

// Context provides an implementation of api.Context. It relies on an Environment and is used to create
// various objects across the project and is being passed to every command as a constructor argument.
type Context struct {
	env    *Environment
	logger *logrus.Logger
}

// NewContext creates a new context from a given environment.
func NewContext(env *Environment) *Context {
	return &Context{env: env}
}

// Input returns the reader for CLI input.
func (ctx *Context) Input() io.Reader {
	return ctx.env.Input
}

// Out returns the writer for CLI output.
func (ctx *Context) Out() io.Writer {
	return ctx.env.Out
}

// ErrOut returns the writer for CLI errors, logs, and informational messages.
func (ctx *Context) ErrOut() io.Writer {
	return ctx.env.ErrOut
}

// EnvLookup lookups environment variables.
func (ctx *Context) EnvLookup(key string) (string, bool) {
	return ctx.env.EnvLookup(key)
}

// User returns the current system user.
func (ctx *Context) User() (*user.User, error) {
	return ctx.env.UserLookup()
}

// Fs returns the filesystem.
func (ctx *Context) Fs() afero.Fs {
	return ctx.env.Fs
}

// Logger returns the CLI logger.
func (ctx *Context) Logger() *logrus.Logger {
	if ctx.logger == nil {
		ctx.logger = &logrus.Logger{
			Out:       ctx.env.ErrOut,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
		}
	}
	return ctx.logger
}

// DCOSDir returns the root directory for the DC/OS CLI.
// It defaults to `~/.dcos` and can be overriden by the `DCOS_DIR` env var.
func (ctx *Context) DCOSDir() string {
	if dcosDir, ok := ctx.env.EnvLookup("DCOS_DIR"); ok {
		return dcosDir
	}
	if usr, err := ctx.env.UserLookup(); err == nil {
		return filepath.Join(usr.HomeDir, ".dcos")
	}

	// Not being able to get the current user is not critical. While it is
	// very unlikely to happen, we can fallback to the current directory.
	return ""
}

// ConfigManager returns the ConfigManager for the context.
func (ctx *Context) ConfigManager() *config.Manager {
	return config.NewManager(config.ManagerOpts{
		Fs:        ctx.env.Fs,
		EnvLookup: ctx.env.EnvLookup,
		Dir:       ctx.DCOSDir(),
	})
}

// Cluster returns the current cluster.
func (ctx *Context) Cluster() (*Cluster, error) {
	conf, err := ctx.ConfigManager().Current()
	if err != nil {
		return nil, err
	}
	return NewCluster(conf), nil
}

// Clusters returns the clusters.
func (ctx *Context) Clusters() []*Cluster {
	confs := ctx.ConfigManager().All()
	var clusters []*Cluster
	for _, conf := range confs {
		clusters = append(clusters, NewCluster(conf))
	}

	return clusters
}

// PluginManager creates and returns the PluginManager for the context
func (ctx *Context) PluginManager() (*plugin.Manager, error) {
	config, err := ctx.ConfigManager().Current()
	if err != nil {
		return nil, err
	}

	configHome := filepath.Dir(config.Path())
	dir := filepath.Join(configHome, "plugins")

	return plugin.NewManager(plugin.ManagerOpts{
		Fs:  ctx.env.Fs,
		Dir: dir,
	}), nil
}

// HTTPClient creates an httpclient.Client for a given cluster.
func (ctx *Context) HTTPClient(c *Cluster, opts ...httpclient.Option) *httpclient.Client {
	var baseOpts []httpclient.Option

	if c.ACSToken() != "" {
		baseOpts = append(baseOpts, httpclient.ACSToken(c.ACSToken()))
	}
	if c.Timeout() > 0 {
		baseOpts = append(baseOpts, httpclient.Timeout(c.Timeout()))
	}
	tlsOpt := httpclient.TLS(&tls.Config{
		InsecureSkipVerify: c.TLS().Insecure,
		RootCAs:            c.TLS().RootCAs,
	})

	baseOpts = append(baseOpts, tlsOpt, httpclient.Logger(ctx.Logger()))
	opts = append(baseOpts, opts...)

	return httpclient.New(c.URL(), opts...)
}

// Prompt is able to prompt for input, password or choices.
func (ctx *Context) Prompt() *prompt.Prompt {
	return prompt.New(ctx.Input(), ctx.Out())
}

// Opener returns a new OS Opener.
func (ctx *Context) Opener() open.Opener {
	return open.NewOsOpener(ctx.Logger())
}

// Login initiates a login based on a set of flags and HTTP client. On success it returns an ACS token.
func (ctx *Context) Login(flags *login.Flags, httpClient *httpclient.Client) (string, error) {
	flow := login.NewFlow(login.FlowOpts{
		Errout: ctx.ErrOut(),
		Prompt: ctx.Prompt(),
		Logger: ctx.Logger(),
		Opener: ctx.Opener(),
	})
	return flow.Start(flags, httpClient)
}
