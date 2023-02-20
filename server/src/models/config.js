const mysql = require("mysql");

const con = mysql.createConnection({
  host: "localhost",
  port: "3306",
  database: "nckh",
  user: "admin",
  password: "secret",
});

con.connect((err) => {
    if (err) {
        console.log("Connection error: " + err);
        return;
    }
    console.log("Connection successful!");
})

module.exports = con;