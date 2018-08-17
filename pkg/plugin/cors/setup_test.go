package cors

import (
	"testing"

	"github.com/hellofresh/janus/pkg/plugin"
	"github.com/hellofresh/janus/pkg/proxy"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	var config Config
	rawConfig := map[string]interface{}{
		"domains":             []string{"*"},
		"methods":             []string{"GET"},
		"request_headers":     []string{"Content-Type", "Authorization"},
		"exposed_headers":     []string{"Test"},
		"options_passthrough": true,
	}

	err := plugin.Decode(rawConfig, &config)
	assert.NoError(t, err)

	assert.IsType(t, []string{}, config.AllowedOrigins)
	assert.Equal(t, []string{"*"}, config.AllowedOrigins)

	assert.IsType(t, []string{}, config.AllowedMethods)
	assert.Equal(t, []string{"GET"}, config.AllowedMethods)

	assert.IsType(t, []string{}, config.AllowedHeaders)
	assert.Equal(t, []string{"Content-Type", "Authorization"}, config.AllowedHeaders)

	assert.IsType(t, []string{}, config.ExposedHeaders)
	assert.Equal(t, []string{"Test"}, config.ExposedHeaders)

	assert.True(t, config.OptionsPassthrough)
}

func TestInvalidConfig(t *testing.T) {
	var config Config
	rawConfig := map[string]interface{}{
		"domains": "*",
	}

	err := plugin.Decode(rawConfig, &config)
	assert.Error(t, err)
}

func TestEmptyPassthrough(t *testing.T) {
	var config Config
	rawConfig := map[string]interface{}{
		"domains":         []string{"*"},
		"methods":         []string{"GET"},
		"request_headers": []string{"Content-Type", "Authorization"},
		"exposed_headers": []string{"Test"},
	}

	err := plugin.Decode(rawConfig, &config)
	assert.NoError(t, err)

	assert.False(t, config.OptionsPassthrough)
}

func TestSetup(t *testing.T) {
	rawConfig := map[string]interface{}{
		"domains":             []string{"*"},
		"methods":             []string{"GET"},
		"request_headers":     []string{"Content-Type", "Authorization"},
		"exposed_headers":     []string{"Test"},
		"options_passthrough": true,
	}
	def := proxy.NewRouterDefinition(proxy.NewDefinition())
	err := setupCors(def, rawConfig)

	assert.NoError(t, err)
	assert.Len(t, def.Middleware(), 1)
}
