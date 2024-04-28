import { useCheckRefreshTokenMutation } from "@/features/Auth";
import { FC, ReactNode, useEffect } from "react";

type Props = {
  children: ReactNode;
};

export const AppInitializer: FC<Props> = ({ children }) => {
  const [getNewToken] = useCheckRefreshTokenMutation();
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
