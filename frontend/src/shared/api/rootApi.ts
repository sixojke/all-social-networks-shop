import { createApi } from "@reduxjs/toolkit/query/react";
import { baseQuery } from "./baseQuery";

export const rootApi = createApi({
  reducerPath: "rootApi",
  baseQuery: baseQuery,
  tagTypes: ["products"],
  endpoints: () => ({}),
});
