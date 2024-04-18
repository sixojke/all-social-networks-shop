import { Button } from "@/shared/components/ui/Buttons/Button";
import { ButtonText } from "@/shared/components/ui/Buttons/ButtonText";

export const Nav = () => {
  return (
    <nav>
      <ul className="flex gap-x-6">
        <li>
          <Button className="bg-main-dark-blue">Каталог</Button>
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
          <Button className="bg-main-dark-blue">Для поставщиков</Button>
        </li>
      </ul>
    </nav>
  );
};
