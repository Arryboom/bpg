// generated by stringer -type=OutputFormat -output=z_gen_output_format_string.go; DO NOT EDIT

package bpg

import "fmt"

const _OutputFormat_name = "OutputFormatRGB24OutputFormatRGBA32OutputFormatRGB48OutputFormatRGBA64OutputFormatCMYK32OutputFormatCMYK64"

var _OutputFormat_index = [...]uint8{0, 17, 35, 52, 70, 88, 106}

func (i OutputFormat) String() string {
	if i+1 >= OutputFormat(len(_OutputFormat_index)) {
		return fmt.Sprintf("OutputFormat(%d)", i)
	}
	return _OutputFormat_name[_OutputFormat_index[i]:_OutputFormat_index[i+1]]
}
