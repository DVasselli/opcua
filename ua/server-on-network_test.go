// Copyright 2018-2019 opcua authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ua

import (
	"testing"
)

func TestDecodeServerOnNetwork(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "single-cap",
			Struct: &ServerOnNetwork{
				RecordID:           1,
				ServerName:         "server-name",
				DiscoveryURL:       "discov-uri",
				ServerCapabilities: []string{"server-cap-1"},
			},
			// Struct: NewServerOnNetwork(
			// 	1,
			// 	"server-name",
			// 	"discov-uri",
			// 	[]string{"server-cap-1"},
			// ),
			Bytes: []byte{
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
			Name: "multiple-caps",
			Struct: &ServerOnNetwork{
				RecordID:           1,
				ServerName:         "server-name",
				DiscoveryURL:       "discov-uri",
				ServerCapabilities: []string{"server-cap-1", "server-cap-2"},
			},
			// Struct: NewServerOnNetwork(
			// 	1,
			// 	"server-name",
			// 	"discov-uri",
			// 	[]string{"server-cap-1", "server-cap-2"},
			// ),
			Bytes: []byte{
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

func TestServerOnNetworkArray(t *testing.T) {
	cases := []CodecTestCase{
		{
			Name: "normal",
			Struct: []*ServerOnNetwork{
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
			// Struct: []*ServerOnNetwork{
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
			// },
			Bytes: []byte{
				// ArraySize
				0x02, 0x00, 0x00, 0x00,
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
