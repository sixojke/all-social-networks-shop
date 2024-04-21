import { rootApi } from "@/shared/api";
import { configureStore } from "@reduxjs/toolkit";
import {
  FLUSH,
  PAUSE,
  PERSIST,
  PURGE,
  REGISTER,
  REHYDRATE,
} from "redux-persist";
import { rootReducer } from "./rootReducer";
import { authApi } from "@/features/Auth";
import { protectedApi } from "@/shared/api/protected";

export const store = configureStore({
  reducer: rootReducer,
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [FLUSH, REHYDRATE, PAUSE, PERSIST, PURGE, REGISTER],
      },
    }).concat(rootApi.middleware, authApi.middleware, protectedApi.middleware),
});

export type RootState = ReturnType<typeof rootReducer>;
