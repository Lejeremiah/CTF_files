import ddddocr

res = ""
ocr = ddddocr.DdddOcr()
for i in range(136):
    with open(f'imgs/{i}.png','rb') as img:
       code = ocr.classification(img.read())

       print(code)
       res += code
print(res)
1594199391770250354455183081054802631580554590456781276981302978243348088576774816981145460077422136047780972200375212293357383685099969525103172039042888918139627966684645793042724447954308373948403404873262837470923601139156304668538304057819343713500158029312192443296076902692735780417298059011568971988619463802818660736654049870484193411780158317168232187100668526865378478661078082009408188033574841574337151898932291631715135266804518790328831268881702387643369637508117317249879868707531954723945940226278368605203277838681081840279552
1594199391770250354455183081054802631580554590456781276981302978243348088576774816981145460077422136047780972200375212293357383685099969525103172039042888918139627966684645793042724447954308373948403404873262837470923601139156304668538304057819343713500158029312192443296076902692735780417298059011568971988619463802818660736654049870484193411780158317168232187100668526865378478661078082009408188033574841574337151898932291631715135266804518790328831268881702387643369637508117317249879868707531954723945940226278368605203277838681081840279552