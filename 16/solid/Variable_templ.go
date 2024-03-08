// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.598
package solid

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Variable() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.38045 3.01185C3.52293 2.62291 3.32314 2.19211 2.9342 2.04963C2.54526 1.90715 2.11446 2.10694 1.97198 2.49588C1.34296 4.21295 1 6.06709 1 7.99965C1 9.9322 1.34296 11.7863 1.97198 13.5034C2.11446 13.8923 2.54526 14.0921 2.9342 13.9497C3.32314 13.8072 3.52293 13.3764 3.38045 12.9874C2.81109 11.4332 2.5 9.75361 2.5 7.99965C2.5 6.24569 2.81109 4.56606 3.38045 3.01185Z\" fill=\"black\"></path> <path d=\"M12.6196 3.01185C12.4771 2.62291 12.6769 2.19211 13.0658 2.04963C13.4547 1.90715 13.8855 2.10694 14.028 2.49588C14.657 4.21295 15 6.06709 15 7.99965C15 9.9322 14.657 11.7863 14.028 13.5034C13.8855 13.8923 13.4547 14.0921 13.0658 13.9497C12.6769 13.8072 12.4771 13.3764 12.6196 12.9874C13.1889 11.4332 13.5 9.75361 13.5 7.99965C13.5 6.24569 13.1889 4.56606 12.6196 3.01185Z\" fill=\"black\"></path> <path d=\"M6.52255 4.78529C6.87451 4.67331 7.2556 4.83423 7.42078 5.16458L8.17869 6.68041L8.99056 5.77833C9.61503 5.08448 10.5879 4.82049 11.4774 5.10353C11.8722 5.22912 12.0903 5.65091 11.9647 6.04563C11.8391 6.44034 11.4174 6.65851 11.0226 6.53292C10.6946 6.42854 10.3358 6.52589 10.1055 6.78178L8.89942 8.12187L9.6334 9.58982L10.0214 9.46567C10.4159 9.33943 10.838 9.5569 10.9643 9.95141C11.0905 10.3459 10.873 10.7681 10.4785 10.8943L9.47854 11.2143C9.12632 11.327 8.74452 11.1662 8.57914 10.8354L7.82132 9.31976L7.00989 10.2213C6.38502 10.9156 5.41111 11.1791 4.52146 10.8944C4.12695 10.7681 3.90948 10.346 4.03572 9.95146C4.16197 9.55695 4.58412 9.33948 4.97862 9.46573C5.30622 9.57056 5.66485 9.47356 5.89495 9.2179L7.10059 7.8783L6.36604 6.40919L5.97736 6.53286C5.58265 6.65845 5.16085 6.44029 5.03526 6.04557C4.90967 5.65086 5.12784 5.22907 5.52255 5.10347L6.52255 4.78529Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
