// generated by stringer -type=Format -output=z_gen_format_string.go; DO NOT EDIT

package bpg

import "fmt"

const _Format_name = "FormatGRAYFormat420Format422Format444Format420VideoFormat422VideoFormatMax"

var _Format_index = [...]uint8{0, 10, 19, 28, 37, 51, 65, 74}

func (i Format) String() string {
	if i+1 >= Format(len(_Format_index)) {
		return fmt.Sprintf("Format(%d)", i)
	}
	return _Format_name[_Format_index[i]:_Format_index[i+1]]
}
