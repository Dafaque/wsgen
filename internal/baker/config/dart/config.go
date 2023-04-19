package dart

import (
	"fmt"
	"wsgen/internal/baker/config"
	"wsgen/internal/baker/config/common"
)

type Config struct{}

func (c Config) GetTypeConverter() config.StringOverrider {
	return normalizeType
}
func (c Config) GetPublicFieldNameConverter() config.StringOverrider {
	return common.Nop
}
func (c Config) GetPublicStructNameConverter() config.StringOverrider {
	return normaizeStructName
}
func (c Config) GetCompleteMessage() string {
	var str string
	str += fmt.Sprintln("WIP")
	return str
}
