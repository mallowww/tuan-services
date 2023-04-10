import http from 'k6/http'

export default function() {
    // http.get('http://localhost:1677/home');
    http.get('http://host.docker.internal:1677/home');
    
    
}