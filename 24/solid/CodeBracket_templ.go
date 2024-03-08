// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CodeBracket(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M14.4473 3.02637C14.847 3.13536 15.0826 3.54766 14.9736 3.94728L10.4736 20.4473C10.3646 20.8469 9.95228 21.0825 9.55266 20.9735C9.15304 20.8645 8.91744 20.4522 9.02643 20.0526L13.5264 3.55261C13.6354 3.15299 14.0477 2.91738 14.4473 3.02637ZM16.7197 6.21961C17.0126 5.92672 17.4874 5.92672 17.7803 6.21961L23.0303 11.4696C23.3232 11.7625 23.3232 12.2374 23.0303 12.5303L17.7803 17.7803C17.4874 18.0732 17.0126 18.0732 16.7197 17.7803C16.4268 17.4874 16.4268 17.0125 16.7197 16.7196L21.4393 11.9999L16.7197 7.28027C16.4268 6.98738 16.4268 6.51251 16.7197 6.21961ZM7.28033 6.21961C7.57322 6.51251 7.57322 6.98738 7.28033 7.28027L2.56066 11.9999L7.28033 16.7196C7.57322 17.0125 7.57322 17.4874 7.28033 17.7803C6.98744 18.0732 6.51256 18.0732 6.21967 17.7803L0.96967 12.5303C0.676777 12.2374 0.676777 11.7625 0.96967 11.4696L6.21967 6.21961C6.51256 5.92672 6.98744 5.92672 7.28033 6.21961Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
