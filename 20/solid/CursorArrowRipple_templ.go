// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CursorArrowRipple() templ.Component {
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
		templ_7745c5c3_Err = CursorArrowRippleWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CursorArrowRippleWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M6.111 11.8891C3.96311 9.7412 3.96311 6.2588 6.111 4.11091C8.25888 1.96303 11.7413 1.96303 13.8892 4.11091C14.9633 5.185 15.5001 6.59127 15.5001 8C15.5001 8.41421 15.8359 8.75 16.2501 8.75C16.6643 8.75 17.0001 8.41421 17.0001 8C17.0001 6.2097 16.3165 4.41694 14.9498 3.05025C12.2162 0.316582 7.78401 0.316582 5.05034 3.05025C2.31667 5.78392 2.31667 10.2161 5.05034 12.9497C5.34323 13.2426 5.8181 13.2426 6.111 12.9497C6.40389 12.6569 6.40389 12.182 6.111 11.8891Z\" fill=\"#0F172A\"></path> <path d=\"M8.23235 6.23223C7.25604 7.20854 7.25604 8.79146 8.23235 9.76777C8.52525 10.0607 8.52525 10.5355 8.23235 10.8284C7.93946 11.1213 7.46459 11.1213 7.17169 10.8284C5.60959 9.26633 5.60959 6.73367 7.17169 5.17157C8.73379 3.60948 11.2664 3.60948 12.8285 5.17157C13.6094 5.95247 14.0001 6.97747 14.0001 8C14.0001 8.41421 13.6643 8.75 13.2501 8.75C12.8359 8.75 12.5001 8.41421 12.5001 8C12.5001 7.35903 12.2562 6.72054 11.7679 6.23223C10.7916 5.25592 9.20866 5.25592 8.23235 6.23223Z\" fill=\"#0F172A\"></path> <path d=\"M10.7657 7.51062C10.5871 7.24492 10.2596 7.12184 9.95029 7.20417C9.64095 7.2865 9.41795 7.5561 9.3951 7.8754L8.90409 14.7363C8.88303 15.0306 9.0365 15.3099 9.29622 15.4499C9.55595 15.59 9.87365 15.5647 10.108 15.3854L11.1508 14.5873L12.1363 18.2653C12.2436 18.6654 12.6548 18.9028 13.0549 18.7956C13.455 18.6884 13.6924 18.2771 13.5852 17.877L12.6083 14.2312L13.9005 14.4349C14.1951 14.4814 14.4893 14.3489 14.6497 14.0974C14.8101 13.846 14.8062 13.5233 14.6398 13.2758L10.7657 7.51062Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
