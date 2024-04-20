import { ReactNode } from "react";

export interface ModalHandlers {
  onOK: (() => void) | null;
  onHide: (() => void) | null;
}

export type ModalContextValue = {
  content: ReactNode;
  modalHandlers: ModalHandlers;
  showModal(content: ReactNode): void;
  hideModal(): void;
  setModalHandlers: (handlers: ModalHandlers) => void;
};
