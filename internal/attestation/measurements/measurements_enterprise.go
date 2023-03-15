//go:build enterprise

/*
Copyright (c) Edgeless Systems GmbH

SPDX-License-Identifier: AGPL-3.0-only
*/

package measurements

import "github.com/edgelesssys/constellation/v2/internal/cloud/cloudprovider"

// Regenerate the measurements by running go generate.
// The enterprise build tag is required to validate the measurements using production
// sigstore certificates.
//go:generate go run -tags enterprise measurement-generator/generate.go

// DefaultsFor provides the default measurements for given cloud provider.
func DefaultsFor(provider cloudprovider.Provider) M {
	switch provider {
	case cloudprovider.AWS:
		return M{
			0: {
				Expected: [32]byte{
					0x73, 0x7f, 0x76, 0x7a, 0x12, 0xf5, 0x4e, 0x70,
					0xee, 0xcb, 0xc8, 0x68, 0x40, 0x11, 0x32, 0x3a,
					0xe2, 0xfe, 0x2d, 0xd9, 0xf9, 0x07, 0x85, 0x57,
					0x79, 0x69, 0xd7, 0xa2, 0x01, 0x3e, 0x8c, 0x12,
				},
				ValidationOpt: WarnOnly,
			},
			2: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			3: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			4: {
				Expected: [32]byte{
					0x4a, 0xd9, 0xdd, 0xb2, 0x58, 0xc6, 0x1e, 0x5f,
					0xd9, 0x5d, 0x5e, 0xe4, 0xde, 0xe4, 0xb1, 0xf6,
					0x02, 0xc7, 0x44, 0x1e, 0xab, 0x44, 0xd7, 0x29,
					0xc5, 0xd9, 0x3a, 0xa4, 0x77, 0xac, 0xd1, 0xfd,
				},
				ValidationOpt: Enforce,
			},
			5: {
				Expected: [32]byte{
					0x73, 0xd1, 0x3b, 0xf0, 0xa9, 0xdc, 0x9a, 0x20,
					0x1a, 0x33, 0x71, 0x50, 0xa0, 0xa4, 0xad, 0xb0,
					0x6e, 0x6b, 0x22, 0x06, 0xee, 0xfa, 0xeb, 0x4f,
					0xbb, 0x0f, 0xe5, 0xda, 0xc6, 0x20, 0x2f, 0xd3,
				},
				ValidationOpt: WarnOnly,
			},
			6: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			7: {
				Expected: [32]byte{
					0x12, 0x0e, 0x49, 0x8d, 0xb2, 0xa2, 0x24, 0xbd,
					0x51, 0x2b, 0x6e, 0xfc, 0x9b, 0x02, 0x34, 0xf8,
					0x43, 0xe1, 0x0b, 0xf0, 0x61, 0xeb, 0x7a, 0x76,
					0xec, 0xca, 0x55, 0x09, 0xa2, 0x23, 0x89, 0x01,
				},
				ValidationOpt: WarnOnly,
			},
			8: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			9: {
				Expected: [32]byte{
					0xa1, 0x3d, 0x79, 0x91, 0x0d, 0x9d, 0x98, 0x48,
					0x0c, 0x0d, 0x5c, 0x3f, 0x19, 0x7d, 0x38, 0x3d,
					0x5d, 0x26, 0x89, 0x5e, 0x35, 0x02, 0x25, 0xe2,
					0x19, 0xbf, 0x28, 0x6a, 0x74, 0x02, 0xe7, 0x80,
				},
				ValidationOpt: Enforce,
			},
			11: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			12: {
				Expected: [32]byte{
					0x0d, 0x01, 0x5e, 0x8d, 0xb4, 0xb9, 0x82, 0x33,
					0x00, 0x1c, 0x54, 0x5c, 0x8d, 0xc4, 0xc4, 0x1f,
					0x88, 0xa9, 0x02, 0xfd, 0x6c, 0x60, 0xfc, 0xd8,
					0x9c, 0x53, 0x46, 0xbb, 0x83, 0x75, 0xf0, 0x15,
				},
				ValidationOpt: Enforce,
			},
			13: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			14: {
				Expected: [32]byte{
					0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22,
					0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9,
					0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c,
					0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f,
				},
				ValidationOpt: WarnOnly,
			},
			15: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
		}
	case cloudprovider.Azure:
		return M{
			1: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			2: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			3: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			4: {
				Expected: [32]byte{
					0xd2, 0x3e, 0x29, 0x3b, 0x73, 0xe5, 0xa9, 0xcd,
					0x2b, 0xdd, 0x51, 0x67, 0x62, 0xf3, 0x08, 0x9f,
					0x7f, 0x16, 0x13, 0xb7, 0xd3, 0x78, 0x90, 0xaa,
					0x0c, 0x3e, 0xe0, 0xb8, 0xc4, 0xaf, 0x5e, 0x5c,
				},
				ValidationOpt: Enforce,
			},
			5: {
				Expected: [32]byte{
					0xb7, 0x31, 0xe1, 0xa1, 0xc7, 0x0e, 0x65, 0xe4,
					0x3a, 0x4e, 0xd5, 0x4d, 0x66, 0x8b, 0x7e, 0x58,
					0x51, 0x63, 0xd9, 0xc4, 0x2f, 0x1e, 0xd4, 0x09,
					0x61, 0x74, 0x45, 0x1f, 0xf0, 0xba, 0x62, 0xf2,
				},
				ValidationOpt: WarnOnly,
			},
			7: {
				Expected: [32]byte{
					0x34, 0x65, 0x47, 0xa8, 0xce, 0x59, 0x57, 0xaf,
					0x27, 0xe5, 0x52, 0x42, 0x7d, 0x6b, 0x9e, 0x6d,
					0x9c, 0xb5, 0x02, 0xf0, 0x15, 0x6e, 0x91, 0x55,
					0x38, 0x04, 0x51, 0xee, 0xa1, 0xb3, 0xf0, 0xed,
				},
				ValidationOpt: WarnOnly,
			},
			8: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			9: {
				Expected: [32]byte{
					0xa1, 0x3d, 0x79, 0x91, 0x0d, 0x9d, 0x98, 0x48,
					0x0c, 0x0d, 0x5c, 0x3f, 0x19, 0x7d, 0x38, 0x3d,
					0x5d, 0x26, 0x89, 0x5e, 0x35, 0x02, 0x25, 0xe2,
					0x19, 0xbf, 0x28, 0x6a, 0x74, 0x02, 0xe7, 0x80,
				},
				ValidationOpt: Enforce,
			},
			11: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			12: {
				Expected: [32]byte{
					0xfd, 0x80, 0xe6, 0xd3, 0x69, 0x2d, 0x7d, 0x39,
					0x10, 0x4a, 0x00, 0x58, 0x90, 0x5f, 0xeb, 0xa2,
					0xdf, 0x5b, 0xc5, 0x95, 0x58, 0x6b, 0xfa, 0x0f,
					0xbd, 0xdf, 0x73, 0x8b, 0x0b, 0x2d, 0x47, 0xe2,
				},
				ValidationOpt: Enforce,
			},
			13: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			14: {
				Expected: [32]byte{
					0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22,
					0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9,
					0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c,
					0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f,
				},
				ValidationOpt: WarnOnly,
			},
			15: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
		}
	case cloudprovider.GCP:
		return M{
			1: {
				Expected: [32]byte{
					0x74, 0x5f, 0x2f, 0xb4, 0x23, 0x5e, 0x46, 0x47,
					0xaa, 0x0a, 0xd5, 0xac, 0xe7, 0x81, 0xcd, 0x92,
					0x9e, 0xb6, 0x8c, 0x28, 0x87, 0x0e, 0x7d, 0xd5,
					0xd1, 0xa1, 0x53, 0x58, 0x54, 0x32, 0x5e, 0x56,
				},
				ValidationOpt: WarnOnly,
			},
			2: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			3: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			4: {
				Expected: [32]byte{
					0x97, 0x65, 0x24, 0x75, 0x9a, 0x1f, 0x06, 0x8e,
					0x75, 0xa5, 0x81, 0x96, 0x5f, 0xa6, 0xdd, 0xb3,
					0x7c, 0x9b, 0x62, 0x60, 0x1c, 0xd6, 0x79, 0x1e,
					0x28, 0xb6, 0xc9, 0x75, 0xda, 0x92, 0xfc, 0xab,
				},
				ValidationOpt: Enforce,
			},
			5: {
				Expected: [32]byte{
					0x9d, 0x5f, 0x8c, 0x2e, 0x04, 0x64, 0xc6, 0x95,
					0x52, 0xe5, 0x93, 0xbe, 0x32, 0xa4, 0xb0, 0x89,
					0xf7, 0xe7, 0x7f, 0x4e, 0xd1, 0x2d, 0x60, 0xcd,
					0xcd, 0x94, 0x12, 0x0b, 0x6c, 0xca, 0xcc, 0xf8,
				},
				ValidationOpt: WarnOnly,
			},
			6: {
				Expected: [32]byte{
					0x3d, 0x45, 0x8c, 0xfe, 0x55, 0xcc, 0x03, 0xea,
					0x1f, 0x44, 0x3f, 0x15, 0x62, 0xbe, 0xec, 0x8d,
					0xf5, 0x1c, 0x75, 0xe1, 0x4a, 0x9f, 0xcf, 0x9a,
					0x72, 0x34, 0xa1, 0x3f, 0x19, 0x8e, 0x79, 0x69,
				},
				ValidationOpt: WarnOnly,
			},
			7: {
				Expected: [32]byte{
					0xb1, 0xe9, 0xb3, 0x05, 0x32, 0x5c, 0x51, 0xb9,
					0x3d, 0xa5, 0x8c, 0xbf, 0x7f, 0x92, 0x51, 0x2d,
					0x8e, 0xeb, 0xfa, 0x01, 0x14, 0x3e, 0x4d, 0x88,
					0x44, 0xe4, 0x0e, 0x06, 0x2e, 0x9b, 0x6c, 0xd5,
				},
				ValidationOpt: WarnOnly,
			},
			8: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			9: {
				Expected: [32]byte{
					0xcb, 0x82, 0x43, 0x95, 0x0b, 0x6f, 0xc8, 0xef,
					0x42, 0x17, 0xde, 0xda, 0xc9, 0xae, 0xc5, 0x05,
					0xf8, 0x09, 0x84, 0x43, 0x54, 0xf2, 0xf6, 0x38,
					0xcd, 0x47, 0xcd, 0x7a, 0xc5, 0xf1, 0x72, 0xe3,
				},
				ValidationOpt: Enforce,
			},
			11: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			12: {
				Expected: [32]byte{
					0x31, 0xb2, 0xc4, 0xc7, 0x92, 0x01, 0xc1, 0xc0,
					0x9a, 0x3e, 0x24, 0x0c, 0xe6, 0x5c, 0xca, 0x90,
					0x23, 0xc7, 0x79, 0x03, 0x95, 0x25, 0x21, 0xf1,
					0x2f, 0x5b, 0x8a, 0x64, 0x9c, 0x65, 0xca, 0x4a,
				},
				ValidationOpt: Enforce,
			},
			13: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			14: {
				Expected: [32]byte{
					0xd7, 0xc4, 0xcc, 0x7f, 0xf7, 0x93, 0x30, 0x22,
					0xf0, 0x13, 0xe0, 0x3b, 0xde, 0xe8, 0x75, 0xb9,
					0x17, 0x20, 0xb5, 0xb8, 0x6c, 0xf1, 0x75, 0x3c,
					0xad, 0x83, 0x0f, 0x95, 0xe7, 0x91, 0x92, 0x6f,
				},
				ValidationOpt: WarnOnly,
			},
			15: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
		}
	case cloudprovider.QEMU:
		return M{
			4: {
				Expected: [32]byte{
					0xdf, 0xd6, 0x2a, 0x25, 0x1a, 0x23, 0x4d, 0x2e,
					0xcc, 0xdb, 0xb4, 0x65, 0x9e, 0x49, 0x90, 0xc3,
					0x30, 0xf3, 0xe4, 0xa4, 0x45, 0x6f, 0x72, 0x74,
					0xc7, 0x69, 0xad, 0x43, 0xc1, 0x74, 0xc7, 0xc4,
				},
				ValidationOpt: Enforce,
			},
			8: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			9: {
				Expected: [32]byte{
					0xa1, 0x3d, 0x79, 0x91, 0x0d, 0x9d, 0x98, 0x48,
					0x0c, 0x0d, 0x5c, 0x3f, 0x19, 0x7d, 0x38, 0x3d,
					0x5d, 0x26, 0x89, 0x5e, 0x35, 0x02, 0x25, 0xe2,
					0x19, 0xbf, 0x28, 0x6a, 0x74, 0x02, 0xe7, 0x80,
				},
				ValidationOpt: Enforce,
			},
			11: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			12: {
				Expected: [32]byte{
					0xc2, 0xb1, 0xe2, 0x4a, 0x8f, 0xf7, 0x34, 0xea,
					0x95, 0x46, 0x62, 0x2d, 0xa8, 0xba, 0x43, 0x8c,
					0x67, 0x99, 0xbe, 0x6b, 0x6c, 0xb4, 0xe6, 0x59,
					0x3d, 0x94, 0x11, 0xe9, 0xea, 0x08, 0x4c, 0x0c,
				},
				ValidationOpt: Enforce,
			},
			13: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
			15: {
				Expected: [32]byte{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				},
				ValidationOpt: Enforce,
			},
		}
	default:
		return nil
	}
}
