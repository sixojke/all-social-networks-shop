import { protectedApi } from "@/shared/api/protected";
import { IGetActiveUserResponse } from "../types";

const userApi = protectedApi.injectEndpoints({
  endpoints: (build) => ({
    getActiveUser: build.query<IGetActiveUserResponse, void>({
      query: () => ({ url: "/user", method: "GET" }),
      providesTags: ["user"],
    }),
  }),
});

export const { useGetActiveUserQuery } = userApi;
