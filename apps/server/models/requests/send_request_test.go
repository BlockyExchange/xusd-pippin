package requests

import (
	"encoding/json"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"
)

func TestDecodeSendRequest(t *testing.T) {
	encoded := `{"action":"send","wallet":"1234","source":"xusd_1","destination":"xusd_2","bpow_key":"abc","amount":"1234"}`
	var decoded SendRequest
	json.Unmarshal([]byte(encoded), &decoded)
	assert.Equal(t, "send", decoded.Action)
	assert.Equal(t, "1234", decoded.Wallet)
	assert.Equal(t, "xusd_1", decoded.Source)
	assert.Equal(t, "xusd_2", decoded.Destination)
	assert.Equal(t, "abc", *decoded.BpowKey)
	assert.Equal(t, "1234", decoded.Amount)
	assert.Nil(t, decoded.Work)
}

func TestDecodeSendRequestNumericAmount(t *testing.T) {
	encoded := `{"action":"send","wallet":"1234","source":"xusd_1","destination":"xusd_2","bpow_key":"abc","amount":12340000000000000000000000000}`
	var decoded SendRequest
	json.Unmarshal([]byte(encoded), &decoded)
	assert.Equal(t, "send", decoded.Action)
	assert.Equal(t, "1234", decoded.Wallet)
	assert.Equal(t, "xusd_1", decoded.Source)
	assert.Equal(t, "xusd_2", decoded.Destination)
	assert.Equal(t, "abc", *decoded.BpowKey)
	assert.Equal(t, "12340000000000000000000000000", decoded.Amount)
	assert.Nil(t, decoded.Work)
}

func TestMapStructureDecodeSendRequest(t *testing.T) {
	request := map[string]interface{}{
		"action":      "send",
		"wallet":      "1234",
		"source":      "xusd_1",
		"destination": "xusd_2",
		"amount":      "1234",
		"bpow_key":    "abc",
	}
	var decoded SendRequest
	mapstructure.Decode(request, &decoded)
	assert.Equal(t, "send", decoded.Action)
	assert.Equal(t, "1234", decoded.Wallet)
	assert.Equal(t, "xusd_1", decoded.Source)
	assert.Equal(t, "xusd_2", decoded.Destination)
	assert.Equal(t, "1234", decoded.Amount)
	assert.Equal(t, "abc", *decoded.BpowKey)
	assert.Nil(t, decoded.Work)
}
