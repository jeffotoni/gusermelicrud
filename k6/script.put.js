import http from 'k6/http';
import { sleep } from 'k6';

const headers = { 'Content-Type': 'application/json' };

export let options = {
  vus: 10,
  duration: '1m',
};

export default function () {
  var url = `http://localhost:8081/v1/user`;
  const data = {
   "first_name":"Jefferson",
   "last_name":"Otoni",
   "birthday":"1980-08-20"
  };
  http.put(url+`/29145037094`,JSON.stringify(data), { headers: headers } );
}
