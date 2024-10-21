# app.py
from flask import Flask, request, render_template_string

app = Flask(__name__)

# HTML template with a vulnerability to XSS
HTML_TEMPLATE = """
<!doctype html>
<title>XSS Demo</title>
<h1>Welcome to the XSS Demo</h1>
<form method="GET" action="/">
    <label for="name">Enter your name:</label>
    <input type="text" name="name" id="name">
    <input type="submit" value="Submit">
</form>
{% if name %}
    <h2>Hello, {{ name }}!</h2>
    <script>alert('XSS Vulnerability!');</script>
{% endif %}
"""

@app.route("/")
def home():
    name = request.args.get("name", "")
    return render_template_string(HTML_TEMPLATE, name=name)

if __name__ == "__main__":
    app.run(debug=True)
