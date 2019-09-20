// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package lxd

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

func easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(in *jlexer.Lexer, out *Metadata) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Architecture":
			out.Architecture = string(in.String())
		case "Config":
			(out.Config).UnmarshalEasyJSON(in)
		case "CreatedAt":
			out.CreatedAt = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Devices":
			(out.Devices).UnmarshalEasyJSON(in)
		case "Ephemeral":
			out.Ephemeral = string(in.String())
		case "Profiles":
			if in.IsNull() {
				in.Skip()
				out.Profiles = nil
			} else {
				in.Delim('[')
				if out.Profiles == nil {
					if !in.IsDelim(']') {
						out.Profiles = make([]string, 0, 4)
					} else {
						out.Profiles = []string{}
					}
				} else {
					out.Profiles = (out.Profiles)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Profiles = append(out.Profiles, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Restore":
			out.Restore = string(in.String())
		case "Status":
			out.Status = string(in.String())
		case "Stateful":
			out.Stateful = string(in.String())
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
func easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(out *jwriter.Writer, in Metadata) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Architecture != "" {
		const prefix string = ",\"Architecture\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Architecture))
	}
	if len(in.Config) != 0 {
		const prefix string = ",\"Config\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Config).MarshalEasyJSON(out)
	}
	if in.CreatedAt != "" {
		const prefix string = ",\"CreatedAt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CreatedAt))
	}
	if in.Description != "" {
		const prefix string = ",\"Description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	if len(in.Devices) != 0 {
		const prefix string = ",\"Devices\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(in.Devices).MarshalEasyJSON(out)
	}
	if in.Ephemeral != "" {
		const prefix string = ",\"Ephemeral\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Ephemeral))
	}
	if len(in.Profiles) != 0 {
		const prefix string = ",\"Profiles\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Profiles {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.Restore != "" {
		const prefix string = ",\"Restore\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Restore))
	}
	if in.Status != "" {
		const prefix string = ",\"Status\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Status))
	}
	if in.Stateful != "" {
		const prefix string = ",\"Stateful\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Stateful))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Metadata) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metadata) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBa0ee0e3EncodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metadata) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metadata) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBa0ee0e3DecodeGithubComSkydiveProjectSkydiveTopologyProbesLxd(l, v)
}