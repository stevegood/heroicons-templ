// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowPath() templ.Component {
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
		templ_7745c5c3_Err = ArrowPathWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ArrowPathWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M13.8356 2.47703C14.2498 2.47703 14.5856 2.81282 14.5856 3.22703V6.40901C14.5856 6.82322 14.2498 7.15901 13.8356 7.15901H10.6536C10.2394 7.15901 9.90362 6.82322 9.90362 6.40901C9.90362 5.9948 10.2394 5.65901 10.6536 5.65901H12.0249L11.1839 4.81802C9.42659 3.06066 6.57734 3.06066 4.81999 4.81802C4.53421 5.10379 4.29556 5.41752 4.10337 5.74996C3.89606 6.10855 3.43729 6.23119 3.07869 6.02388C2.7201 5.81656 2.59746 5.3578 2.80477 4.9992C3.06184 4.55455 3.38024 4.13644 3.75933 3.75736C6.10247 1.41421 9.90146 1.41421 12.2446 3.75736L13.0856 4.59835V3.22703C13.0856 2.81282 13.4214 2.47703 13.8356 2.47703ZM12.925 9.97662C13.2836 10.184 13.4062 10.6427 13.1989 11.0013C12.9418 11.4458 12.6235 11.8637 12.2446 12.2426C9.90146 14.5858 6.10247 14.5858 3.75933 12.2426L2.91833 11.4017V12.773C2.91833 13.1872 2.58255 13.523 2.16834 13.523C1.75412 13.523 1.41834 13.1872 1.41834 12.773L1.41833 9.59099C1.41833 9.17678 1.75412 8.84099 2.16833 8.84099H5.35031C5.76453 8.84099 6.10032 9.17678 6.10032 9.59099C6.10032 10.0052 5.76453 10.341 5.35032 10.341H3.97899L4.81999 11.182C6.57734 12.9393 9.42659 12.9393 11.1839 11.182C11.4696 10.8963 11.7082 10.5827 11.9003 10.2504C12.1077 9.89186 12.5665 9.76927 12.925 9.97662Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
