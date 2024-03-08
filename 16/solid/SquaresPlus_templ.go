// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SquaresPlus() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M2 3.5C2 2.67157 2.67157 2 3.5 2H5.5C6.32843 2 7 2.67157 7 3.5V5.5C7 6.32843 6.32843 7 5.5 7H3.5C2.67157 7 2 6.32843 2 5.5V3.5Z\" fill=\"black\"></path> <path d=\"M2 10.5C2 9.67157 2.67157 9 3.5 9H5.5C6.32843 9 7 9.67157 7 10.5V12.5C7 13.3284 6.32843 14 5.5 14H3.5C2.67157 14 2 13.3284 2 12.5V10.5Z\" fill=\"black\"></path> <path d=\"M10.5 2C9.67157 2 9 2.67157 9 3.5V5.5C9 6.32843 9.67157 7 10.5 7H12.5C13.3284 7 14 6.32843 14 5.5V3.5C14 2.67157 13.3284 2 12.5 2H10.5Z\" fill=\"black\"></path> <path d=\"M11.5 9C11.9142 9 12.25 9.33579 12.25 9.75V10.75H13.25C13.6642 10.75 14 11.0858 14 11.5C14 11.9142 13.6642 12.25 13.25 12.25H12.25V13.25C12.25 13.6642 11.9142 14 11.5 14C11.0858 14 10.75 13.6642 10.75 13.25V12.25H9.75C9.33579 12.25 9 11.9142 9 11.5C9 11.0858 9.33579 10.75 9.75 10.75H10.75V9.75C10.75 9.33579 11.0858 9 11.5 9Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
