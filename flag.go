package moon

import "github.com/meowdhavan/moon/converter"

type flag struct {
	longNames   []string
	shortName   string
	about       string
	requiresVal bool
	isRequired  bool
	setValue    func(string) error
	isValueSet  bool
}

func (c *Command) AddStringFlag(target *string, longNames []string, shortName string, about string, isRequired bool) {
	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: true,
		isRequired: isRequired,
		setValue: func(s string) error {
			v, err := converter.ToString(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
		isValueSet: false,
	})
}

func (c *Command) AddMultiStringFlag(target *[]string, longNames []string, shortName string, about string) {
	*target = []string{}

	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: true,
		isRequired: false,
		setValue: func(s string) error {
			v, err := converter.ToString(s)
			if err != nil {
				return err
			}

			*target = append(*target, v)
			return nil
		},
		isValueSet: false,
	})
}

func (c *Command) AddBoolFlag(target *bool, longNames []string, shortName string, about string, isRequired bool) {
	*target = false

	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: false,
		isRequired: isRequired,
		setValue: func(s string) error {
			v, err := converter.ToBool(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
		isValueSet: false,
	})
}

func (c *Command) AddMultiBoolFlag(target *int, longNames []string, shortName string, about string) {
	*target = 0

	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: false,
		isRequired: false,
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
		isValueSet: false,
	})
}

func (c *Command) AddIntFlag(target *int, longNames []string, shortName string, about string, isRequired bool) {
	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: true,
		isRequired: isRequired,
		setValue: func(s string) error {
			v, err := converter.ToInt(s)
			if err != nil {
				return err
			}

			*target = v
			return nil
		},
		isValueSet: false,
	})
}

func (c *Command) AddMultiIntFlag(target *[]int, longNames []string, shortName string, about string) {
	c.flags = append(c.flags, &flag{
		longNames: longNames,
		shortName: shortName,
		about: about,
		requiresVal: true,
		isRequired: false,
		setValue: func(s string) error {
			v, err := converter.ToInt(s)
			if err != nil {
				return err
			}

			*target = append(*target, v)
			return nil
		},
		isValueSet: false,
	})
}