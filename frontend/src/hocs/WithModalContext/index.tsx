import { ModalContext, ModalHandlers } from "@/shared/contexts/Modal";
import { FC, PropsWithChildren, ReactNode, useCallback, useState } from "react";

export const WithModalContext: FC<PropsWithChildren> = ({ children }) => {
  const [modalContent, setModalContent] = useState<ReactNode>(null);
  const [modalHandlers, setModalHandlers] = useState<ModalHandlers>({
    onOK: null,
    onHide: null,
  });

  const showModal = useCallback((content: ReactNode) => {
    setModalContent(content);
  }, []);

  const hideModal = useCallback(() => {
    setModalContent(null);
    setModalHandlers({
      onOK: null,
      onHide: null,
    });
  }, []);

  return (
    <ModalContext.Provider
      value={{
        content: modalContent,
        modalHandlers,
        showModal,
        hideModal,
        setModalHandlers,
      }}
    >
      {children}
    </ModalContext.Provider>
  );
};
