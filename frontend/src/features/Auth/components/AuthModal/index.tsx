import { useContext, useState } from "react";
import { SignInContent } from "./SignInContent";
import { ModalContext } from "@/shared/contexts/Modal";
import { FormProvider, useForm } from "react-hook-form";
import classNames from "classnames";
import { SignUpContent } from "./SignUpContent";

type ContentType = "signIn" | "signUp" | "dropPassword";

const TITLES = {
  signIn: "Авторизация",
  signUp: "Регистрация",
};

export const AuthModal = () => {
  const modalContext = useContext(ModalContext);
  const onHide = () => {
    modalContext?.hideModal();
  };
  const changeModalTabButtonClassname =
    "text-lg duration-300 text-main-blue border-b-2 border-main-white hover:text-main-dark-blue hover:text-main-dark-blue hover:border-b-2 hover:border-main-dark-blue";

  const [contentType, setContentType] = useState<ContentType>("signIn");

  const getContent = () => {
    if (contentType === "signIn") {
      return <SignInContent />;
    }
    if (contentType === "signUp") {
      return <SignUpContent />;
    }
  };

  const getBottomButtons = () => {
    if (contentType === "signIn") {
      return (
        <>
          <button
            onClick={() => setContentType("signUp")}
            className={classNames(changeModalTabButtonClassname)}
          >
            Создать аккаунт
          </button>
          <div className="h-2/4 w-[2px] bg-main-dark-blue" />
          <button
            onClick={() => setContentType("dropPassword")}
            className={classNames(changeModalTabButtonClassname)}
          >
            Сбросить пароль
          </button>
        </>
      );
    }
  };
  return (
    <div className="bg-main-white w-[680px] h-[550px] rounded-2xl flex items-center flex-col justify-between py-5 px-5">
      <div className="flex flex-col w-full items-center justify-between">
        <p
          onClick={onHide}
          className="rounded-full cursor-pointer select-none w-7 h-7 self-end bg-main-blue flex justify-center items-center text-main-white font-semibold hover:bg-main-dark-blue duration-150"
        >
          X
        </p>
        <p className="text-main-dark-blue font-semibold text-2xl mt-3">
          {TITLES[contentType as keyof typeof TITLES]}
        </p>
        <div className="mt-7 w-full px-24">{getContent()}</div>
      </div>
      <div className="flex gap-x-2 items-center">{getBottomButtons()}</div>
    </div>
  );
};
