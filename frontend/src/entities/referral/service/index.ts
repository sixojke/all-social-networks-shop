import { rootApi } from "@/shared/api";
import { protectedApi } from "@/shared/api/protected";
import {
  IAddVisitorRequest,
  ICreateLinkRequest,
  ICreateLinkResponse,
  IDeleteLinkRequest,
  IGetStatsRequest,
  IGetStatsResponse,
} from "../types";

const referalApi = rootApi.injectEndpoints({
  endpoints: (build) => ({
    addVisitor: build.mutation<void, IAddVisitorRequest>({
      query: (data) => ({
        url: "/referal-system/visitor",
        method: "POST",
        params: data,
      }),
    }),
  }),
});

const referalApiProtected = protectedApi.injectEndpoints({
  endpoints: (build) => ({
    createReferralLink: build.mutation<ICreateLinkResponse, ICreateLinkRequest>(
      {
        query: (data) => ({
          url: "/admin/referral-system/create-code",
          body: data,
          method: "POST",
        }),
        invalidatesTags: ["referral"],
      }
    ),
    deleteReferralLink: build.mutation<void, IDeleteLinkRequest>({
      query: ({ code }) => ({
        url: `/admin/referral-system/${code}`,
        method: "DELETE",
      }),
      invalidatesTags: ["referral"],
    }),
    getRefferalStats: build.query<IGetStatsResponse, IGetStatsRequest>({
      query: (params) => ({
        url: `/admin/referral-system/stats`,
        params,
        method: "GET",
      }),
      providesTags: ["referral"],
    }),
  }),
});

export const { useAddVisitorMutation } = referalApi;

export const {
  useCreateReferralLinkMutation,
  useDeleteReferralLinkMutation,
  useGetRefferalStatsQuery,
} = referalApiProtected;
