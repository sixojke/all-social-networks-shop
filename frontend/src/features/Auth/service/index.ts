import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import {
  CheckTokenRequest,
  CheckTokenResponse,
  SignInRequest,
  SignInResponse,
  SignUpRequest,
  SignUpResponse,
  SignUpVerifyRequest,
} from "../types";
import { useRouter } from "next/navigation";

const baseUrl = "http://localhost:8009/api/v1/users/";

const authBaseQuery = fetchBaseQuery({
  baseUrl,
  credentials: "include",
});

export const authApi = createApi({
  reducerPath: "auth",
  baseQuery: authBaseQuery,
  tagTypes: ["auth"],
  endpoints: (build) => ({
    signUp: build.mutation<SignUpResponse, SignUpRequest>({
      query: (data) => ({
        url: "sign-up",
        method: "POST",
        body: data,
      }),
      invalidatesTags: ["auth"],
    }),
    signUpVerify: build.mutation<void, SignUpVerifyRequest>({
      query: (data) => ({
        url: "verify",
        method: "POST",
        body: data,
      }),
      invalidatesTags: ["auth"],
    }),
    signIn: build.mutation<SignInResponse, SignInRequest>({
      query: (data) => ({ url: "sign-in", method: "POST", body: data }),
      invalidatesTags: ["auth"],
    }),
    checkRefreshToken: build.mutation<CheckTokenResponse, CheckTokenRequest>({
      query: (data) => ({ url: "auth/refresh", body: data, method: "POST" }),
      invalidatesTags: ["auth"],
    }),
  }),
});

export const {
  useSignUpMutation,
  useCheckRefreshTokenMutation,
  useSignInMutation,
  useSignUpVerifyMutation,
} = authApi;

export const useLogout = () => {
  const router = useRouter();
  const logout = () => {
    router.replace("/catalog");
    localStorage.clear();
    router.refresh();
  };
  return logout;
};
