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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.5 2C2.67157 2 2 2.67157 2 3.5V5C2 6.14856 2.14913 7.26341 2.42949 8.32576C3.61908 12.8334 7.16665 16.3809 11.6742 17.5705C12.7366 17.8509 13.8514 18 15 18H16.5C17.3284 18 18 17.3284 18 16.5V15.3516C18 14.6486 17.5117 14.0399 16.8254 13.8873L13.6024 13.1711C12.8276 12.999 12.0528 13.4602 11.8348 14.2233L11.5682 15.1561C11.4509 15.5669 11.0134 15.7989 10.6184 15.636C7.79126 14.47 5.53001 12.2087 4.36396 9.38159C4.20107 8.98665 4.4331 8.54913 4.84388 8.43176L5.77667 8.16525C6.53984 7.9472 7.00105 7.17238 6.82887 6.39757L6.11265 3.1746C5.96014 2.4883 5.35142 2 4.64837 2H3.5ZM13.2803 2.21967C12.9874 1.92678 12.5126 1.92678 12.2197 2.21967C11.9268 2.51256 11.9268 2.98744 12.2197 3.28033L13.9393 5L12.2197 6.71967C11.9268 7.01256 11.9268 7.48744 12.2197 7.78033C12.5126 8.07322 12.9874 8.07322 13.2803 7.78033L15 6.06066L16.7197 7.78033C17.0126 8.07322 17.4874 8.07322 17.7803 7.78033C18.0732 7.48744 18.0732 7.01256 17.7803 6.71967L16.0607 5L17.7803 3.28033C18.0732 2.98744 18.0732 2.51256 17.7803 2.21967C17.4874 1.92678 17.0126 1.92678 16.7197 2.21967L15 3.93934L13.2803 2.21967Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}