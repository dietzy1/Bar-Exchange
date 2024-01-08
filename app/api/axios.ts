/** @format */

import axios from "axios";

export const http = axios.create({
  baseURL: "http://localhost:9000/",
  timeout: 20000,
  headers: {
    "Content-Type": "application/json",
  },
});

/* 
http.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});
 */
