// generated by stringer -type=ExtensionTag -output=z_gen_extension_tag_string.go; DO NOT EDIT

package bpg

import "fmt"

const _ExtensionTag_name = "ExtensionTagEXIFExtensionTagICCPExtensionTagXMPExtensionTagTHUMBNAIL"

var _ExtensionTag_index = [...]uint8{16, 32, 47, 68}

func (i ExtensionTag) String() string {
	if i < 0 || i >= ExtensionTag(len(_ExtensionTag_index)) {
		return fmt.Sprintf("ExtensionTag(%d)", i)
	}
	hi := _ExtensionTag_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _ExtensionTag_index[i-1]
	}
	return _ExtensionTag_name[lo:hi]
}