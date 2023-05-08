import axios from 'axios';
import store from '../store/store';

const client = axios.create({
    baseURL: "https://nckh-be.onrender.com",
    // baseURL:"http://localhost:8080",
    headers: {
        'Accept': "application/json",
        'Content-Type': "application/json",
    }
})

client.interceptors.request.use(config => {
    const authToken = store.getState().login.token.accessToken;
    if (authToken) {
        config.headers['Authorization'] = "Bearer " + authToken
    }
    return config
})

export default client