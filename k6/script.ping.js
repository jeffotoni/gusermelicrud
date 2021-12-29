import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

export let options = {
  vus: 10,
  duration: '30s',
};

export default function () {
  var url = `http://localhost:8081/v1/user`;
  var url2 = `http://localhost:8082/v1/user`;
  http.get(url+`/ping`,{ headers: headers });
	http.get(url2+`/ping`,{ headers: headers });
}

