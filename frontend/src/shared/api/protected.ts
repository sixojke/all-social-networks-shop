import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

let accessToken;
let refreshToken;

if (typeof window !== "undefined") {
  accessToken = localStorage.getItem("accessToken");
  refreshToken = localStorage.getItem("refreshToken");
}

const baseUrl = "http://localhost:8009/api/v1/";

export const baseQuery = fetchBaseQuery({
  baseUrl: baseUrl,
  credentials: "same-origin",
  headers: {
    Authorization: `Bearer=${accessToken}`,
  },
});

export const protectedApi = createApi({
  reducerPath: "protectedApi",
  tagTypes: ["user"],
  baseQuery,
  endpoints: () => ({}),
});
