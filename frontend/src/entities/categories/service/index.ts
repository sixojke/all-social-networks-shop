import { rootApi } from "@/shared/api";
import { IGetAllCategoriesResponse } from "../types";

const categoriesApi = rootApi.injectEndpoints({
  endpoints: (build) => ({
    getAllCategories: build.query<IGetAllCategoriesResponse, void>({
      query: () => ({
        url: "/categories",
        method: "GET",
      }),
    }),
  }),
});

export const { useGetAllCategoriesQuery } = categoriesApi;
