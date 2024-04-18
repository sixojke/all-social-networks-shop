import { rootApi } from "@/shared/api";
import { IGetAllProductsRequest, IGetAllProductsResponse } from "../types";

const productsApi = rootApi.injectEndpoints({
  endpoints: (build) => ({
    getAllProducts: build.query<
      IGetAllProductsResponse,
      IGetAllProductsRequest
    >({
      query: (data) => ({
        url: "/products/",
        method: "GET",
        params: data,
      }),
      providesTags: ["products"],
    }),
  }),
});

export const { useGetAllProductsQuery } = productsApi;
