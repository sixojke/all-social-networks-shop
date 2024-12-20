import { ICategory } from "@/entities/categories";
import { FC } from "react";

type Props = {
  categories: ICategory[];
  isLoading: boolean;
};

export const Footer: FC<Props> = ({ categories }) => {
  return (
    <footer className="flex justify-center px-[450px] flex-col gap-y-8 mb-16">
      <section className="bg-main-light-gray border border-main-black border-opacity-60 min-h-[204px] rounded-[10px] w-full flex flex-col items-center py-3">
        <p className="text-main-black text-[28px] font-semibold">
          Все наши категории
        </p>
        <div className="flex gap-4 justify-center mt-9">
          {Array.from({ length: 10 }, (_, index) => ({ id: index })).map(
            (item) => (
              <div
                key={item.id}
                className="w-12 h-12 bg-main-black rounded-[10px]"
              />
            )
          )}
        </div>
      </section>
      <section className="w-full h-[163px] bg-main-dark-green rounded-[10px]" />
    </footer>
  );
};
