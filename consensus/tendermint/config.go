// Copyright 2015, 2016 Eris Industries (UK) Ltd.
// This file is part of Eris-RT

// Eris-RT is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Eris-RT is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Eris-RT.  If not, see <http://www.gnu.org/licenses/>.

// version provides the current Eris-DB version and a VersionIdentifier
// for the modules to identify their version with.

package tendermint

import (
  "time"

  tendermintConfig "github.com/tendermint/go-config"
  viper            "github.com/spf13/viper"
)

// NOTE [ben] Compiler check to ensure TendermintConfig successfully implements
// tendermint/go-config/config.Config
var _ tendermintConfig.Config = (*TendermintConfig)(nil)

// Tendermint has a self-rolled configuration type defined
// in tendermint/go-config but over an interface type, which is implemented
// by default in tendermint/tendermint/config/tendermint.go
// However, for Eris-DB purposes we can choose different rules for how to load
// the tendermint configuration and set the defaults.  Hence we re-implement
// go-config.Config on a viper subtree of the loaded Eris-DB configuration file.
type TendermintConfig struct {
  subTree *viper.Viper
}

// implement interface github.com/tendermint/go-config/config.Config
// so that `TMROOT` and config can be circumvented
func (tmintConfig *TendermintConfig) Get(key string) interface{} {
  return tmintConfig.subTree.Get(key)
}

func (tmintConfig *TendermintConfig) GetBool(key string) bool {
  return tmintConfig.subTree.GetBool(key)
}

func (tmintConfig *TendermintConfig) GetFloat64(key string) float64 {
  return tmintConfig.subTree.GetFloat64(key)
}

func (tmintConfig *TendermintConfig) GetInt(key string) int {
  return tmintConfig.subTree.GetInt(key)
}

func (tmintConfig *TendermintConfig) GetString(key string) string {
  return tmintConfig.subTree.GetString(key)
}

func (tmintConfig *TendermintConfig) GetStringSlice(key string) []string {
  return tmintConfig.subTree.GetStringSlice(key)
}

func (tmintConfig *TendermintConfig) GetTime(key string) time.Time {
  return tmintConfig.subTree.GetTime(key)
}

func (tmintConfig *TendermintConfig) GetMap(key string) map[string]interface{} {
  return tmintConfig.subTree.GetStringMap(key)
}

func (tmintConfig *TendermintConfig) GetMapString(key string) map[string]string {
  return tmintConfig.subTree.GetStringMapString(key)
}

func (tmintConfig *TendermintConfig) GetConfig(key string) tendermintConfig.Config {
  return &TendermintConfig {
    subTree: tmintConfig.subTree.Sub(key),
  }
}

func (tmintConfig *TendermintConfig) IsSet(key string) bool {
  return tmintConfig.IsSet(key)
}

func (tmintConfig *TendermintConfig) Set(key string, value interface{}) {
  tmintConfig.subTree.Set(key, value)
}

func (tmintConfig *TendermintConfig) SetDefault(key string, value interface{}) {
  tmintConfig.subTree.SetDefault(key, value)
}