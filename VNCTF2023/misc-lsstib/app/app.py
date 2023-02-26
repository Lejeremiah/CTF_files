from flask import *
import os
from lsb import lsb_decode

app = Flask(__name__, static_folder='static', static_url_path='/static')
app.config['ALLOWED_EXTENSIONS'] = {'png'} 

def allowed_file(filename):
    return '.' in filename and \
           filename.rsplit('.', 1)[1] in app.config['ALLOWED_EXTENSIONS']

@app.route('/')
def index():
    return redirect('/upload')

@app.route('/upload', methods=['GET', 'POST'])
def upload_file():
    if request.method == 'POST':
        f = request.files['file']
        if f and allowed_file(f.filename):
            return render_template_string(lsb_decode(f))
        else:
            return 'This is not a Upload challenge!!!\nPlz Upload right png!!!'
    else:
        return render_template('index.html')

if __name__ == '__main__':
    app.run(host='0.0.0.0',port=5000)