import axios from "axios";
import { env } from '$env/dynamic/public'

// Create a instance of axios to use the same base url.
let instanceBaseURL = "http://" + env.PUBLIC_HOSTNAME + ":443";
const axiosAPI = axios.create({
  baseURL : instanceBaseURL
});

// implement a method to execute all the request from here.
const apiRequest = (method: string, url: string, request: any) => {
    const headers = {
        // "authorization": "",
        "Access-Control-Allow-Origin": '*',
        'Content-Type': 'application/json',
    };
    //using the axios instance to perform the request that received from each http method
    return axiosAPI({
        method,
        url,
        data: request,
        headers,

      }).then(res => {
        return Promise.resolve(res.data);
      })
      .catch(err => {
        return Promise.reject(err);
      });
};

// function to execute the http get request
const get = (url: any,  request: any) => apiRequest("get",url,request);

// function to execute the http delete request
const deleteRequest = (url: any, request: any) =>  apiRequest("delete", url, request);

// function to execute the http post request
const post = (url: any, request: any) => apiRequest("post", url, request);

// function to execute the http put request
const put = (url: any, request: any) => apiRequest("put", url, request);

// function to execute the http path request
const patch = (url: any, request: any) =>  apiRequest("patch", url, request);

// expose your method to other services or actions
const Api ={
    get,
    delete: deleteRequest,
    post,
    put,
    patch
};
export default Api;
