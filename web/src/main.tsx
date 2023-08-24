/** @format */

import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";

// Create an Axios instance with your desired configuration
/* const http: AxiosInstance = axios.create({
  baseURL: "http://localhost:9000",
  timeout: 1000,
  headers: {
    "Content-Type": "application/json",
  },
});

function axiosFetch(input, init) {
  // Convert RequestInfo or URL to a string if necessary
  const url = input instanceof URL ? input.toString() : input;

  // Convert RequestInit to Axios config
  const axiosConfig = {
    method: init?.method || "GET", // Default to GET if no method is provided
    headers: init?.headers || {},
    data: init?.body || undefined,
    // Handle other options as needed
  };

  return axios(url, axiosConfig)
    .then((response) => {
      // Convert Axios response to a Response-like object
      const responseInit = {
        status: response.status,
        statusText: response.statusText,
        headers: response.headers,
      };

      return new Response(response.data, responseInit);
    })
    .catch((error) => {
      // Handle Axios error or transform it into a Response-like object
      const responseInit = {
        status: error.response?.status || 500,
        statusText: error.message,
        headers: error.response?.headers || {},
      };

      return Promise.reject(new Response(undefined, responseInit));
    });
} */

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
