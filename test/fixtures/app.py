from flask import Flask

app = Flask(__name__)

@app.route("/users")
def get_users():
    return {"users": []}

@app.route("/users/:id")
def get_user(id):
    return {"user": {"id": id}}

@app.route("/api/v1/products")
def get_products():
    return {"products": []}

@app.route("/api/v1/products/:id")
def get_product(id):
    return {"product": {"id": id}}

@app.route("/admin/dashboard")
def admin_dashboard():
    return {"dashboard": "data"}

@app.route("/settings")
def get_settings():
    return {"settings": {}}

@app.route("/login")
def login():
    return {"status": "login"}
