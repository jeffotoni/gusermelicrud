import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

export let options = {
  vus: 10,
  duration: '30s',
};

export default function () {
  var url2 = `http://localhost:8082/v1/user`;
	http.get(url2+`?names=\[Paul\]`,{ headers: headers });
}

