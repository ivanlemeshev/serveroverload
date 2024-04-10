import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '2m', target: 1000 },
    { duration: '2m', target: 1000 },
    { duration: '2m', target: 2000 },
    { duration: '2m', target: 2000 },
  ],
};

export default function () {
  http.get('http://localhost:8000/password/32');
  sleep(0.1); // 100ms
}
