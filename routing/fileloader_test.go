package routing

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vaulty/vaulty/encryption/noneenc"
	"github.com/vaulty/vaulty/secrets/memorystorage"
	"github.com/vaulty/vaulty/transformer"
	"github.com/vaulty/vaulty/transformer/json"
)

func TestLoadFromFile(t *testing.T) {
	enc := noneenc.New()

	secretsStorage := memorystorage.New(&memorystorage.Params{
		Encrypter: enc,
	})

	loader := &fileLoader{
		enc:            enc,
		secretsStorage: secretsStorage,
		transformerFactory: map[string]transformer.Factory{
			"json": json.Factory,
		},
	}

	routes, err := loader.Load("./testdata/routes.json")
	require.NoError(t, err)

	require.Equal(t, "in1", routes[0].Name)
	require.Equal(t, "in2", routes[1].Name)
	require.Equal(t, "inAll", routes[2].Name)
	require.Equal(t, "out1", routes[3].Name)
	require.Equal(t, "outAll", routes[4].Name)
}
