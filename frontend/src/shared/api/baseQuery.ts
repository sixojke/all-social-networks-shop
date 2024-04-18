import { fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const baseUrl = "http://localhost:8009/api/v1/";

export const baseQuery = fetchBaseQuery({
  mode: "no-cors",
  prepareHeaders: (headers) => {
    headers.set("Content-Type", "application/json");
    return headers;
  },
  baseUrl: baseUrl,
  credentials: "include",
});
