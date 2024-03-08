// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Tv() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M19.5 6H4.5V15H19.5V6Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.375 3C2.33947 3 1.5 3.83947 1.5 4.875V16.125C1.5 17.1605 2.33947 18 3.375 18H9.75V19.5H6C5.58579 19.5 5.25 19.8358 5.25 20.25C5.25 20.6642 5.58579 21 6 21H18C18.4142 21 18.75 20.6642 18.75 20.25C18.75 19.8358 18.4142 19.5 18 19.5H14.25V18H20.625C21.6605 18 22.5 17.1605 22.5 16.125V4.875C22.5 3.83947 21.6605 3 20.625 3H3.375ZM3.375 16.5H20.625C20.8321 16.5 21 16.3321 21 16.125V4.875C21 4.66789 20.8321 4.5 20.625 4.5H3.375C3.16789 4.5 3 4.66789 3 4.875V16.125C3 16.3321 3.16789 16.5 3.375 16.5Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
