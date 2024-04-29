import { rootApi } from "@/shared/api";
import { configureStore } from "@reduxjs/toolkit";
import { rootReducer } from "./rootReducer";
import { authApi } from "@/features/Auth";
import { protectedApi } from "@/shared/api/protected";

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(
      rootApi.middleware,
      authApi.middleware,
      protectedApi.middleware
    ),
});

export type RootState = ReturnType<typeof rootReducer>;
