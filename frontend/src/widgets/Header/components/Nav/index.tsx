import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";
import supplierIcon from "@/assets/icons/supplier.svg";
import bottomArrowIcon from "@/assets/icons/arrow-bottom.svg";
import Image from "next/image";

export const Nav = () => {
  return (
    <nav>
      <ul className="flex gap-x-6">
        <li>
          <Button className="bg-main-dark-blue">
            <div className="flex items-center gap-x-3">
              Каталог
              <Image alt="" src={bottomArrowIcon} height={22} width={24} />
            </div>
          </Button>
        </li>
        <li>
          <ButtonText>Новости</ButtonText>
        </li>
        <li>
          <ButtonText>Правила и помощь</ButtonText>
        </li>
        <li>
          <ButtonText>Контакты</ButtonText>
        </li>
        <li>
          <Button className="bg-main-dark-blue">
            <div className="flex items-center gap-x-3">
              Для поставщиков
              <Image alt="" src={supplierIcon} height={21} width={21} />
            </div>
          </Button>
        </li>
      </ul>
    </nav>
  );
};
