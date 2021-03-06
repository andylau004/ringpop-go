package ringpop

import "errors"

var (
	// ErrNotBootstrapped is returned by public methods which require the ring to
	// be bootstrapped before they can operate correctly.
	ErrNotBootstrapped = errors.New("ringpop is not bootstrapped")

	// ErrEphemeralIdentity is returned by the identity resolver if TChannel is
	// using port 0 and is not listening (and thus has not been assigned a port by
	// the OS).
	ErrEphemeralIdentity = errors.New("unable to resolve this node's identity from channel that is not yet listening")

	// ErrChannelNotListening is returned on bootstrap if TChannel is not
	// listening.
	ErrChannelNotListening = errors.New("tchannel is not listening")
)
