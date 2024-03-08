// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Lifebuoy() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M19.449 8.44818L16.3878 10.9992C16.5374 11.6574 16.5374 12.3426 16.3878 13.0008L19.449 15.5518C20.517 13.3118 20.517 10.6882 19.449 8.44818ZM15.5518 19.449L13.0008 16.3878C12.3426 16.5374 11.6574 16.5374 10.9992 16.3878L8.44818 19.449C10.6882 20.517 13.3118 20.517 15.5518 19.449ZM4.55102 15.5518L7.6122 13.0008C7.4626 12.3426 7.4626 11.6574 7.6122 10.9992L4.55102 8.44818C3.48299 10.6882 3.48299 13.3118 4.55102 15.5518ZM8.44818 4.55102L10.9992 7.6122C11.6574 7.4626 12.3426 7.4626 13.0008 7.6122L15.5518 4.55102C13.3118 3.48299 10.6882 3.48299 8.44818 4.55102ZM17.1055 3.6912C17.7424 4.08325 18.3435 4.55493 18.8943 5.10571C19.4451 5.65649 19.9167 6.25755 20.3088 6.89448C22.2304 10.0163 22.2304 13.9837 20.3088 17.1055C19.9167 17.7424 19.4451 18.3435 18.8943 18.8943C18.3435 19.4451 17.7424 19.9167 17.1055 20.3088C13.9837 22.2304 10.0163 22.2304 6.89448 20.3088C6.25755 19.9167 5.65649 19.4451 5.10571 18.8943C4.55493 18.3435 4.08325 17.7424 3.6912 17.1055C1.7696 13.9837 1.7696 10.0163 3.6912 6.89448C4.08325 6.25755 4.55493 5.65649 5.10571 5.10571C5.65649 4.55493 6.25755 4.08325 6.89448 3.6912C10.0163 1.7696 13.9837 1.7696 17.1055 3.6912ZM14.1213 9.87868C13.7958 9.55313 13.4158 9.31907 13.0115 9.17471C12.359 8.94176 11.641 8.94176 10.9886 9.17471C10.5842 9.31907 10.2042 9.55313 9.87868 9.87868C9.55313 10.2042 9.31907 10.5842 9.17471 10.9885C8.94176 11.641 8.94176 12.359 9.17471 13.0114C9.31907 13.4158 9.55313 13.7958 9.87868 14.1213C10.2042 14.4469 10.5842 14.6809 10.9886 14.8253C11.641 15.0582 12.359 15.0582 13.0115 14.8253C13.4158 14.6809 13.7958 14.4469 14.1213 14.1213C14.4469 13.7958 14.6809 13.4158 14.8253 13.0115C15.0582 12.359 15.0582 11.641 14.8253 10.9885C14.6809 10.5842 14.4469 10.2042 14.1213 9.87868Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}