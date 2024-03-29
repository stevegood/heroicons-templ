// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HandThumbUp() templ.Component {
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
		templ_7745c5c3_Err = HandThumbUpWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func HandThumbUpWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M7.49281 18.5C7.06823 18.5 6.67296 18.2635 6.51759 17.8684C6.18349 17.0187 6 16.0933 6 15.125C6 13.3759 6.59874 11.7667 7.60244 10.491C7.75335 10.2993 7.97456 10.1821 8.20214 10.094C8.67496 9.91091 9.09254 9.57968 9.4141 9.16967C10.1873 8.18384 11.1617 7.3634 12.2755 6.77021C12.9977 6.38563 13.6243 5.81428 13.9281 5.05464C14.1408 4.5231 14.25 3.95587 14.25 3.38338V2.75C14.25 2.33579 14.5858 2 15 2C16.2426 2 17.25 3.00736 17.25 4.25C17.25 5.40163 16.9904 6.49263 16.5266 7.46771C16.261 8.02604 16.6336 8.75 17.2519 8.75H20.3777C21.4044 8.75 22.3233 9.44399 22.432 10.4649C22.4769 10.8871 22.5 11.3158 22.5 11.75C22.5 14.5976 21.5081 17.2136 19.851 19.2712C19.4634 19.7525 18.8642 20 18.2462 20H14.2302C13.7466 20 13.2661 19.922 12.8072 19.7691L9.69278 18.7309C9.23393 18.578 8.75342 18.5 8.26975 18.5H7.49281Z\" fill=\"#0F172A\"></path> <path d=\"M2.33149 10.7271C1.79481 12.0889 1.5 13.5725 1.5 15.125C1.5 16.3451 1.68208 17.5226 2.02056 18.632C2.27991 19.482 3.10418 20 3.99289 20H4.90067C5.3462 20 5.62137 19.5017 5.42423 19.1022C4.83248 17.9029 4.5 16.5528 4.5 15.125C4.5 13.4168 4.97588 11.8198 5.8023 10.4593C6.0473 10.0559 5.77404 9.5 5.30212 9.5H4.24936C3.41733 9.5 2.63655 9.95303 2.33149 10.7271Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
