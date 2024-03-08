// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BookOpen(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M10.75 16.8195C11.9579 15.9871 13.4212 15.5 15 15.5C15.7103 15.5 16.3964 15.5985 17.0459 15.7822C17.272 15.8462 17.515 15.8005 17.7024 15.6587C17.8899 15.5169 18 15.2955 18 15.0605V4.06055C18 3.72495 17.7771 3.4302 17.4541 3.33886C16.6731 3.11796 15.8497 3 15 3C13.4636 3 12.016 3.38549 10.75 4.06487V16.8195Z\" fill=\"#0F172A\"></path> <path d=\"M9.25 4.06487C7.98396 3.38549 6.5364 3 5 3C4.15029 3 3.32689 3.11796 2.54588 3.33886C2.22295 3.4302 2 3.72495 2 4.06055V15.0605C2 15.2955 2.11014 15.5169 2.29756 15.6587C2.48497 15.8005 2.728 15.8462 2.95412 15.7822C3.60361 15.5985 4.28967 15.5 5 15.5C6.57884 15.5 8.04208 15.9871 9.25 16.8195V4.06487Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
