conn = new Mongo("mongodb://root:123456@localhost:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false");
db = conn.getDB("meliusers");
//db.auth("root", "123456");
printjson(db.users.createIndex({ first_name: "text", cpf: "text", last_name: "text", email:"text" }));
printjson({"msg":"ok"})
