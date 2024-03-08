// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PaintBrush(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M12.6133 1.25783C12.8654 1.08971 13.1617 1 13.4648 1C14.3127 1 15 1.68733 15 2.53518C15 2.83827 14.9103 3.13457 14.7422 3.38675L12.8381 6.24283C11.9595 7.56078 10.7169 8.57786 9.27774 9.18189C8.87694 8.03308 7.96692 7.12306 6.81812 6.72226C7.42214 5.28315 8.43922 4.04052 9.75718 3.16189L12.6133 1.25783Z\" fill=\"black\"></path> <path d=\"M5.49976 8C4.11905 8 2.99976 9.11929 2.99976 10.5C2.99976 10.7761 2.77591 11 2.49976 11C2.42716 11 2.36004 10.985 2.29976 10.9586C2.01467 10.8338 1.68215 10.8981 1.46405 11.12C1.24595 11.342 1.1876 11.6756 1.31739 11.9585C1.86958 13.1618 3.086 14 4.49976 14C6.43276 14 7.99976 12.433 7.99976 10.5C7.99976 9.11929 6.88048 8 5.49976 8Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
