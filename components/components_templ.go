// Code generated by templ@v0.2.334 DO NOT EDIT.

package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/Ed1123/us-visa-wait-times/usvisa"
import "fmt"

func Table(cities []usvisa.CityWaitTime) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<table><thead><tr><th>")
		if err != nil {
			return err
		}
		var_2 := `City`
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th><th>")
		if err != nil {
			return err
		}
		var_3 := `Wait Time`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</th></tr></thead><tbody>")
		if err != nil {
			return err
		}
		for _, city := range cities {
			_, err = templBuffer.WriteString("<tr><td>")
			if err != nil {
				return err
			}
			var var_4 string = city.CityName
			_, err = templBuffer.WriteString(templ.EscapeString(var_4))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</td>")
			if err != nil {
				return err
			}
			if city.BusinessTourismVisitor.Days != nil {
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				var var_5 string = fmt.Sprint(*city.BusinessTourismVisitor.Days)
				_, err = templBuffer.WriteString(templ.EscapeString(var_5))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
			} else if city.BusinessTourismVisitor.Message != nil {
				_, err = templBuffer.WriteString("<td>")
				if err != nil {
					return err
				}
				var var_6 string = string(*city.BusinessTourismVisitor.Message)
				_, err = templBuffer.WriteString(templ.EscapeString(var_6))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</td>")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("</tr>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</tbody></table>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func Index() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_7 := templ.GetChildren(ctx)
		if var_7 == nil {
			var_7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<html><head><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if err != nil {
			return err
		}
		var_8 := `US Visa Wait Times`
		_, err = templBuffer.WriteString(var_8)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title><link rel=\"stylesheet\" href=\"https://cdnjs.cloudflare.com/ajax/libs/bulma/0.7.2/css/bulma.min.css\"></head><body><header class=\"section\"><div class=\"content\"><h1 class=\"title\">")
		if err != nil {
			return err
		}
		var_9 := `US Visa Wait Times`
		_, err = templBuffer.WriteString(var_9)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h1><p class=\"subtitle\">")
		if err != nil {
			return err
		}
		var_10 := `Wait times for US visa appointments at consulates around the world.`
		_, err = templBuffer.WriteString(var_10)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p></div></header><section class=\"section\"><div class=\"content\"><h2 class=\"subtitle\">")
		if err != nil {
			return err
		}
		var_11 := `APIs`
		_, err = templBuffer.WriteString(var_11)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ul><li><a href=\"/wait-times\">")
		if err != nil {
			return err
		}
		var_12 := `Wait Times`
		_, err = templBuffer.WriteString(var_12)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li><a href=\"/wait-times-with-country\">")
		if err != nil {
			return err
		}
		var_13 := `Wait Times with Countries`
		_, err = templBuffer.WriteString(var_13)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div><div class=\"content\"><h2 class=\"subtitle\">")
		if err != nil {
			return err
		}
		var_14 := `Tables`
		_, err = templBuffer.WriteString(var_14)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</h2><ul><li><a href=\"/table-js\">")
		if err != nil {
			return err
		}
		var_15 := `Working table from ChatGPT`
		_, err = templBuffer.WriteString(var_15)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li><li><a href=\"/table\">")
		if err != nil {
			return err
		}
		var_16 := `Table made with just Go/Templ`
		_, err = templBuffer.WriteString(var_16)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ul></div></section></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}