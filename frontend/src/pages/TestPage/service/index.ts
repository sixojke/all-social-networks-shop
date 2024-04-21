import { protectedApi } from "@/shared/api/protected";
import { url } from "inspector";
import { method } from "lodash";

const testApi = protectedApi.injectEndpoints({
  endpoints: (build) => ({
    test: build.query({
      query: () => ({
        url: "/users/test",
        method: "GET",
      }),
    }),
  }),
});

export const { useTestQuery } = testApi;
