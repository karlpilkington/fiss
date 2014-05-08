package main

import (
    "bytes"
    "compress/gzip"
    "fmt"
    "io"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

func directory_list_go_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x94, 0x52,
		0xcb, 0x8e, 0x9c, 0x30, 0x10, 0x3c, 0x87, 0xaf, 0x68, 0x39, 0x4a, 0x6e,
		0x18, 0xcf, 0x2a, 0x9b, 0x48, 0x1b, 0x96, 0x43, 0x5e, 0xd2, 0x1e, 0x26,
		0x8a, 0x34, 0xfc, 0x80, 0x77, 0xdc, 0x2c, 0x96, 0xc0, 0xac, 0x8c, 0xb3,
		0xd2, 0xc4, 0xe1, 0xdf, 0xd3, 0x0d, 0x1e, 0xe6, 0x11, 0xe5, 0x10, 0x2e,
		0x54, 0x97, 0xab, 0xab, 0x9b, 0xc2, 0x31, 0x1a, 0x6c, 0xac, 0x43, 0x10,
		0x75, 0x6d, 0x43, 0x87, 0x62, 0x9a, 0x32, 0x80, 0x18, 0xe5, 0x56, 0xef,
		0x5b, 0xe2, 0xa7, 0x09, 0xde, 0x3e, 0x85, 0x8f, 0xcc, 0xfc, 0xd0, 0xa1,
		0xa5, 0xd3, 0x18, 0xd1, 0x19, 0x7a, 0x67, 0x67, 0xad, 0x9f, 0x07, 0x17,
		0xd0, 0x05, 0x6e, 0x2e, 0xdb, 0x4d, 0xb5, 0x8a, 0xcb, 0x82, 0xaa, 0xac,
		0x34, 0xf6, 0xa5, 0x22, 0xb9, 0xfc, 0xa4, 0x47, 0x7c, 0x70, 0xcd, 0x20,
		0xb7, 0x83, 0xa9, 0x6d, 0x8f, 0xf2, 0xdb, 0xe0, 0x7b, 0x1d, 0x40, 0xdc,
		0x28, 0xf5, 0x3e, 0x57, 0x9b, 0x5c, 0xdd, 0xc0, 0xe6, 0xf6, 0x4e, 0xbd,
		0xbb, 0x53, 0xb7, 0x90, 0xab, 0x0f, 0x4a, 0xc1, 0x76, 0x57, 0x8b, 0x79,
		0xe8, 0x45, 0x33, 0xf2, 0xa0, 0x62, 0xb6, 0x5d, 0xcc, 0x01, 0xca, 0xa0,
		0x1f, 0x3b, 0x64, 0xc4, 0xb8, 0x45, 0x6d, 0x16, 0xcc, 0x95, 0xaf, 0xb2,
		0x57, 0xc4, 0x55, 0xdf, 0x75, 0x8f, 0x65, 0x41, 0x60, 0x29, 0xc9, 0xc8,
		0x36, 0x16, 0xcd, 0x89, 0x82, 0x7d, 0xa7, 0xc7, 0xf1, 0x5e, 0xb8, 0x9f,
		0x3d, 0x7a, 0xbb, 0x17, 0xd5, 0xce, 0xfe, 0xba, 0xea, 0x48, 0x65, 0xb2,
		0x2e, 0xd8, 0x3b, 0xa1, 0xd3, 0xcc, 0x32, 0x3c, 0x0e, 0xe6, 0x70, 0x14,
		0xc5, 0xe8, 0xb5, 0x7b, 0x42, 0x90, 0x5f, 0x5d, 0xf0, 0x16, 0xc7, 0x39,
		0xe1, 0xb4, 0x18, 0x04, 0x0e, 0xfd, 0x5e, 0xc4, 0xf8, 0xec, 0xad, 0x0b,
		0x0d, 0x88, 0x37, 0xaf, 0x5f, 0x04, 0xc8, 0xdd, 0x81, 0x64, 0x02, 0x8c,
		0x0e, 0x3a, 0x7f, 0xa6, 0x28, 0x59, 0xe1, 0xb1, 0x63, 0x08, 0x92, 0x4e,
		0xe6, 0x7d, 0x0c, 0x07, 0xcd, 0xdf, 0xc4, 0x41, 0x53, 0xb5, 0x72, 0xff,
		0x1b, 0x6f, 0xea, 0xe6, 0x27, 0x46, 0xdb, 0x80, 0x7c, 0x18, 0xbf, 0x58,
		0x4f, 0x7b, 0xb2, 0xe1, 0x5f, 0x99, 0x5c, 0xa8, 0xb1, 0x1b, 0xf1, 0x5f,
		0x42, 0xda, 0x84, 0xf3, 0x83, 0xdf, 0xd0, 0xf4, 0x61, 0x24, 0x74, 0x35,
		0x69, 0xb9, 0x47, 0x67, 0x4b, 0xaf, 0x82, 0xf3, 0x74, 0x57, 0x61, 0xa2,
		0x8e, 0xd1, 0x12, 0x5c, 0xfe, 0x78, 0xba, 0x08, 0x49, 0xf6, 0x27, 0x00,
		0x00, 0xff, 0xff, 0xdf, 0x1f, 0xd7, 0xf2, 0xd1, 0x02, 0x00, 0x00,
		},
		"directory-list.go.html",
	)
}

func error_go_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x8f,
		0x41, 0xcb, 0x82, 0x40, 0x10, 0x86, 0xef, 0xfe, 0x8a, 0x61, 0x3f, 0xbe,
		0xab, 0x12, 0x78, 0x0a, 0xf1, 0x12, 0x1d, 0xbb, 0x84, 0x7f, 0x40, 0xf0,
		0xb5, 0x84, 0x5a, 0x75, 0x9c, 0x8a, 0x98, 0xf6, 0xbf, 0x37, 0x6e, 0x04,
		0x79, 0xf0, 0xb2, 0xec, 0xfb, 0xf2, 0x3c, 0x33, 0x8c, 0x6a, 0x83, 0xb6,
		0xf3, 0x20, 0x57, 0x55, 0x9d, 0x5c, 0xe0, 0x42, 0x48, 0x88, 0xf6, 0xcc,
		0x3d, 0x27, 0xaa, 0xf0, 0x8d, 0xe5, 0xe4, 0x07, 0xda, 0xf5, 0x5e, 0xe0,
		0x65, 0xc6, 0x8a, 0xf3, 0xa6, 0x8c, 0x60, 0x91, 0xd9, 0xcf, 0x62, 0x5e,
		0x1e, 0x30, 0x4d, 0xf5, 0x09, 0x5b, 0x6b, 0x72, 0x6b, 0x06, 0x46, 0xa9,
		0x9a, 0x82, 0x39, 0x8d, 0x60, 0x08, 0x45, 0x36, 0x77, 0x91, 0x3d, 0xd6,
		0x8f, 0xcf, 0x9e, 0x25, 0x6d, 0xaf, 0xc8, 0xb3, 0xbd, 0x0a, 0xcd, 0x1e,
		0xbd, 0x68, 0xe0, 0xce, 0x4b, 0x4b, 0xee, 0xff, 0x8f, 0xee, 0x6e, 0x39,
		0x01, 0xe3, 0x0d, 0x93, 0xac, 0xfa, 0x8c, 0x71, 0xd5, 0xff, 0xde, 0xf6,
		0x0e, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x44, 0xe1, 0x7f, 0x00, 0x01, 0x00,
		0x00,
		},
		"error.go.html",
	)
}

func layout_go_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x84, 0x94,
		0x4d, 0x6f, 0xd4, 0x3c, 0x10, 0xc7, 0xcf, 0x4f, 0x3f, 0xc5, 0x3c, 0xa1,
		0x28, 0xbb, 0xa2, 0x79, 0x59, 0x21, 0x90, 0x48, 0x93, 0x72, 0x28, 0x1c,
		0x11, 0x1c, 0xca, 0x01, 0x21, 0x0e, 0x5e, 0xdb, 0x89, 0xdd, 0x3a, 0x76,
		0xb0, 0x27, 0x6c, 0x57, 0x55, 0xbf, 0x3b, 0xf6, 0x7a, 0x53, 0x92, 0x15,
		0xa8, 0xf1, 0x21, 0x13, 0x7b, 0xe6, 0xe7, 0xf1, 0xfc, 0x27, 0xae, 0xff,
		0xff, 0xf0, 0xf9, 0xfa, 0xe6, 0xdb, 0x97, 0x8f, 0x20, 0xb0, 0x57, 0x57,
		0x67, 0x75, 0x7c, 0x01, 0xd4, 0x82, 0x13, 0x16, 0x0c, 0x6f, 0xa2, 0x44,
		0xc5, 0xa3, 0xfd, 0xf0, 0x80, 0xbc, 0x1f, 0x14, 0x41, 0x0e, 0xc9, 0xcd,
		0x4d, 0x98, 0x4f, 0x20, 0x7f, 0x7c, 0x8c, 0x7e, 0xc5, 0xcc, 0xb1, 0x56,
		0x52, 0xdf, 0x81, 0xb0, 0xbc, 0x6d, 0xd2, 0xa2, 0x68, 0x8d, 0x46, 0x97,
		0x77, 0xc6, 0x74, 0x8a, 0x93, 0x41, 0xba, 0x9c, 0x9a, 0xbe, 0xa0, 0xce,
		0xbd, 0x6f, 0x49, 0x2f, 0xd5, 0xbe, 0xf9, 0xba, 0x1d, 0x35, 0x8e, 0xaf,
		0x3e, 0x19, 0x6d, 0x52, 0xb0, 0x5c, 0x35, 0xa9, 0xc3, 0xbd, 0xe2, 0x4e,
		0x70, 0x8e, 0x29, 0xe0, 0x7e, 0xe0, 0x4d, 0x8a, 0xfc, 0x1e, 0x43, 0x4c,
		0x7a, 0xdc, 0xc0, 0x51, 0x2b, 0x07, 0x04, 0x67, 0x69, 0x93, 0x14, 0x05,
		0xb9, 0x25, 0xf7, 0xa7, 0x1b, 0x84, 0xb9, 0x42, 0xc9, 0xad, 0x2b, 0x6e,
		0x7f, 0x8e, 0xdc, 0xee, 0x8b, 0x4d, 0xbe, 0xd9, 0xe4, 0xe5, 0xf1, 0x2b,
		0xef, 0xa5, 0xce, 0x6f, 0x5d, 0x72, 0x55, 0x17, 0x11, 0xb5, 0xe0, 0xc6,
		0x8f, 0xf3, 0x55, 0x3b, 0x6a, 0x8a, 0xd2, 0xe8, 0xd5, 0xfa, 0xe1, 0xec,
		0xbf, 0xf3, 0x55, 0x82, 0xf6, 0x3b, 0x23, 0x48, 0xb2, 0x81, 0xa0, 0xf8,
		0x91, 0xac, 0x73, 0xaa, 0x24, 0xbd, 0x9b, 0x79, 0x81, 0x77, 0x0b, 0x91,
		0x3b, 0xa9, 0x99, 0xd9, 0xe5, 0xca, 0x50, 0x12, 0x16, 0xf2, 0x50, 0x09,
		0x68, 0x3c, 0x10, 0x85, 0x74, 0xeb, 0x3c, 0x30, 0x56, 0x49, 0x80, 0x24,
		0xeb, 0xcb, 0x10, 0xf1, 0xe8, 0x5f, 0x70, 0x78, 0x26, 0xeb, 0x34, 0xad,
		0x50, 0x90, 0x68, 0x6f, 0x0d, 0xdb, 0x47, 0x35, 0x8e, 0x21, 0x5b, 0x42,
		0xef, 0x3a, 0x6b, 0x46, 0xcd, 0x32, 0x6a, 0x94, 0xb1, 0x15, 0xbc, 0xd8,
		0x94, 0x61, 0x4c, 0xcc, 0x69, 0x96, 0xb7, 0x61, 0x4c, 0xb3, 0x41, 0x95,
		0x2c, 0x2a, 0x50, 0x41, 0x1a, 0x35, 0x80, 0x83, 0x06, 0x17, 0x90, 0xa6,
		0x93, 0xd7, 0x40, 0x18, 0x93, 0xba, 0xcb, 0x14, 0x6f, 0xb1, 0x82, 0x4d,
		0x39, 0xdc, 0xc7, 0x95, 0x28, 0x3a, 0x92, 0xad, 0xe2, 0x80, 0x16, 0xae,
		0xa0, 0xd2, 0x06, 0x57, 0x55, 0x2b, 0xad, 0xc3, 0x8c, 0x0a, 0xa9, 0x98,
		0x2f, 0x06, 0xec, 0x84, 0x44, 0x9e, 0xb9, 0x81, 0x50, 0x5e, 0x81, 0x36,
		0x3b, 0x4b, 0x86, 0x4b, 0x5f, 0x1c, 0x86, 0xc2, 0xb3, 0x3c, 0x6a, 0x46,
		0x59, 0x1c, 0xa9, 0x27, 0xb6, 0x93, 0x3a, 0x43, 0x33, 0x78, 0xbf, 0x37,
		0xd3, 0x9e, 0x30, 0x85, 0xbe, 0x2b, 0x5f, 0x2e, 0xb2, 0x10, 0x8b, 0xe0,
		0xd0, 0x2a, 0x19, 0x51, 0xb2, 0xd3, 0x55, 0xc8, 0xfa, 0xe4, 0x24, 0x15,
		0x78, 0xde, 0x01, 0xfa, 0x64, 0x4c, 0xac, 0x08, 0x63, 0x17, 0xa7, 0xc0,
		0x65, 0x0d, 0x66, 0xe9, 0x4c, 0x0b, 0x56, 0x76, 0x62, 0xb9, 0x32, 0xcb,
		0x01, 0xfe, 0x24, 0x71, 0x4c, 0x97, 0x2d, 0xe8, 0x74, 0xb4, 0x2e, 0xc8,
		0x33, 0x18, 0xa9, 0x91, 0xdb, 0xb9, 0x67, 0xae, 0xc7, 0x9e, 0x5b, 0x49,
		0xff, 0x75, 0x3c, 0x38, 0x6c, 0xfc, 0x54, 0x9c, 0x79, 0xb1, 0x07, 0xcb,
		0x97, 0xc7, 0xb2, 0x95, 0x46, 0x11, 0x95, 0x59, 0x19, 0xc6, 0xd6, 0xcf,
		0xb5, 0xd0, 0xeb, 0x32, 0x8c, 0x25, 0x23, 0xb4, 0x5e, 0x20, 0x09, 0xf3,
		0x8b, 0xdb, 0xe7, 0x00, 0xe5, 0xe1, 0x59, 0x02, 0xc4, 0x73, 0x41, 0x6f,
		0xcb, 0x30, 0xe6, 0x35, 0xf0, 0x3f, 0xc2, 0xd4, 0xfb, 0x75, 0x31, 0x5d,
		0x4b, 0x75, 0xc8, 0xe4, 0x2f, 0x77, 0xd2, 0xb5, 0xef, 0x69, 0xae, 0x71,
		0xba, 0x95, 0xea, 0x22, 0xfa, 0xf9, 0xc0, 0x70, 0xb1, 0xfd, 0x0e, 0x00,
		0x00, 0xff, 0xff, 0x02, 0x7f, 0xb8, 0x1a, 0xef, 0x04, 0x00, 0x00,
		},
		"layout.go.html",
	)
}

func root_list_go_html() ([]byte, error) {
	return bindata_read([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x5c, 0x90,
		0xcd, 0x6a, 0xc3, 0x30, 0x10, 0x84, 0xcf, 0xf5, 0x53, 0x2c, 0x2a, 0xed,
		0xcd, 0xb6, 0x1c, 0x9a, 0x16, 0x52, 0xd5, 0x97, 0x42, 0xa0, 0x50, 0xf7,
		0xd0, 0xf8, 0x05, 0x94, 0x6a, 0x1d, 0x0b, 0x6c, 0x29, 0xd8, 0x4b, 0xc0,
		0x08, 0xbf, 0x7b, 0x25, 0xff, 0x34, 0x25, 0xb7, 0xd1, 0xc7, 0xac, 0x76,
		0x66, 0x9d, 0x53, 0x58, 0x69, 0x83, 0xc0, 0xca, 0x52, 0x53, 0x83, 0x6c,
		0x1c, 0x23, 0x00, 0xe7, 0x92, 0x42, 0xfe, 0xd4, 0x9e, 0x8f, 0x23, 0x3c,
		0x9e, 0xe8, 0x15, 0xbe, 0xad, 0x25, 0xf8, 0xd4, 0x3d, 0x45, 0xce, 0xa1,
		0x51, 0xde, 0x15, 0xfd, 0x1b, 0x7d, 0xb7, 0x86, 0xd0, 0x50, 0x18, 0x16,
		0x75, 0x96, 0xef, 0x75, 0x83, 0xfd, 0xd0, 0x13, 0xb6, 0xd3, 0x5c, 0x2f,
		0x52, 0x0f, 0x23, 0xa1, 0xf4, 0x25, 0xf7, 0x9f, 0x0b, 0x92, 0xc7, 0x06,
		0x83, 0x0a, 0xba, 0x46, 0xa9, 0x66, 0x1d, 0x5e, 0x5d, 0x1e, 0xdd, 0x79,
		0x96, 0x7f, 0xc9, 0x16, 0x45, 0xea, 0xc5, 0xfc, 0x2c, 0xac, 0xd2, 0x95,
		0x46, 0x35, 0xa3, 0xc5, 0x9c, 0x06, 0xf7, 0xa2, 0xae, 0xbf, 0x08, 0x3a,
		0x5a, 0x35, 0xac, 0x26, 0xe7, 0x3a, 0x69, 0x4e, 0x08, 0x49, 0xc8, 0xf1,
		0x61, 0x2a, 0xdb, 0x4f, 0xfd, 0x96, 0x65, 0x40, 0xa1, 0xf2, 0x1b, 0x73,
		0xee, 0xdc, 0x69, 0x43, 0x15, 0xb0, 0x87, 0xfb, 0x0b, 0x83, 0xe4, 0x30,
		0x78, 0x1b, 0x03, 0x25, 0x49, 0xc6, 0x67, 0x49, 0x75, 0x70, 0x24, 0x21,
		0x92, 0xa7, 0x53, 0x22, 0x95, 0xff, 0x01, 0xbf, 0x5c, 0x5d, 0x99, 0x0f,
		0x5a, 0xea, 0x16, 0x93, 0xbd, 0xed, 0x5a, 0x49, 0xc0, 0x36, 0x9c, 0x3f,
		0xc7, 0x3c, 0x8b, 0xf9, 0x06, 0xb2, 0xed, 0x8e, 0x3f, 0xed, 0xf8, 0x16,
		0x62, 0xfe, 0xc2, 0x39, 0x14, 0x87, 0x92, 0xad, 0xd3, 0xb7, 0x85, 0x42,
		0xee, 0xf9, 0xc8, 0x0b, 0x5f, 0x2b, 0x79, 0x39, 0xdf, 0x4e, 0xa4, 0xd3,
		0x31, 0x17, 0xdb, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xd6, 0xa7,
		0x5b, 0xc5, 0x01, 0x00, 0x00,
		},
		"root-list.go.html",
	)
}


// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	if f, ok := _bindata[name]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string] func() ([]byte, error) {
	"directory-list.go.html": directory_list_go_html,
	"error.go.html": error_go_html,
	"layout.go.html": layout_go_html,
	"root-list.go.html": root_list_go_html,

}
