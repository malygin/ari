package native

import "github.com/CyCoreSystems/ari"

type nativeBridge struct {
	conn *Conn
}

// Data returns the details of a bridge
// Equivalent to Get /bridges/{bridgeId}
func (b *nativeBridge) Data(id string) (bd ari.BridgeData, err error) {
	err = Get(b.conn, "/bridges/"+id, &bd)
	return
}

// AddChannel adds a channel to a bridge
// Equivalent to Post /bridges/{id}/addChannel
func (b *nativeBridge) AddChannel(bridgeID string, channelID string) (err error) {

	type request struct {
		ChannelID string `json:"channel"`
		Role      string `json:"role,omitempty"`
	}

	req := request{channelID, ""}
	err = Post(b.conn, "/bridges/"+bridgeID+"/addChannel", nil, &req)
	return
}

// RemoveChannel removes the specified channel from a bridge
// Equivalent to Post /bridges/{id}/removeChannel
func (b *nativeBridge) RemoveChannel(id string, channelID string) (err error) {
	req := struct {
		ChannelID string `json:"channel"`
	}{
		ChannelID: channelID,
	}

	//pass request
	err = Post(b.conn, "/bridges/"+id+"/removeChannel", nil, &req)
	return
}

// BridgeDelete shuts down a bridge. If any channels are in this bridge,
// they will be removed and resume whatever they were doing beforehand.
// This means that the channels themselves are not deleted.
// Equivalent to DELETE /bridges/{id}
func (b *nativeBridge) Delete(id string) (err error) {
	err = Delete(b.conn, "/bridges/"+id, nil, "")
	return
}