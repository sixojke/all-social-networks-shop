import { protectedApi } from "@/shared/api/protected";
import { IGetAdminLogsRequest, IGetAdminLogsResponse } from "../types";

const logsApi = protectedApi.injectEndpoints({
  endpoints: (build) => ({
    getAdminLogs: build.query<IGetAdminLogsResponse, IGetAdminLogsRequest>({
      query: (params) => ({
        url: "/admin/log",
        method: "GET",
        params,
      }),
    }),
  }),
});

export const { useGetAdminLogsQuery } = logsApi;
