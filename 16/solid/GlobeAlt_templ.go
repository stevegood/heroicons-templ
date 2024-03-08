// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeAlt() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.7571 4.5C3.93668 4.71745 4.13258 4.92081 4.34293 5.10822C4.49575 4.49822 4.69651 3.93309 4.93881 3.42979C4.49497 3.72769 4.09682 4.08864 3.7571 4.5ZM8 1C5.58576 1 3.45737 2.22284 2.20015 4.07948C1.44254 5.19829 1 6.5486 1 8C1 8.15013 1.00474 8.29925 1.01408 8.44721C1.24509 12.1052 4.28419 15 8 15C11.7158 15 14.7549 12.1052 14.9859 8.44721C14.9953 8.29925 15 8.15013 15 8C15 6.5486 14.5575 5.19829 13.7999 4.07948C12.5426 2.22284 10.4142 1 8 1ZM8 2.5C7.52424 2.5 6.90855 2.88577 6.36732 3.9266C6.07368 4.49129 5.83596 5.19363 5.68416 5.99028C6.38814 6.3175 7.17278 6.5 8 6.5C8.82722 6.5 9.61186 6.3175 10.3158 5.99028C10.164 5.19363 9.92632 4.49129 9.63268 3.9266C9.09145 2.88577 8.47576 2.5 8 2.5ZM11.6571 5.10822C11.5043 4.49822 11.3035 3.93309 11.0612 3.42979C11.505 3.72769 11.9032 4.08864 12.2429 4.5C12.0633 4.71745 11.8674 4.92081 11.6571 5.10822ZM10.4909 7.54371C9.7171 7.83848 8.87743 8 8 8C7.12257 8 6.2829 7.83848 5.50908 7.54371C5.50307 7.69396 5.5 7.84613 5.5 8C5.5 8.7654 5.57609 9.48859 5.71082 10.1444C6.43245 10.3754 7.20162 10.5 8 10.5C8.79838 10.5 9.56755 10.3754 10.2892 10.1444C10.4239 9.48859 10.5 8.7654 10.5 8C10.5 7.84613 10.4969 7.69396 10.4909 7.54371ZM11.9239 9.39308C11.974 8.9414 12 8.4753 12 8C12 7.5869 11.9803 7.18075 11.9423 6.78498C12.3471 6.50859 12.7214 6.19087 13.059 5.83797C13.3429 6.50123 13.5 7.23175 13.5 8C13.5 8.03342 13.4997 8.06677 13.4991 8.10004C13.0361 8.59906 12.5063 9.0348 11.9239 9.39308ZM9.75208 11.8295C9.18534 11.9413 8.5995 12 8 12C7.4005 12 6.81466 11.9413 6.24792 11.8295C6.28658 11.9134 6.3264 11.9947 6.36732 12.0734C6.90855 13.1142 7.52424 13.5 8 13.5C8.47576 13.5 9.09145 13.1142 9.63268 12.0734C9.6736 11.9947 9.71342 11.9134 9.75208 11.8295ZM11.0613 12.5699C11.2526 12.1726 11.418 11.7368 11.5543 11.2708C12.0106 11.0744 12.4475 10.8413 12.8611 10.5754C12.4342 11.3795 11.8135 12.0651 11.0613 12.5699ZM4.93867 12.5699C4.74741 12.1726 4.58205 11.7368 4.44572 11.2708C3.98936 11.0744 3.55251 10.8413 3.1389 10.5754C3.56584 11.3795 4.18651 12.0651 4.93867 12.5699ZM2.50089 8.10004C2.96391 8.59906 3.49374 9.0348 4.07614 9.39308C4.02604 8.9414 4 8.4753 4 8C4 7.5869 4.01967 7.18075 4.05771 6.78498C3.65289 6.50859 3.27862 6.19087 2.94101 5.83797C2.65714 6.50123 2.5 7.23175 2.5 8C2.5 8.03342 2.5003 8.06677 2.50089 8.10004Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
