import router from "./route"
import axios from 'axios';

axios.interceptors.request.use(
    config => {
        if (localStorage.getItem('token')) {

            config.headers.token = localStorage.getItem('token');
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    });

axios.interceptors.response.use(
    response => {
        if (response.data.Code === 3) {
            localStorage.removeItem('token');
            router.replace('/login');
            console.log("token过期");
        }
        return response;
    },
    error => {
        return Promise.reject(error);
    }
);
export default axios;
