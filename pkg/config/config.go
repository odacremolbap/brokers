// Copyright 2022 TriggerMesh Inc.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"context"

	"gopkg.in/yaml.v3"
)

func Parse(config string) (*Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal([]byte(config), c); err != nil {
		return nil, err
	}

	if err := c.Validate(context.Background()); err != nil {
		return nil, err
	}

	return c, nil
}
