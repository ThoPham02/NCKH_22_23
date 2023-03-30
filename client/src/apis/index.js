import axios from 'axios';

const client = axios.create({
    baseURL: "https://nckh-be.onrender.com",
    headers: {
        Accept: "application/json",
        'Content-Type': "application/json"
    }
})

export default client