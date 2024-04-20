import { authApi } from "@/features/Auth";
import { rootApi } from "@/shared/api";
import { combineReducers } from "@reduxjs/toolkit";

export const rootReducer = combineReducers({
  [rootApi.reducerPath]: rootApi.reducer,
  [authApi.reducerPath]: authApi.reducer,
});
