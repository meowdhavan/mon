package moon

import (
	"github.com/meowdhavan/moon/converter"
)

type flag struct {
	longName    string
	aliases     []string
	shortName   string
	about       string
	requiresVal bool
	setValue    func(string) error
	isValueSet  bool
	env         *string
	defaultVal  *string
	isRequired  bool
}

type flagOption func(*flag)

func Alias(alias string) flagOption {
	return func(f *flag) {
		f.aliases = append(f.aliases, alias)
	}
}

func Env(env string) flagOption {
	return func(f *flag) {
		*f.env = env
	}
}

func Default(defaultVal string) flagOption {
	return func(f *flag) {
		f.defaultVal = &defaultVal
	}
}

func Required() flagOption {
	return func(f *flag) {
		f.isRequired = true
	}
}

func (c *Command) StringFlag(target *string, longName string, shortName string, about string, options ...flagOption) {
	f := &flag{
		longName:    longName,
		aliases:     []string{},
		shortName:   shortName,
		about:       about,
		requiresVal: true,
		setValue: func(s string) error {
			v, err := converter.ToString(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}

func (c *Command) MultiStringFlag(target *[]string, longName string, shortName string, about string, options ...flagOption) {
	*target = []string{}

	f := &flag{
		longName:    longName,
		aliases:     []string{},
		shortName:   shortName,
		about:       about,
		requiresVal: true,
		setValue: func(s string) error {
			v, err := converter.ToString(s)
			if err != nil {
				return err
			}

			*target = append(*target, v)
			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}

func (c *Command) BoolFlag(target *bool, longName string, shortName string, about string, options ...flagOption) {
	*target = false

	f := &flag{
		longName:  longName,
		aliases:   []string{},
		shortName: shortName,
		about:     about,
		setValue: func(s string) error {
			v, err := converter.ToBool(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}

func (c *Command) MultiBoolFlag(target *int, longName string, shortName string, about string, options ...flagOption) {
	*target = 0

	f := &flag{
		longName:  longName,
		aliases:   []string{},
		shortName: shortName,
		about:     about,
		setValue: func(s string) error {
			v, err := converter.ToBool(s)
			if err != nil {
				return err
			}

			if v {
				*target++
			}

			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}

func (c *Command) IntFlag(target *int, longName string, shortName string, about string, options ...flagOption) {
	f := &flag{
		longName:  longName,
		aliases:   []string{},
		shortName: shortName,
		about:     about,

		requiresVal: true,
		setValue: func(s string) error {
			v, err := converter.ToInt(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}

func (c *Command) MultiIntFlag(target *[]int, longName string, shortName string, about string, options ...flagOption) {
	f := &flag{
		longName:    longName,
		aliases:     []string{},
		shortName:   shortName,
		about:       about,
		requiresVal: true,
		setValue: func(s string) error {
			v, err := converter.ToInt(s)
			if err != nil {
				return err
			}

			*target = append(*target, v)
			return nil
		},
	}

	for _, opt := range options {
		opt(f)
	}

	c.flags = append(c.flags, f)
}
