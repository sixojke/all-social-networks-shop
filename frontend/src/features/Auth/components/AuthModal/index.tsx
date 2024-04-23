import { useContext, useState } from "react";
import { SignInContent } from "./SignInContent";
import { ModalContext } from "@/shared/contexts/Modal";
import classNames from "classnames";
import { SignUpContent } from "./SignUpContent";
import { VerifyContent } from "./VerifyContent";
import { isNumber } from "lodash";
import { Loading } from "./Loading";
import CloseIcon from "@/assets/icons/close-icon.svg";
import Image from "next/image";
import { ErrorContent } from "./ErrorContent";

export type ContentType =
  | "signIn"
  | "signUp"
  | "dropPassword"
  | "verify"
  | "loading"
  | "error";

const TITLES = {
  signIn: "Авторизация",
  signUp: "Регистрация",
  verify: "Введите код",
  loading: "Обработка запроса",
  error: "Ошибка",
};

export const AuthModal = () => {
  const modalContext = useContext(ModalContext);
  const onHide = () => {
    modalContext?.hideModal();
  };
  const changeModalTabButtonClassname =
    "text-lg duration-300 text-main-dark-green border-b-2 border-main-white hover:text-main-dark-green hover:text-main-black hover:border-b-2 hover:border-main-black";

  const [contentType, setContentType] = useState<ContentType>("signIn");
  const [userId, setUserId] = useState<number | null>(null);
  const setErrorContent = () => {
    setContentType("error");
  };
  const getContent = () => {
    switch (contentType) {
      case "signIn":
        return <SignInContent setErrorContent={setErrorContent} />;

      case "signUp":
        return (
          <SignUpContent
            setErrorContent={setErrorContent}
            setUserId={setUserId}
            setContentType={setContentType}
          />
        );

      case "loading":
        return <Loading />;

      case "verify":
        return (
          isNumber(userId) && (
            <VerifyContent setErrorContent={setErrorContent} userId={userId} />
          )
        );

      case "error":
        return <ErrorContent setContentType={() => setContentType("signIn")} />;
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
          <div className="h-2/4 w-[2px] bg-main-dark-green" />
          <button
            onClick={() => setContentType("dropPassword")}
            className={classNames(changeModalTabButtonClassname)}
          >
            Сбросить пароль
          </button>
        </>
      );
    }
    if (contentType === "signUp") {
      return (
        <button
          onClick={() => setContentType("signIn")}
          className={classNames(changeModalTabButtonClassname)}
        >
          Уже есть аккаунт?
        </button>
      );
    }
  };
  return (
    <div className="bg-white w-[680px] h-[550px] rounded-[10px] flex items-center flex-col justify-between py-5 px-5">
      <div className="flex flex-col w-full items-center justify-between">
        <p
          onClick={onHide}
          className="rounded-full cursor-pointer select-none w-7 h-7 self-end bg-main-dark-green flex justify-center items-center hover:bg-main-black duration-150"
        >
          <Image src={CloseIcon} width={12} height={12} alt="" />
        </p>
        <p className="text-main-dark-green font-semibold text-2xl mt-3">
          {TITLES[contentType as keyof typeof TITLES]}
        </p>
        <div className="mt-7 w-full px-24">{getContent()}</div>
      </div>
      <div className="flex gap-x-2 items-center">{getBottomButtons()}</div>
    </div>
  );
};
