import http from 'k6/http'

export let options = {
    vus: 5,
    duration: '5s',
    // iterations: 10, 
}

export default function() {
    // http.get('http://localhost:8088/home');
    http.get('http://host.docker.internal:8088/home');
}