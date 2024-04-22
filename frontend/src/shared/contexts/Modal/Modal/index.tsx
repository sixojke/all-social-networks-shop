import { memo, useContext } from "react";
import { Modal } from "@/shared/components/ui/Modal";
import { ModalContext } from "../context";

export const AppModal = memo(() => {
  const context = useContext(ModalContext);

  const handleHideModal = () => {
    context?.modalHandlers?.onHide?.();
    context?.hideModal();
  };
  return (
    <Modal active={!!context?.content} onHide={handleHideModal}>
      {context?.content}
    </Modal>
  );
});

AppModal.displayName = "AppModal";
