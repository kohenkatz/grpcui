package standalone

import (
	"bytes"
	"compress/gzip"
	"io"
)

// GetIndexTemplate returns raw, uncompressed file data.
func GetIndexTemplate() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x94, 0x56,
		0xdb, 0x6e, 0xdb, 0x46, 0x10, 0x7d, 0xcf, 0x57, 0x0c, 0xd6, 0x28, 0x90,
		0x14, 0xbc, 0x88, 0x94, 0x1d, 0xb4, 0x0c, 0x25, 0xa0, 0x30, 0xd0, 0xcb,
		0x4b, 0x11, 0x18, 0x2d, 0xf2, 0xbc, 0x5c, 0x0e, 0xc9, 0x8d, 0x97, 0xbb,
		0xcc, 0xee, 0x48, 0x96, 0x22, 0xe8, 0xdf, 0x0b, 0x5e, 0x25, 0x4a, 0x32,
		0x1a, 0xbf, 0xd8, 0xe4, 0x5c, 0xce, 0x39, 0x9c, 0x99, 0x9d, 0x55, 0x5a,
		0x51, 0xad, 0xd6, 0xef, 0x00, 0xd2, 0x0a, 0x79, 0xde, 0x3e, 0x00, 0xa4,
		0x24, 0x49, 0xe1, 0xba, 0x7c, 0xfa, 0xfc, 0x08, 0xff, 0xfe, 0x95, 0x86,
		0xfd, 0xeb, 0xbb, 0xde, 0xa7, 0xa4, 0x7e, 0x06, 0x8b, 0x6a, 0xc5, 0xa4,
		0x30, 0x9a, 0x01, 0xed, 0x1b, 0x5c, 0x31, 0x59, 0xf3, 0x12, 0xc3, 0x46,
		0x97, 0x0c, 0x2a, 0x8b, 0xc5, 0x8a, 0x15, 0x7c, 0xdb, 0xfa, 0x83, 0xd6,
		0x34, 0xa6, 0x3a, 0x61, 0x65, 0x43, 0xe0, 0xac, 0x58, 0xb1, 0xaf, 0xdf,
		0x36, 0x68, 0xf7, 0xfe, 0x32, 0x58, 0x06, 0x51, 0x50, 0x4b, 0x1d, 0x7c,
		0x75, 0x6c, 0x9d, 0x86, 0x7d, 0xc8, 0x35, 0x97, 0xa3, 0xbd, 0x42, 0x57,
		0x21, 0xd2, 0x48, 0x30, 0x00, 0x6c, 0xa4, 0x1f, 0x05, 0x51, 0x3c, 0x80,
		0x08, 0xe7, 0xd8, 0xfa, 0x55, 0xb2, 0x79, 0xec, 0x5b, 0x09, 0x4b, 0xdb,
		0x08, 0xff, 0x05, 0x33, 0xbf, 0x30, 0xb6, 0x7e, 0x8d, 0x69, 0x1e, 0x74,
		0x8b, 0xa2, 0x03, 0xee, 0x33, 0x7f, 0x86, 0x43, 0xf7, 0x1f, 0x20, 0x33,
		0x3b, 0xdf, 0xc9, 0xef, 0x52, 0x97, 0x09, 0x64, 0xc6, 0xe6, 0x68, 0xfd,
		0xcc, 0xec, 0x3e, 0x75, 0xde, 0x63, 0x9f, 0x98, 0xf0, 0x82, 0xd0, 0x7a,
		0x90, 0x64, 0x58, 0x18, 0x8b, 0x6f, 0x49, 0xcd, 0x4c, 0xbe, 0x9f, 0xe2,
		0x6b, 0x6e, 0x4b, 0xa9, 0x13, 0x58, 0x7c, 0x1a, 0x0c, 0x0d, 0xcf, 0xf3,
		0x2e, 0x7b, 0xb2, 0x14, 0x46, 0x93, 0x5f, 0xf0, 0x5a, 0xaa, 0x7d, 0x02,
		0x4f, 0x26, 0x33, 0x64, 0x3c, 0x60, 0x7f, 0xa2, 0xda, 0x22, 0x49, 0xc1,
		0xe1, 0x6f, 0xdc, 0x20, 0xf3, 0x60, 0x32, 0x78, 0xf0, 0x9b, 0x95, 0x5c,
		0x79, 0xe0, 0xb8, 0x76, 0xbe, 0x43, 0x2b, 0x8b, 0x19, 0x94, 0x93, 0xdf,
		0x31, 0x81, 0xe8, 0xbe, 0xd9, 0x8d, 0x66, 0x25, 0x35, 0xfa, 0x15, 0xca,
		0xb2, 0xa2, 0x04, 0xe2, 0xc5, 0xc9, 0x21, 0x8c, 0x32, 0x36, 0x81, 0xbb,
		0xe5, 0x72, 0x39, 0x9a, 0x72, 0xe9, 0x1a, 0xc5, 0xf7, 0x09, 0x48, 0xdd,
		0xa5, 0x65, 0xca, 0x88, 0xe7, 0xd1, 0x59, 0x4b, 0xed, 0xbf, 0xc8, 0x9c,
		0xaa, 0x04, 0xa2, 0xc5, 0xe2, 0xa7, 0xf9, 0x67, 0x6f, 0x88, 0x8c, 0xf6,
		0x40, 0xea, 0x66, 0x43, 0x1e, 0x38, 0x54, 0x28, 0xc8, 0x03, 0xc2, 0x1d,
		0x71, 0x8b, 0x7c, 0xaa, 0xc8, 0x4c, 0x8c, 0xd4, 0x15, 0x5a, 0x49, 0xe7,
		0xfa, 0xaf, 0x8c, 0x83, 0xc8, 0x99, 0x75, 0xe0, 0x94, 0x75, 0x79, 0xd6,
		0x99, 0xb6, 0x1b, 0x67, 0x75, 0xdd, 0xa2, 0x6d, 0xcb, 0xa5, 0x7c, 0xae,
		0x64, 0xa9, 0x13, 0xa8, 0x65, 0x9e, 0x2b, 0x9c, 0xe5, 0x57, 0x91, 0x07,
		0x55, 0xec, 0x41, 0xb5, 0xf4, 0xa0, 0xba, 0xf7, 0xa0, 0x7a, 0xb8, 0x2d,
		0x33, 0x0a, 0xa2, 0x59, 0x1e, 0xf1, 0x4c, 0xe1, 0x05, 0xb3, 0xef, 0x1a,
		0x2e, 0xe6, 0x9d, 0x1d, 0x1c, 0xc2, 0x28, 0xc5, 0x1b, 0x87, 0x09, 0x8c,
		0x4f, 0x33, 0xb0, 0xa0, 0xd8, 0x28, 0xd5, 0x97, 0xd5, 0x6f, 0xb7, 0x82,
		0xd4, 0xa7, 0x8f, 0xba, 0x2a, 0x36, 0x40, 0xc6, 0xc5, 0x73, 0x69, 0xcd,
		0x46, 0xe7, 0xfe, 0xd8, 0xbe, 0x7b, 0xc1, 0x17, 0x5c, 0xcc, 0x41, 0x2f,
		0x91, 0x6a, 0xbe, 0x1b, 0x5b, 0x17, 0x3f, 0x7c, 0x3c, 0x9b, 0x81, 0xd7,
		0x27, 0x74, 0x19, 0x37, 0x3b, 0x88, 0xef, 0xaf, 0xc6, 0xe5, 0xa5, 0x92,
		0x84, 0xd7, 0x72, 0x26, 0x21, 0xb0, 0xb1, 0xea, 0x7d, 0x7f, 0x36, 0xb3,
		0xb2, 0x5b, 0x49, 0x1f, 0x40, 0x1b, 0xdf, 0x62, 0x83, 0x9c, 0x40, 0xa0,
		0x26, 0xb4, 0xa1, 0x30, 0x5b, 0xb4, 0xb7, 0x25, 0x57, 0xd1, 0xa4, 0xba,
		0x1b, 0xe8, 0x97, 0xa1, 0x0b, 0xda, 0xd8, 0x9a, 0xab, 0x1b, 0xb3, 0xbe,
		0x7c, 0xb8, 0xf1, 0x39, 0xb0, 0x80, 0xa8, 0xfd, 0x80, 0x57, 0x4f, 0xde,
		0x15, 0x6d, 0xfc, 0x46, 0xda, 0x1b, 0xa5, 0xb9, 0x13, 0x8b, 0x62, 0x51,
		0x14, 0xff, 0x5f, 0xdb, 0x0b, 0x0d, 0xc4, 0x6d, 0x89, 0x34, 0xe7, 0xef,
		0x56, 0x57, 0x02, 0x92, 0xb8, 0x92, 0xe2, 0x92, 0xe7, 0xd7, 0x13, 0xc9,
		0x80, 0xe9, 0x2b, 0x2c, 0x28, 0x81, 0x49, 0xd3, 0xe9, 0x88, 0x04, 0x5d,
		0x2b, 0x94, 0x29, 0xcd, 0xe5, 0x5c, 0xfd, 0x72, 0x35, 0x07, 0x7e, 0x66,
		0x88, 0x4c, 0x9d, 0xc0, 0xc7, 0x2b, 0x8f, 0x1d, 0xce, 0xc2, 0xe2, 0x82,
		0xe1, 0xae, 0x83, 0x6f, 0x37, 0xf0, 0x0f, 0x6c, 0xbd, 0xf8, 0x4c, 0x5f,
		0xfb, 0x37, 0x0d, 0xa7, 0x0d, 0x9d, 0x86, 0xe3, 0x9d, 0x98, 0xb6, 0x2b,
		0x74, 0x58, 0xf7, 0xb9, 0xdc, 0x82, 0x50, 0xdc, 0xb9, 0x15, 0xbb, 0x3e,
		0x27, 0xc3, 0x9d, 0x30, 0x0f, 0xbb, 0xf4, 0xb5, 0x97, 0x6d, 0xb4, 0x4e,
		0xdb, 0x5d, 0x31, 0x44, 0x4c, 0xf5, 0x60, 0x67, 0xd7, 0x48, 0xfb, 0x1e,
		0xb8, 0x6d, 0xc9, 0x80, 0x2b, 0x5a, 0xb1, 0x3f, 0x9e, 0x3e, 0x3f, 0xb2,
		0x70, 0x0d, 0x5f, 0x30, 0xeb, 0x6e, 0xe5, 0x2a, 0x3a, 0xe1, 0x1d, 0x0e,
		0x20, 0x0b, 0xd0, 0x86, 0xe0, 0x3d, 0x7e, 0x83, 0xe0, 0x9f, 0xbe, 0x79,
		0x8c, 0x7d, 0x80, 0xe3, 0x71, 0x0a, 0x6a, 0x69, 0xe3, 0xf5, 0xa3, 0xd1,
		0x1a, 0x05, 0x61, 0x0e, 0x64, 0x20, 0x75, 0x0d, 0xd7, 0xa3, 0x88, 0xbe,
		0xe5, 0x6c, 0x7d, 0x38, 0x4c, 0x08, 0xc7, 0x63, 0x1a, 0xb6, 0x21, 0xeb,
		0x34, 0xac, 0xe2, 0x19, 0x1d, 0xea, 0xfc, 0x84, 0x9d, 0x86, 0xb9, 0xdc,
		0x0e, 0xe5, 0x39, 0x3d, 0x1e, 0x0e, 0xc1, 0x17, 0xcc, 0x7e, 0x37, 0xb6,
		0x7e, 0x34, 0x9a, 0x50, 0x93, 0xeb, 0x12, 0xd2, 0xb0, 0x2f, 0x66, 0x1a,
		0xf6, 0x3f, 0x3d, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0xf6, 0x28,
		0x40, 0x82, 0x08, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}