// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Newspaper() templ.Component {
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
		templ_7745c5c3_Err = NewspaperWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func NewspaperWithAttrs(attrs templ.Attributes) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.125 3C3.08947 3 2.25 3.83947 2.25 4.875V18C2.25 19.6569 3.59315 21 5.25 21H20.25C18.5931 21 17.25 19.6569 17.25 18V4.875C17.25 3.83947 16.4105 3 15.375 3H4.125ZM12 9.75C11.5858 9.75 11.25 10.0858 11.25 10.5C11.25 10.9142 11.5858 11.25 12 11.25H13.5C13.9142 11.25 14.25 10.9142 14.25 10.5C14.25 10.0858 13.9142 9.75 13.5 9.75H12ZM11.25 7.5C11.25 7.08579 11.5858 6.75 12 6.75H13.5C13.9142 6.75 14.25 7.08579 14.25 7.5C14.25 7.91421 13.9142 8.25 13.5 8.25H12C11.5858 8.25 11.25 7.91421 11.25 7.5ZM6 12.75C5.58579 12.75 5.25 13.0858 5.25 13.5C5.25 13.9142 5.58579 14.25 6 14.25H13.5C13.9142 14.25 14.25 13.9142 14.25 13.5C14.25 13.0858 13.9142 12.75 13.5 12.75H6ZM5.25 16.5C5.25 16.0858 5.58579 15.75 6 15.75H13.5C13.9142 15.75 14.25 16.0858 14.25 16.5C14.25 16.9142 13.9142 17.25 13.5 17.25H6C5.58579 17.25 5.25 16.9142 5.25 16.5ZM6 6.75C5.58579 6.75 5.25 7.08579 5.25 7.5V10.5C5.25 10.9142 5.58579 11.25 6 11.25H9C9.41421 11.25 9.75 10.9142 9.75 10.5V7.5C9.75 7.08579 9.41421 6.75 9 6.75H6Z\" fill=\"#0F172A\"></path> <path d=\"M18.75 6.75H20.625C21.2463 6.75 21.75 7.25368 21.75 7.875V18C21.75 18.8284 21.0784 19.5 20.25 19.5C19.4216 19.5 18.75 18.8284 18.75 18V6.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
