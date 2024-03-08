// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PhoneXMark() templ.Component {
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
		templ_7745c5c3_Err = PhoneXMarkWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PhoneXMarkWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.85462 7.28621L4.92241 6.75232C5.31094 6.55806 5.52962 6.13518 5.46357 5.70584L5.02389 2.84794C4.94884 2.36011 4.52909 2 4.03552 2H3C2.44772 2 2 2.44772 2 3V5C2 5.70879 2.08208 6.39936 2.23758 7.06237C3.01505 10.3773 5.62265 12.985 8.93763 13.7624C9.60064 13.9179 10.2912 14 11 14H13C13.5523 14 14 13.5523 14 13V11.9645C14 11.4709 13.6399 11.0512 13.1521 10.9761L10.2942 10.5364C9.86482 10.4704 9.44195 10.6891 9.24768 11.0776L8.71379 12.1454C6.4109 11.4092 4.59083 9.58911 3.85462 7.28621Z\" fill=\"black\"></path> <path d=\"M13.7803 2.21967C14.0732 2.51256 14.0732 2.98744 13.7803 3.28033L12.5607 4.5L13.7803 5.71967C14.0732 6.01256 14.0732 6.48744 13.7803 6.78033C13.4874 7.07322 13.0126 7.07322 12.7197 6.78033L11.5 5.56066L10.2803 6.78033C9.98744 7.07322 9.51256 7.07322 9.21967 6.78033C8.92678 6.48744 8.92678 6.01256 9.21967 5.71967L10.4393 4.5L9.21967 3.28033C8.92678 2.98744 8.92678 2.51256 9.21967 2.21967C9.51256 1.92678 9.98744 1.92678 10.2803 2.21967L11.5 3.43934L12.7197 2.21967C13.0126 1.92678 13.4874 1.92678 13.7803 2.21967Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
