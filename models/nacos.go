package models


type AKConfig struct {
        AKID string `mapstructure:"akid" json:"akid"`
        AKSK string     `mapstructure:"aksk" json:"aksk"`
}