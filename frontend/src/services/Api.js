import axios from "axios";
import { env } from '$env/dynamic/public'

// Create a instance of axios to use the same base url.
const axiosAPI = axios.create({
  baseURL : "http://" + env.PUBLIC_HOSTNAME + ":443"
});

// implement a method to execute all the request from here.
const apiRequest = (/** @type {string} */ method, /** @type {String} */ url, /** @type {any} */ request) => {
    console.log("baseURL", axiosAPI.defaults.baseURL);
    const headers = {
        "authorization": "",
        "Access-Control-Allow-Origin": "*",
    };
    //using the axios instance to perform the request that received from each http method
    return axiosAPI({
        method,
        url,
        data: request,
        headers
      }).then(res => {
        return Promise.resolve(res.data);
      })
      .catch(err => {
        return Promise.reject(err);
      });
};

// function to execute the http get request
const get = (/** @type {any} */ url, /** @type {any} */ request) => apiRequest("get",url,request);

// function to execute the http delete request
const deleteRequest = (/** @type {any} */ url, /** @type {any} */ request) =>  apiRequest("delete", url, request);

// function to execute the http post request
const post = (/** @type {any} */ url, /** @type {any} */ request) => apiRequest("post", url, request);

// function to execute the http put request
const put = (/** @type {any} */ url, /** @type {any} */ request) => apiRequest("put", url, request);

// function to execute the http path request
const patch = (/** @type {any} */ url, /** @type {any} */ request) =>  apiRequest("patch", url, request);

// expose your method to other services or actions
const Api ={
    get,
    delete: deleteRequest,
    post,
    put,
    patch
};
export default Api;
