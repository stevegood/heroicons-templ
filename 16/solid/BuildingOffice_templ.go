// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BuildingOffice() templ.Component {
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
		templ_7745c5c3_Err = BuildingOfficeWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func BuildingOfficeWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.75 2C3.33579 2 3 2.33579 3 2.75C3 3.16421 3.33579 3.5 3.75 3.5H4V12.5H3.75C3.33579 12.5 3 12.8358 3 13.25C3 13.6642 3.33579 14 3.75 14H6C6.27614 14 6.5 13.7761 6.5 13.5V10.5C6.5 10.2239 6.72386 10 7 10H9C9.27614 10 9.5 10.2239 9.5 10.5V13.5C9.5 13.7761 9.72386 14 10 14H12.25C12.6642 14 13 13.6642 13 13.25C13 12.8358 12.6642 12.5 12.25 12.5H12V3.5H12.25C12.6642 3.5 13 3.16421 13 2.75C13 2.33579 12.6642 2 12.25 2H3.75ZM6.5 4C6.22386 4 6 4.22386 6 4.5V5C6 5.27614 6.22386 5.5 6.5 5.5H7C7.27614 5.5 7.5 5.27614 7.5 5V4.5C7.5 4.22386 7.27614 4 7 4H6.5ZM6 7C6 6.72386 6.22386 6.5 6.5 6.5H7C7.27614 6.5 7.5 6.72386 7.5 7V7.5C7.5 7.77614 7.27614 8 7 8H6.5C6.22386 8 6 7.77614 6 7.5V7ZM9 4C8.72386 4 8.5 4.22386 8.5 4.5V5C8.5 5.27614 8.72386 5.5 9 5.5H9.5C9.77614 5.5 10 5.27614 10 5V4.5C10 4.22386 9.77614 4 9.5 4H9ZM8.5 7C8.5 6.72386 8.72386 6.5 9 6.5H9.5C9.77614 6.5 10 6.72386 10 7V7.5C10 7.77614 9.77614 8 9.5 8H9C8.72386 8 8.5 7.77614 8.5 7.5V7Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
