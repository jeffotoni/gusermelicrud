import http from 'k6/http';
import { SharedArray } from "k6/data";
import { sleep } from 'k6';

var payload = new SharedArray("users", function () {
  var f = JSON.parse(open("../userwrite/json/user.post.k6.json"));
  //console.log(JSON.stringify(f));
  return f;
});

export let options = {
  vus: 10,
  duration: '1m',
  summaryTrendStats: ['avg', 'min', 'med', 'max', 'p(95)', 'p(99)', 'p(99.99)', 'count'],

  thresholds: {
    http_req_duration: ['avg<100', 'p(95)<200'],
    'http_req_connecting{cdnAsset:true}': ['p(95)<100'],
  },

};

export default function () {
  var url = `http://localhost:8081/v1/user`;
  var params = {
    headers: {
      'Content-Type': 'application/json',
    },
  };

  console.log(JSON.stringify(payload))
  return
  
  http.post(url,  JSON.stringify(payload), params);
}

