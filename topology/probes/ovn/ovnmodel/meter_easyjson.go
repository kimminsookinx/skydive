// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package ovnmodel

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonA2a2364dDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(in *jlexer.Lexer, out *Meter) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "UUID":
			out.UUID = string(in.String())
		case "Bands":
			if in.IsNull() {
				in.Skip()
				out.Bands = nil
			} else {
				in.Delim('[')
				if out.Bands == nil {
					if !in.IsDelim(']') {
						out.Bands = make([]string, 0, 4)
					} else {
						out.Bands = []string{}
					}
				} else {
					out.Bands = (out.Bands)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Bands = append(out.Bands, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ExternalIDs":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.ExternalIDs = make(map[string]string)
				} else {
					out.ExternalIDs = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.ExternalIDs)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Fair":
			if in.IsNull() {
				in.Skip()
				out.Fair = nil
			} else {
				in.Delim('[')
				if out.Fair == nil {
					if !in.IsDelim(']') {
						out.Fair = make([]bool, 0, 64)
					} else {
						out.Fair = []bool{}
					}
				} else {
					out.Fair = (out.Fair)[:0]
				}
				for !in.IsDelim(']') {
					var v3 bool
					v3 = bool(in.Bool())
					out.Fair = append(out.Fair, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Name":
			out.Name = string(in.String())
		case "Unit":
			out.Unit = string(in.String())
		case "ExternalIDsMeta":
			(out.ExternalIDsMeta).UnmarshalEasyJSON(in)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA2a2364dEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(out *jwriter.Writer, in Meter) {
	out.RawByte('{')
	first := true
	_ = first
	if in.UUID != "" {
		const prefix string = ",\"UUID\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.UUID))
	}
	if len(in.Bands) != 0 {
		const prefix string = ",\"Bands\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.Bands {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	if len(in.ExternalIDs) != 0 {
		const prefix string = ",\"ExternalIDs\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v6First := true
			for v6Name, v6Value := range in.ExternalIDs {
				if v6First {
					v6First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v6Name))
				out.RawByte(':')
				out.String(string(v6Value))
			}
			out.RawByte('}')
		}
	}
	if len(in.Fair) != 0 {
		const prefix string = ",\"Fair\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Fair {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.Bool(bool(v8))
			}
			out.RawByte(']')
		}
	}
	if in.Name != "" {
		const prefix string = ",\"Name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Unit != "" {
		const prefix string = ",\"Unit\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Unit))
	}
	if len(in.ExternalIDsMeta) != 0 {
		const prefix string = ",\"ExternalIDsMeta\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.ExternalIDsMeta).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Meter) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA2a2364dEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Meter) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA2a2364dEncodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Meter) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA2a2364dDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Meter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA2a2364dDecodeGithubComSkydiveProjectSkydiveTopologyProbesOvnOvnmodel(l, v)
}
