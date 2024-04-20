import { HeaderLayout } from "@/widgets/Header";
import { FC, ReactNode } from "react";
import { Inter } from "next/font/google";

type Props = {
  children: ReactNode;
};

const inter = Inter({ subsets: ["latin"] });

export const AppLayout: FC<Props> = ({ children }) => {
  return (
    <div className={inter.className}>
      <HeaderLayout />
      {children}
    </div>
  );
};
