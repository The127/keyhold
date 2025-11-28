package config

import (
	"fmt"
	"keyhold/internal/args"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
}

var C Config

var k = koanf.New(".")

func Init() error {
	if args.ConfigFilePath() != "" {
		_, err := os.Stat(args.ConfigFilePath())
		if err != nil {
			return fmt.Errorf("config file not found: %w", err)
		}

		err = k.Load(file.Provider(args.ConfigFilePath()), yaml.Parser())
		if err != nil {
			return fmt.Errorf("failed to load config file: %w", err)
		}
	}

	err := k.Load(env.Provider(".", env.Opt{
		Prefix: "KEYHOLD_",
		TransformFunc: func(k, v string) (string, any) {
			// Transform the key.
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "KEYHOLD_")), "_", ".")

			if strings.Contains(v, " ") {
				return k, strings.Split(v, " ")
			}

			return k, v
		},
	}), nil)
	if err != nil {
		return fmt.Errorf("failed to load env config: %w", err)
	}

	err = k.Unmarshal("", &C)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	err = setDefaults()
	if err != nil {
		return fmt.Errorf("failed to set defaults: %w", err)
	}

	return nil
}

func setDefaults() error {
	return nil
}
