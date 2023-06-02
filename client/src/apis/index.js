import axios from "axios";
import store from "../store/store";

// let url = "http://localhost:8080";
let url = "https://nckh-be.onrender.com"

const client = axios.create({
  baseURL: url,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json"
  }
});

client.interceptors.request.use((config) => {
  const authToken = store.getState().login.token.accessToken;
  if (authToken) {
    config.headers["Authorization"] = "Bearer " + authToken;
  }
  return config;
});


export default client;
