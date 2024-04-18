import { fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const baseUrl = "http://localhost:8010/api/v1/";

export const baseQuery = fetchBaseQuery({
  prepareHeaders: (headers) => {
    headers.set("Content-Type", "application/json");
    return headers;
  },
  baseUrl: baseUrl,
  credentials: "include",
});
