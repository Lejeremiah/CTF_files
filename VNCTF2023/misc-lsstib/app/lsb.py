from PIL import Image
import numpy as np
from Crypto.Util.number import long_to_bytes

def lsb_decode(pic):
    pic_data = np.array(Image.open(pic).convert('RGB'))
    a,b,c = pic_data.shape
    # print(a,b,c)# a:y b:x
    res_data = pic_data.reshape(a*b*c)
    data = ''.join(str(i%2) for i in res_data)
    res = ''.join(chr(int(data[i:i+8],2)) for i in range(0,len(data),8))
    return res