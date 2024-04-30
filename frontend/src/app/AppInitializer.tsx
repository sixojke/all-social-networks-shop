import { useAddVisitorMutation } from "@/entities/referral";
import { useReferral } from "@/entities/referral/hooks/useReferral";
import { useCheckRefreshTokenMutation } from "@/features/Auth";
import { useParams, useSearchParams } from "next/navigation";
import { FC, ReactNode, useEffect } from "react";

type Props = {
  children: ReactNode;
};

export const AppInitializer: FC<Props> = ({ children }) => {
  const [getNewToken] = useCheckRefreshTokenMutation();
  useReferral();
  const refresh = () => {
    const refresh_token = localStorage.getItem("refreshToken");
    if (refresh_token) {
      getNewToken({ refresh_token })
        .unwrap()
        .then((res) => {
          localStorage.setItem("accessToken", res.accessToken);
          localStorage.setItem("refreshToken", res.refreshToken);
        });
    }
  };
  useEffect(() => {
    refresh();
    setTimeout(() => {
      refresh();
    }, 600000);
  }, []);
  return <>{children}</>;
};
