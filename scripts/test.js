import http from 'k6/http'

export let option = {
    vus: 10,
    duration: '1s',
    iterations: 10, 
}

export default function() {
    http.get('http://localhost:8088/home');
    // http.get('http://host.docker.internal:1677/home');
}