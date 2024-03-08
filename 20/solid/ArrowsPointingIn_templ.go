// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowsPointingIn(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.28033 2.21967C2.98744 1.92678 2.51256 1.92678 2.21967 2.21967C1.92678 2.51256 1.92678 2.98744 2.21967 3.28033L5.43934 6.5H2.75C2.33579 6.5 2 6.83579 2 7.25C2 7.66421 2.33579 8 2.75 8H7.25C7.66421 8 8 7.66421 8 7.25V2.75C8 2.33579 7.66421 2 7.25 2C6.83579 2 6.5 2.33579 6.5 2.75V5.43934L3.28033 2.21967Z\" fill=\"#0F172A\"></path> <path d=\"M13.5 2.75C13.5 2.33579 13.1642 2 12.75 2C12.3358 2 12 2.33579 12 2.75V7.25C12 7.66421 12.3358 8 12.75 8H17.25C17.6642 8 18 7.66421 18 7.25C18 6.83579 17.6642 6.5 17.25 6.5H14.5607L17.7803 3.28033C18.0732 2.98744 18.0732 2.51256 17.7803 2.21967C17.4874 1.92678 17.0126 1.92678 16.7197 2.21967L13.5 5.43934V2.75Z\" fill=\"#0F172A\"></path> <path d=\"M3.28033 17.7803L6.5 14.5607V17.25C6.5 17.6642 6.83579 18 7.25 18C7.66421 18 8 17.6642 8 17.25V12.75C8 12.3358 7.66421 12 7.25 12H2.75C2.33579 12 2 12.3358 2 12.75C2 13.1642 2.33579 13.5 2.75 13.5H5.43934L2.21967 16.7197C1.92678 17.0126 1.92678 17.4874 2.21967 17.7803C2.51256 18.0732 2.98744 18.0732 3.28033 17.7803Z\" fill=\"#0F172A\"></path> <path d=\"M13.5 14.5607L16.7197 17.7803C17.0126 18.0732 17.4874 18.0732 17.7803 17.7803C18.0732 17.4874 18.0732 17.0126 17.7803 16.7197L14.5607 13.5H17.25C17.6642 13.5 18 13.1642 18 12.75C18 12.3358 17.6642 12 17.25 12H12.75C12.3358 12 12 12.3358 12 12.75V17.25C12 17.6642 12.3358 18 12.75 18C13.1642 18 13.5 17.6642 13.5 17.25V14.5607Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
