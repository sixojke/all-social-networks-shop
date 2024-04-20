import { fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const baseUrl = "http://localhost:8009/api/v1/";

export const baseQuery = fetchBaseQuery({
  baseUrl: baseUrl,
  credentials: "same-origin",
});
