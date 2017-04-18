package arimocks

import ari "github.com/CyCoreSystems/ari"
import mock "github.com/stretchr/testify/mock"
import time "time"

// Channel is an autogenerated mock type for the Channel type
type Channel struct {
	mock.Mock
}

// Answer provides a mock function with given fields: key
func (_m *Channel) Answer(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Busy provides a mock function with given fields: key
func (_m *Channel) Busy(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Congestion provides a mock function with given fields: key
func (_m *Channel) Congestion(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Continue provides a mock function with given fields: key, context, extension, priority
func (_m *Channel) Continue(key *ari.Key, context string, extension string, priority int) error {
	ret := _m.Called(key, context, extension, priority)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string, string, int) error); ok {
		r0 = rf(key, context, extension, priority)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0
func (_m *Channel) Create(_a0 ari.ChannelCreateRequest) (ari.ChannelHandle, error) {
	ret := _m.Called(_a0)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(ari.ChannelCreateRequest) ari.ChannelHandle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ari.ChannelCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Data provides a mock function with given fields: key
func (_m *Channel) Data(key *ari.Key) (*ari.ChannelData, error) {
	ret := _m.Called(key)

	var r0 *ari.ChannelData
	if rf, ok := ret.Get(0).(func(*ari.Key) *ari.ChannelData); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ari.ChannelData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Dial provides a mock function with given fields: key, caller, timeout
func (_m *Channel) Dial(key *ari.Key, caller string, timeout time.Duration) error {
	ret := _m.Called(key, caller, timeout)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string, time.Duration) error); ok {
		r0 = rf(key, caller, timeout)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *Channel) Get(key *ari.Key) ari.ChannelHandle {
	ret := _m.Called(key)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(*ari.Key) ari.ChannelHandle); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	return r0
}

// Hangup provides a mock function with given fields: key, reason
func (_m *Channel) Hangup(key *ari.Key, reason string) error {
	ret := _m.Called(key, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Hold provides a mock function with given fields: key
func (_m *Channel) Hold(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: _a0
func (_m *Channel) List(_a0 *ari.Key) ([]*ari.Key, error) {
	ret := _m.Called(_a0)

	var r0 []*ari.Key
	if rf, ok := ret.Get(0).(func(*ari.Key) []*ari.Key); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ari.Key)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MOH provides a mock function with given fields: key, moh
func (_m *Channel) MOH(key *ari.Key, moh string) error {
	ret := _m.Called(key, moh)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string) error); ok {
		r0 = rf(key, moh)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Mute provides a mock function with given fields: key, dir
func (_m *Channel) Mute(key *ari.Key, dir ari.Direction) error {
	ret := _m.Called(key, dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, ari.Direction) error); ok {
		r0 = rf(key, dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Originate provides a mock function with given fields: _a0
func (_m *Channel) Originate(_a0 ari.OriginateRequest) (ari.ChannelHandle, error) {
	ret := _m.Called(_a0)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(ari.OriginateRequest) ari.ChannelHandle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(ari.OriginateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Play provides a mock function with given fields: key, playbackID, mediaURI
func (_m *Channel) Play(key *ari.Key, playbackID string, mediaURI string) (ari.PlaybackHandle, error) {
	ret := _m.Called(key, playbackID, mediaURI)

	var r0 ari.PlaybackHandle
	if rf, ok := ret.Get(0).(func(*ari.Key, string, string) ari.PlaybackHandle); ok {
		r0 = rf(key, playbackID, mediaURI)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.PlaybackHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key, string, string) error); ok {
		r1 = rf(key, playbackID, mediaURI)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Record provides a mock function with given fields: key, name, opts
func (_m *Channel) Record(key *ari.Key, name string, opts *ari.RecordingOptions) (ari.LiveRecordingHandle, error) {
	ret := _m.Called(key, name, opts)

	var r0 ari.LiveRecordingHandle
	if rf, ok := ret.Get(0).(func(*ari.Key, string, *ari.RecordingOptions) ari.LiveRecordingHandle); ok {
		r0 = rf(key, name, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.LiveRecordingHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key, string, *ari.RecordingOptions) error); ok {
		r1 = rf(key, name, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ring provides a mock function with given fields: key
func (_m *Channel) Ring(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendDTMF provides a mock function with given fields: key, dtmf, opts
func (_m *Channel) SendDTMF(key *ari.Key, dtmf string, opts *ari.DTMFOptions) error {
	ret := _m.Called(key, dtmf, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, string, *ari.DTMFOptions) error); ok {
		r0 = rf(key, dtmf, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Silence provides a mock function with given fields: key
func (_m *Channel) Silence(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Snoop provides a mock function with given fields: key, snoopID, opts
func (_m *Channel) Snoop(key *ari.Key, snoopID string, opts *ari.SnoopOptions) (ari.ChannelHandle, error) {
	ret := _m.Called(key, snoopID, opts)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(*ari.Key, string, *ari.SnoopOptions) ari.ChannelHandle); ok {
		r0 = rf(key, snoopID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ari.Key, string, *ari.SnoopOptions) error); ok {
		r1 = rf(key, snoopID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StageOriginate provides a mock function with given fields: _a0
func (_m *Channel) StageOriginate(_a0 ari.OriginateRequest) ari.ChannelHandle {
	ret := _m.Called(_a0)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(ari.OriginateRequest) ari.ChannelHandle); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	return r0
}

// StageSnoop provides a mock function with given fields: key, snoopID, opts
func (_m *Channel) StageSnoop(key *ari.Key, snoopID string, opts *ari.SnoopOptions) ari.ChannelHandle {
	ret := _m.Called(key, snoopID, opts)

	var r0 ari.ChannelHandle
	if rf, ok := ret.Get(0).(func(*ari.Key, string, *ari.SnoopOptions) ari.ChannelHandle); ok {
		r0 = rf(key, snoopID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.ChannelHandle)
		}
	}

	return r0
}

// StopHold provides a mock function with given fields: key
func (_m *Channel) StopHold(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopMOH provides a mock function with given fields: key
func (_m *Channel) StopMOH(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopRing provides a mock function with given fields: key
func (_m *Channel) StopRing(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopSilence provides a mock function with given fields: key
func (_m *Channel) StopSilence(key *ari.Key) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: key, n
func (_m *Channel) Subscribe(key *ari.Key, n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(*ari.Key, ...string) ari.Subscription); ok {
		r0 = rf(key, n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}

// Unmute provides a mock function with given fields: key, dir
func (_m *Channel) Unmute(key *ari.Key, dir ari.Direction) error {
	ret := _m.Called(key, dir)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ari.Key, ari.Direction) error); ok {
		r0 = rf(key, dir)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Variables provides a mock function with given fields: key
func (_m *Channel) Variables(key *ari.Key) ari.Variables {
	ret := _m.Called(key)

	var r0 ari.Variables
	if rf, ok := ret.Get(0).(func(*ari.Key) ari.Variables); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Variables)
		}
	}

	return r0
}
