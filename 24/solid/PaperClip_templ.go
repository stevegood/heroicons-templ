// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func PaperClip() templ.Component {
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
		templ_7745c5c3_Err = PaperClipWithAttrs(templ.Attributes{}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func PaperClipWithAttrs(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M18.9696 3.65901C18.0909 2.78033 16.6663 2.78033 15.7876 3.65901L4.84835 14.5983C3.38388 16.0628 3.38388 18.4372 4.84835 19.9016C6.31282 21.3661 8.68718 21.3661 10.1517 19.9016L17.8446 12.2087C18.1375 11.9158 18.6124 11.9158 18.9053 12.2087C19.1982 12.5015 19.1982 12.9764 18.9053 13.2693L11.2123 20.9623C9.16206 23.0126 5.83794 23.0126 3.78769 20.9623C1.73744 18.9121 1.73744 15.5879 3.78769 13.5377L14.7269 2.59835C16.1914 1.13388 18.5658 1.13388 20.0302 2.59835C21.4947 4.06281 21.4947 6.43718 20.0302 7.90165L9.09735 18.8346C9.09452 18.8374 9.09167 18.8403 9.08878 18.8432L9.08191 18.85L9.07971 18.8522L9.07747 18.8544C8.19765 19.7196 6.78319 19.7152 5.90901 18.841C5.03033 17.9623 5.03033 16.5377 5.90901 15.659L13.7195 7.84833C14.0124 7.55543 14.4873 7.55542 14.7802 7.84831C15.0731 8.1412 15.0731 8.61608 14.7802 8.90897L6.96968 16.7197C6.67678 17.0125 6.67678 17.4874 6.96967 17.7803C7.26041 18.0711 7.73056 18.0732 8.02383 17.7867L18.9696 6.84099C19.8483 5.96231 19.8483 4.53769 18.9696 3.65901Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
