// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func EyeSlash() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.28033 2.21967C2.98744 1.92678 2.51256 1.92678 2.21967 2.21967C1.92678 2.51256 1.92678 2.98744 2.21967 3.28033L12.7197 13.7803C13.0126 14.0732 13.4874 14.0732 13.7803 13.7803C14.0732 13.4874 14.0732 13.0126 13.7803 12.7197L12.4577 11.397C13.438 10.5863 14.1937 9.51366 14.6176 8.2863C14.681 8.10274 14.6811 7.90313 14.6179 7.71951C13.672 4.97316 11.0653 3 7.99777 3C6.85414 3 5.77457 3.27425 4.82123 3.76057L3.28033 2.21967ZM6.47602 5.41536L7.61147 6.55081C7.73539 6.51767 7.86563 6.5 8 6.5C8.82843 6.5 9.5 7.17157 9.5 8C9.5 8.13437 9.48233 8.26461 9.44919 8.38853L10.5846 9.52398C10.8486 9.07734 11 8.55636 11 8C11 6.34315 9.65685 5 8 5C7.44364 5 6.92266 5.15145 6.47602 5.41536Z\" fill=\"black\"></path> <path d=\"M7.81206 10.9942L9.62754 12.8097C9.10513 12.9341 8.56002 13 7.99952 13C4.93197 13 2.32527 11.0268 1.3794 8.28049C1.31616 8.09687 1.31625 7.89727 1.37965 7.71371C1.63675 6.96935 2.01588 6.28191 2.49314 5.67529L5.00579 8.18794C5.09895 9.69509 6.30491 10.901 7.81206 10.9942Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
