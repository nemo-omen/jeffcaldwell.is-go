// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "jeffcaldwell.is/view/layout"

func Index(current, remoteAddr string) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"content-container page\"><div class=\"section-header highlight\"><h2>About</h2></div><div class=\"post\"><p>Jeff Caldwell is a web developer who was born in Spain, raised in Texas, and lived in Brooklyn for 15 years.</p><p>He moved back to Texas to go to school a few years ago. Luckily, this resulted in Jeff obtaining a Bachelor of Science in Computer Science. Currently, he works as the supervisor of a small team of local digital journalists in San Angelo, Texas, where he lives with his fiancee.</p><p>He likes reading, writing, drawing, photography, and videography, although he hasn't had much time to commit to those pursuits while going to school and working full-time.</p><p>Jeff's looking to transition out of journalism and into software engineering and web development. He's spent the last few years learning frontend development. His favorite framework (sorry, library) is Svelte, but he can also work with React, Vue, or just plain 'ol JavaScript. He's has been learning a little more on backend development with Go. If you know of anyone looking for a web developer who loves to learn and has strong verbal and written communication skills, let him know on <a href=\"https://hachyderm.io/@trainingmontage\" rel=\"me\">Mastodon</a>.</p></div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layout.Base("JeffCaldwell — About", current, remoteAddr).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
