// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BoltSlash(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.21967 2.21967C2.51256 1.92678 2.98744 1.92678 3.28033 2.21967L17.7803 16.7197C18.0732 17.0126 18.0732 17.4874 17.7803 17.7803C17.4874 18.0732 17.0126 18.0732 16.7197 17.7803L2.21967 3.28033C1.92678 2.98744 1.92678 2.51256 2.21967 2.21967Z\" fill=\"#0F172A\"></path> <path d=\"M4.73016 7.91214L2.19108 10.7499C1.99385 10.9704 1.9446 11.2861 2.06533 11.5562C2.18607 11.8262 2.45423 12 2.75001 12H8.81805L4.73016 7.91214Z\" fill=\"#0F172A\"></path> <path d=\"M9.23329 12.4153L8.01666 18.0929C7.9454 18.4255 8.10685 18.7644 8.41002 18.9185C8.71318 19.0727 9.08215 19.0036 9.30894 18.7501L12.2647 15.4467L9.23329 12.4153Z\" fill=\"#0F172A\"></path> <path d=\"M15.2699 12.0879L17.8089 9.25013C18.0062 9.0297 18.0554 8.71393 17.9347 8.4439C17.814 8.17388 17.5458 8.00003 17.25 8.00003H11.182L15.2699 12.0879Z\" fill=\"#0F172A\"></path> <path d=\"M10.7667 7.58476L11.9834 1.90718C12.0546 1.57461 11.8932 1.23571 11.59 1.08152C11.2868 0.927338 10.9179 0.996463 10.6911 1.24994L7.73537 4.55339L10.7667 7.58476Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
