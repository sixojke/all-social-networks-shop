import { Button } from "@/shared/components/ui/Buttons/Button";
import { ModalContext } from "@/shared/contexts/Modal";
import { useContext } from "react";
import { AuthModal } from "../AuthModal";

export const AuthButton = () => {
  const modalContext = useContext(ModalContext);
  const showModal = () => {
    modalContext?.showModal(<AuthModal />);
  };
  return (
    <Button onClick={showModal} className="bg-main-dark-green">
      Войти
    </Button>
  );
};
