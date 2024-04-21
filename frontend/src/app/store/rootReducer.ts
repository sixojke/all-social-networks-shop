import { authApi } from "@/features/Auth";
import { rootApi } from "@/shared/api";
import { protectedApi } from "@/shared/api/protected";
import { combineReducers } from "@reduxjs/toolkit";

export const rootReducer = combineReducers({
  [rootApi.reducerPath]: rootApi.reducer,
  [authApi.reducerPath]: authApi.reducer,
  [protectedApi.reducerPath]: protectedApi.reducer,
});
