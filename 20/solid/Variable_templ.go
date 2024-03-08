// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Variable(classNames string) templ.Component {
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
		var templ_7745c5c3_Var2 = []any{classNames}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var2...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ.CSSClasses(templ_7745c5c3_Var2).String()))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15.212 2.07905C15.5825 1.894 16.033 2.0444 16.218 2.41498C17.3587 4.6992 18 7.2758 18 10C18 12.7243 17.3587 15.3009 16.218 17.5851C16.033 17.9557 15.5825 18.1061 15.212 17.921C14.8414 17.736 14.691 17.2856 14.876 16.915C15.9151 14.8342 16.5 12.4865 16.5 10C16.5 7.51359 15.9151 5.16592 14.876 3.0851C14.691 2.71452 14.8414 2.2641 15.212 2.07905ZM4.78803 2.07905C5.15861 2.2641 5.30901 2.71452 5.12397 3.0851C4.08491 5.16591 3.5 7.51359 3.5 10C3.5 12.4865 4.08491 14.8342 5.12397 16.915C5.30901 17.2856 5.15861 17.736 4.78803 17.921C4.41745 18.1061 3.96703 17.9557 3.78198 17.5851C2.64135 15.3009 2 12.7243 2 10C2 7.2758 2.64135 4.6992 3.78198 2.41498C3.96703 2.0444 4.41745 1.894 4.78803 2.07905ZM7.07277 5.63347C7.86037 5.10841 8.93143 5.43515 9.29184 6.31042L10.1483 8.39032L11.2942 6.61933C11.9758 5.56592 13.3873 5.27355 14.4313 5.96953L14.666 6.12601C15.0107 6.35577 15.1038 6.82142 14.874 7.16607C14.6443 7.51071 14.1786 7.60384 13.834 7.37408L13.5993 7.21761C13.2513 6.98561 12.7808 7.08307 12.5536 7.43421L10.8431 10.0777L12.0952 13.1185L12.834 12.626C13.1786 12.3962 13.6443 12.4894 13.874 12.834C14.1038 13.1787 14.0107 13.6443 13.666 13.8741L12.9272 14.3666C12.1396 14.8917 11.0686 14.5649 10.7082 13.6897L9.85173 11.6098L8.70579 13.3808C8.02418 14.4342 6.61265 14.7265 5.56869 14.0306L5.33397 13.8741C4.98933 13.6443 4.8962 13.1787 5.12596 12.834C5.35573 12.4894 5.82138 12.3962 6.16603 12.626L6.40074 12.7825C6.74873 13.0145 7.21923 12.917 7.44644 12.5659L9.15693 9.92239L7.90482 6.88155L7.16603 7.37408C6.82138 7.60384 6.35573 7.51071 6.12596 7.16607C5.8962 6.82142 5.98933 6.35577 6.33397 6.126L7.07277 5.63347Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
