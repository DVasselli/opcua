// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"testing"
	"time"
)

func TestFindServersOnNetworkResponse(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "single-server",
			Struct: &FindServersOnNetworkResponse{
				ResponseHeader: &ResponseHeader{
					Timestamp:          time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					RequestHandle:      1,
					ServiceDiagnostics: &DiagnosticInfo{},
					StringTable:        []string{},
					AdditionalHeader:   NewExtensionObject(nil),
				},
				LastCounterResetTime: time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
				Servers: []*ServerOnNetwork{
					&ServerOnNetwork{
						RecordID:           1,
						ServerName:         "server-name",
						DiscoveryURL:       "discov-uri",
						ServerCapabilities: []string{"server-cap-1"},
					},
				},
			},
			// Struct: NewFindServersOnNetworkResponse(
			// 	NewResponseHeader(
			// 		time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
			// 		1, 0, NewNullDiagnosticInfo(), []string{}, NewExtensionObject(nil),
			// 	),
			// 	time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
			// 	NewServerOnNetwork(
			// 		1,
			// 		"server-name",
			// 		"discov-uri",
			// 		[]string{"server-cap-1"},
			// 	),
			// ),
			Bytes: []byte{
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ServiceResult
				0x00, 0x00, 0x00, 0x00,
				// ServiceDiagnostics
				0x00,
				// StringTable
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// LastCounterResetTime
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// Servers
				// ArraySize
				0x01, 0x00, 0x00, 0x00,
				// ApplicationURI
				// RecordID
				0x01, 0x00, 0x00, 0x00,
				// ServerName
				0x0b, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
				// DiscoveryURI
				0x0a, 0x00, 0x00, 0x00,
				0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69,
				// ServerCapabilities
				0x01, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x61, 0x70, 0x2d, 0x31,
			},
		},
		{
			Name: "multiple-servers",
			Struct: &FindServersOnNetworkResponse{
				ResponseHeader: &ResponseHeader{
					Timestamp:          time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
					RequestHandle:      1,
					ServiceDiagnostics: &DiagnosticInfo{},
					StringTable:        []string{},
					AdditionalHeader:   NewExtensionObject(nil),
				},
				LastCounterResetTime: time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
				Servers: []*ServerOnNetwork{
					&ServerOnNetwork{
						RecordID:           1,
						ServerName:         "server-name",
						DiscoveryURL:       "discov-uri",
						ServerCapabilities: []string{"server-cap-1"},
					},
					&ServerOnNetwork{
						RecordID:           1,
						ServerName:         "server-name",
						DiscoveryURL:       "discov-uri",
						ServerCapabilities: []string{"server-cap-1", "server-cap-2"},
					},
				},
			},
			// Struct: NewFindServersOnNetworkResponse(
			// 	NewResponseHeader(
			// 		time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
			// 		1, 0, NewNullDiagnosticInfo(), []string{}, NewExtensionObject(nil),
			// 	),
			// 	time.Date(2018, time.August, 10, 23, 0, 0, 0, time.UTC),
			// 	NewServerOnNetwork(
			// 		1,
			// 		"server-name",
			// 		"discov-uri",
			// 		[]string{"server-cap-1"},
			// 	),
			// 	NewServerOnNetwork(
			// 		1,
			// 		"server-name",
			// 		"discov-uri",
			// 		[]string{"server-cap-1", "server-cap-2"},
			// 	),
			// ),
			Bytes: []byte{
				// Timestamp
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// RequestHandle
				0x01, 0x00, 0x00, 0x00,
				// ServiceResult
				0x00, 0x00, 0x00, 0x00,
				// ServiceDiagnostics
				0x00,
				// StringTable
				0x00, 0x00, 0x00, 0x00,
				// AdditionalHeader
				0x00, 0x00, 0x00,
				// LastCounterResetTime
				0x00, 0x98, 0x67, 0xdd, 0xfd, 0x30, 0xd4, 0x01,
				// Servers
				// ArraySize
				0x02, 0x00, 0x00, 0x00,
				// ApplicationURI
				// RecordID
				0x01, 0x00, 0x00, 0x00,
				// ServerName
				0x0b, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
				// DiscoveryURI
				0x0a, 0x00, 0x00, 0x00,
				0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69,
				// ServerCapabilities
				0x01, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x61, 0x70, 0x2d, 0x31,
				// RecordID
				0x01, 0x00, 0x00, 0x00,
				// ServerName
				0x0b, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x6e, 0x61, 0x6d, 0x65,
				// DiscoveryURI
				0x0a, 0x00, 0x00, 0x00,
				0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x2d, 0x75, 0x72, 0x69,
				// ServerCapabilities
				0x02, 0x00, 0x00, 0x00,
				0x0c, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x61, 0x70, 0x2d, 0x31,
				0x0c, 0x00, 0x00, 0x00,
				0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x63, 0x61, 0x70, 0x2d, 0x32,
			},
		},
	}
	RunCodecTest(t, cases)
}
