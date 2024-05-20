package paginator

type Flag string

const (
	ENABLED  Flag = "ENABLED"
	DISABLED Flag = "DISABLED"
)

var defaultConfig = Config{
	Keys:     []string{"ID"},
	Limit:    10,
	Order:    DESC,
	TupleCmp: DISABLED,
}

// Option for paginator
type Option interface {
	Apply(p *Paginator)
}

// Config for paginator
type Config struct {
	Rules    []Rule
	Keys     []string
	Limit    int
	Order    Order
	After    string
	Before   string
	TupleCmp Flag
}

// Apply applies config to paginator
func (c *Config) Apply(p *Paginator) {
	if c.Rules != nil {
		p.SetRules(c.Rules...)
	}
	// only set keys when no rules presented
	if c.Rules == nil && c.Keys != nil {
		p.SetKeys(c.Keys...)
	}
	if c.Limit != 0 {
		p.SetLimit(c.Limit)
	}
	if c.Order != "" {
		p.SetOrder(c.Order)
	}
	if c.After != "" {
		p.SetAfterCursor(c.After)
	}
	if c.Before != "" {
		p.SetBeforeCursor(c.Before)
	}
	if c.TupleCmp != "" {
		p.SetAllowTupleCmp(c.TupleCmp == ENABLED)
	}
}

// WithRules configures rules for paginator
func WithRules(rules ...Rule) Option {
	return &Config{
		Rules: rules,
	}
}

// WithKeys configures keys for paginator
func WithKeys(keys ...string) Option {
	return &Config{
		Keys: keys,
	}
}

// WithLimit configures limit for paginator
func WithLimit(limit int) Option {
	return &Config{
		Limit: limit,
	}
}

// WithOrder configures order for paginator
func WithOrder(order Order) Option {
	return &Config{
		Order: order,
	}
}

// WithAfter configures after cursor for paginator
func WithAfter(c string) Option {
	return &Config{
		After: c,
	}
}

// WithBefore configures before cursor for paginator
func WithBefore(c string) Option {
	return &Config{
		Before: c,
	}
}

// WithAllowTupleCmp enables tuple comparison optimization
func WithTupleCmp(flag Flag) Option {
	return &Config{
		TupleCmp: flag,
	}
}
