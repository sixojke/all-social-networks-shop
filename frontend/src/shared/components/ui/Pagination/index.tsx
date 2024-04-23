import { FC } from "react";
import ReactPaginate from "react-paginate";
import styles from "./styles.module.css";

type Props = {
  currentPage: number;
  countPage: number;
  onChange: (selectedItem: number) => void;
};

const Arrow: FC<{ arrow: ">" | "<" }> = ({ arrow }) => {
  return (
    <span className="bg-main-dark-green duration-150 rounded-md py-[7px] px-[10px] text-main-white cursor-pointer hover:bg-main-black ">
      {arrow}
    </span>
  );
};

export const Pagination: FC<Props> = ({ currentPage, countPage, onChange }) => {
  return (
    <ReactPaginate
      previousLabel={<Arrow arrow="<" />}
      nextLabel={<Arrow arrow=">" />}
      containerClassName={styles.paginate}
      activeClassName={styles.active}
      pageClassName={styles.page}
      forcePage={currentPage - 1}
      pageCount={countPage}
      pageRangeDisplayed={2}
      marginPagesDisplayed={1}
      onPageChange={(page) => onChange(page.selected)}
    />
  );
};
