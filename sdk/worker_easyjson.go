// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package sdk

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	time "time"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson82a45abeDecodeGithubComOvhCdsSdk(in *jlexer.Lexer, out *Model) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "image":
			out.Image = string(in.String())
		case "capabilities":
			if in.IsNull() {
				in.Skip()
				out.Capabilities = nil
			} else {
				in.Delim('[')
				if out.Capabilities == nil {
					if !in.IsDelim(']') {
						out.Capabilities = make([]Requirement, 0, 1)
					} else {
						out.Capabilities = []Requirement{}
					}
				} else {
					out.Capabilities = (out.Capabilities)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Requirement
					if data := in.Raw(); in.Ok() {
						in.AddError((v1).UnmarshalJSON(data))
					}
					out.Capabilities = append(out.Capabilities, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "communication":
			out.Communication = string(in.String())
		case "template":
			if in.IsNull() {
				in.Skip()
				out.Template = nil
			} else {
				if out.Template == nil {
					out.Template = new(map[string]string)
				}
				if in.IsNull() {
					in.Skip()
				} else {
					in.Delim('{')
					if !in.IsDelim('}') {
						*out.Template = make(map[string]string)
					} else {
						*out.Template = nil
					}
					for !in.IsDelim('}') {
						key := string(in.String())
						in.WantColon()
						var v2 string
						v2 = string(in.String())
						(*out.Template)[key] = v2
						in.WantComma()
					}
					in.Delim('}')
				}
			}
		case "run_script":
			out.RunScript = string(in.String())
		case "disabled":
			out.Disabled = bool(in.Bool())
		case "restricted":
			out.Restricted = bool(in.Bool())
		case "need_registration":
			out.NeedRegistration = bool(in.Bool())
		case "last_registration":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.LastRegistration).UnmarshalJSON(data))
			}
		case "user_last_modified":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.UserLastModified).UnmarshalJSON(data))
			}
		case "created_by":
			easyjson82a45abeDecodeGithubComOvhCdsSdk1(in, &out.CreatedBy)
		case "provision":
			out.Provision = int64(in.Int64())
		case "group_id":
			out.GroupID = int64(in.Int64())
		case "group":
			easyjson82a45abeDecodeGithubComOvhCdsSdk2(in, &out.Group)
		case "nb_spawn_err":
			out.NbSpawnErr = int64(in.Int64())
		case "last_spawn_err":
			out.LastSpawnErr = string(in.String())
		case "date_last_spawn_err":
			if in.IsNull() {
				in.Skip()
				out.DateLastSpawnErr = nil
			} else {
				if out.DateLastSpawnErr == nil {
					out.DateLastSpawnErr = new(time.Time)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.DateLastSpawnErr).UnmarshalJSON(data))
				}
			}
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
func easyjson82a45abeEncodeGithubComOvhCdsSdk(out *jwriter.Writer, in Model) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"image\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Image))
	}
	{
		const prefix string = ",\"capabilities\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Capabilities == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Capabilities {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Raw((v4).MarshalJSON())
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"communication\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Communication))
	}
	{
		const prefix string = ",\"template\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Template == nil {
			out.RawString("null")
		} else {
			if *in.Template == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
				out.RawString(`null`)
			} else {
				out.RawByte('{')
				v5First := true
				for v5Name, v5Value := range *in.Template {
					if v5First {
						v5First = false
					} else {
						out.RawByte(',')
					}
					out.String(string(v5Name))
					out.RawByte(':')
					out.String(string(v5Value))
				}
				out.RawByte('}')
			}
		}
	}
	{
		const prefix string = ",\"run_script\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RunScript))
	}
	{
		const prefix string = ",\"disabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Disabled))
	}
	{
		const prefix string = ",\"restricted\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Restricted))
	}
	{
		const prefix string = ",\"need_registration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.NeedRegistration))
	}
	{
		const prefix string = ",\"last_registration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.LastRegistration).MarshalJSON())
	}
	{
		const prefix string = ",\"user_last_modified\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.UserLastModified).MarshalJSON())
	}
	{
		const prefix string = ",\"created_by\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a45abeEncodeGithubComOvhCdsSdk1(out, in.CreatedBy)
	}
	{
		const prefix string = ",\"provision\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Provision))
	}
	{
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.GroupID))
	}
	{
		const prefix string = ",\"group\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson82a45abeEncodeGithubComOvhCdsSdk2(out, in.Group)
	}
	{
		const prefix string = ",\"nb_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.NbSpawnErr))
	}
	{
		const prefix string = ",\"last_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LastSpawnErr))
	}
	{
		const prefix string = ",\"date_last_spawn_err\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.DateLastSpawnErr == nil {
			out.RawString("null")
		} else {
			out.Raw((*in.DateLastSpawnErr).MarshalJSON())
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Model) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson82a45abeEncodeGithubComOvhCdsSdk(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Model) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson82a45abeEncodeGithubComOvhCdsSdk(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Model) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson82a45abeDecodeGithubComOvhCdsSdk(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Model) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson82a45abeDecodeGithubComOvhCdsSdk(l, v)
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk2(in *jlexer.Lexer, out *Group) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "admins":
			if in.IsNull() {
				in.Skip()
				out.Admins = nil
			} else {
				in.Delim('[')
				if out.Admins == nil {
					if !in.IsDelim(']') {
						out.Admins = make([]User, 0, 1)
					} else {
						out.Admins = []User{}
					}
				} else {
					out.Admins = (out.Admins)[:0]
				}
				for !in.IsDelim(']') {
					var v6 User
					easyjson82a45abeDecodeGithubComOvhCdsSdk1(in, &v6)
					out.Admins = append(out.Admins, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]User, 0, 1)
					} else {
						out.Users = []User{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v7 User
					easyjson82a45abeDecodeGithubComOvhCdsSdk1(in, &v7)
					out.Users = append(out.Users, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "tokens":
			if in.IsNull() {
				in.Skip()
				out.Tokens = nil
			} else {
				in.Delim('[')
				if out.Tokens == nil {
					if !in.IsDelim(']') {
						out.Tokens = make([]Token, 0, 1)
					} else {
						out.Tokens = []Token{}
					}
				} else {
					out.Tokens = (out.Tokens)[:0]
				}
				for !in.IsDelim(']') {
					var v8 Token
					easyjson82a45abeDecodeGithubComOvhCdsSdk3(in, &v8)
					out.Tokens = append(out.Tokens, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson82a45abeEncodeGithubComOvhCdsSdk2(out *jwriter.Writer, in Group) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Admins) != 0 {
		const prefix string = ",\"admins\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v9, v10 := range in.Admins {
				if v9 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk1(out, v10)
			}
			out.RawByte(']')
		}
	}
	if len(in.Users) != 0 {
		const prefix string = ",\"users\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v11, v12 := range in.Users {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk1(out, v12)
			}
			out.RawByte(']')
		}
	}
	if len(in.Tokens) != 0 {
		const prefix string = ",\"tokens\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v13, v14 := range in.Tokens {
				if v13 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk3(out, v14)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk3(in *jlexer.Lexer, out *Token) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "group_id":
			out.GroupID = int64(in.Int64())
		case "group_name":
			out.GroupName = string(in.String())
		case "token":
			out.Token = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "creator":
			out.Creator = string(in.String())
		case "expiration":
			out.Expiration = Expiration(in.Int())
		case "created":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Created).UnmarshalJSON(data))
			}
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
func easyjson82a45abeEncodeGithubComOvhCdsSdk3(out *jwriter.Writer, in Token) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"group_id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.GroupID))
	}
	{
		const prefix string = ",\"group_name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.GroupName))
	}
	{
		const prefix string = ",\"token\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Token))
	}
	{
		const prefix string = ",\"description\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"creator\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Creator))
	}
	{
		const prefix string = ",\"expiration\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Expiration))
	}
	{
		const prefix string = ",\"created\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Created).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson82a45abeDecodeGithubComOvhCdsSdk1(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "username":
			out.Username = string(in.String())
		case "fullname":
			out.Fullname = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "admin":
			out.Admin = bool(in.Bool())
		case "groups":
			if in.IsNull() {
				in.Skip()
				out.Groups = nil
			} else {
				in.Delim('[')
				if out.Groups == nil {
					if !in.IsDelim(']') {
						out.Groups = make([]Group, 0, 1)
					} else {
						out.Groups = []Group{}
					}
				} else {
					out.Groups = (out.Groups)[:0]
				}
				for !in.IsDelim(']') {
					var v15 Group
					easyjson82a45abeDecodeGithubComOvhCdsSdk2(in, &v15)
					out.Groups = append(out.Groups, v15)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "origin":
			out.Origin = string(in.String())
		case "permissions":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Permissions).UnmarshalJSON(data))
			}
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
func easyjson82a45abeEncodeGithubComOvhCdsSdk1(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"fullname\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Fullname))
	}
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"admin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Admin))
	}
	if len(in.Groups) != 0 {
		const prefix string = ",\"groups\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v16, v17 := range in.Groups {
				if v16 > 0 {
					out.RawByte(',')
				}
				easyjson82a45abeEncodeGithubComOvhCdsSdk2(out, v17)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"origin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Origin))
	}
	if true {
		const prefix string = ",\"permissions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Permissions).MarshalJSON())
	}
	out.RawByte('}')
}
