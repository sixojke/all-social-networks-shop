import { rootApi } from "@/shared/api";
import {
  IGetAllSubcategoriesRequest,
  IGetAllSubcategoriesResponse,
} from "../types";

const subcategoriesApi = rootApi.injectEndpoints({
  endpoints: (build) => ({
    getAllSubcategories: build.query<
      IGetAllSubcategoriesResponse,
      IGetAllSubcategoriesRequest
    >({
      query: ({ category_id }) => ({
        url: `/categories/${category_id}`,
        method: "GET",
      }),
    }),
  }),
});

export const { useGetAllSubcategoriesQuery } = subcategoriesApi;
