package pdf_grid

import(
	"gridpaging"
	"github.com/signintech/gopdf"
)

func RectText(pdf *gopdf.GoPdf,x,y,w,h,textHeight,border float64,fnt,txt string,sz,align int){//align 0:left,1:center,2:right
	if border>0 {
		pdf.SetLineWidth(border)
		pdf.SetFillColor(255,255,255)
		pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
		pdf.SetFillColor(0,0,0)
	}
	err := pdf.SetFont(fnt, "", sz)
	if err==nil {
		ww,_ := pdf.MeasureTextWidth(txt)
		xx := x+1
		yy := y+(h-textHeight)/2
		switch align{
			case 1: 
				xx = x+(w-ww)/2
			case 2:
				xx = x+w-ww-1
		}
		pdf.SetX(xx)
		pdf.SetY(yy)
		pdf.Cell(nil,txt)
	}
}

func WriteGrid( pdf *gopdf.GoPdf,x,y,w,rowHeight,textHeight float64, fields []string, widths []float64,lines []gridpaging.Dataline,
										clines []gridpaging.Cellline,istart,nlns int,fontname string,fontsize int ){
	pdf.SetFont(fontname, "", fontsize)
	ncols := len(fields)
	nlines := nlns
	nn := len(lines)-istart+1
	if nlines>nn {
		nlines = nn
	}

	h := float64(nlines)*rowHeight
	pdf.SetLineWidth(1)
	pdf.SetFillColor(255,255,255)
	pdf.RectFromUpperLeftWithStyle(x, y, w, h, "FD")
	pdf.SetFillColor(0,0,0)
	pdf.SetLineWidth(0.1)
	xx := x
	yy := y	
	for i:=0;i<ncols;i++ {	//title and vertical lines
		RectText(pdf,xx,yy,widths[i],rowHeight,textHeight,0,"stsong",fields[i],fontsize,1)
		xx += widths[i]
		if i<ncols-1 {
			pdf.Line(xx, yy, xx, yy+h)
		}
	}
	yy += rowHeight
	pdf.Line( x,yy,x+w,yy )
	for j:=0;j<nlines-1;j++ {	//skip title line -> nlines-1
		dtline := lines[istart+j]
		if dtline.LineNOInRow==0 {
			pdf.Line( x,yy,x+w,yy )	//horizontal line
		}
		xx = x
		for i:=0;i<ncols;i++ {
			align := 1
			if i>0 { align = 0 }
			txt := dtline.Cols[i]
			if len(txt)>0 {
				hh := 1
				if clines[dtline.IRow].Lines[i]==1 {
					hh = dtline.NLinesInRow-dtline.LineNOInRow
					if hh>nlines-1-j {
						hh = nlines-1-j
					}
				}
				RectText(pdf,xx,yy,widths[i],float64(hh)*rowHeight,textHeight,0,fontname,txt,fontsize,align)	//dtline.Bg
			}
			xx += widths[i]
		}
	    yy += rowHeight
	}
}
