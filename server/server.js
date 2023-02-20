const app = require("./src/app");

const PORT = 8000;

app.listen(PORT, () => {
    console.log("Server listening on port::" +PORT);
})

process.on("SIGOUT", () => {
    console.log("Server closed")
})