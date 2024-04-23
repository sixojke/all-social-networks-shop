import { useGetActiveUserQuery } from "@/entities/user";
import { AuthModal } from "@/features/Auth/components/AuthModal";
import { Loader } from "@/shared/components/ui/Loader";
import { ModalContext } from "@/shared/contexts/Modal";
import { useRouter } from "next/router";
import { FC, ReactNode, useContext, useEffect, useState } from "react";

type Props = {
  children: ReactNode;
};

export const ProtectedComponent: FC<Props> = ({ children }) => {
  const [isAuth, setIsAuth] = useState(false);
  const router = useRouter();
  const modalContext = useContext(ModalContext);
  const { data: activeUser, isLoading, isError } = useGetActiveUserQuery();
  useEffect(() => {
    if (!isLoading) {
      if (isError) {
        router.replace("/catalog");
        modalContext?.showModal(<AuthModal />);
        return;
      }
      setIsAuth(true);
    }
    return;
  }, [activeUser, isLoading]);
  if (isLoading) {
    return (
      <div className="w-full flex justify-center mt-72 items-center">
        <Loader />
      </div>
    );
  }
  if (isAuth) {
    return <>{children}</>;
  }
};
