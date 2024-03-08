// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChartBar() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M12 2C11.4477 2 11 2.44772 11 3V13C11 13.5523 11.4477 14 12 14H13C13.5523 14 14 13.5523 14 13V3C14 2.44772 13.5523 2 13 2H12Z\" fill=\"black\"></path> <path d=\"M6.5 6C6.5 5.44772 6.94772 5 7.5 5H8.5C9.05228 5 9.5 5.44772 9.5 6V13C9.5 13.5523 9.05228 14 8.5 14H7.5C6.94772 14 6.5 13.5523 6.5 13V6Z\" fill=\"black\"></path> <path d=\"M2 9C2 8.44772 2.44772 8 3 8H4C4.55228 8 5 8.44772 5 9V13C5 13.5523 4.55228 14 4 14H3C2.44772 14 2 13.5523 2 13V9Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
